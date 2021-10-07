package jwt

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"github.com/ONSdigital/dp-authorisation/v2/permissions"
	"github.com/golang-jwt/jwt/v4"
	"github.com/pkg/errors"
)

var (
	ErrFailedToParsePublicKey     = errors.New("error parsing public key for jwt verification")
	ErrUnexpectedKeyType          = errors.New("unexpected public key type for jwt verification")
	ErrInvalidSignature           = errors.New("jwt verification failed due to an invalid signature")
	ErrTokenExpired               = errors.New("jwt token has expired")
	ErrTokenNotYetValid           = errors.New("jwt token is not yet valid")
	ErrTokenMalformed             = errors.New("jwt token is malformed")
	ErrTokenInvalid               = errors.New("jwt token is not valid") // more generic error to catch any other cases
	ErrTokenUnsupportedEncryption = errors.New("only rsa encrypted jwt tokens are supported")
	ErrNoUserID                   = errors.New("jwt token does not have a user id")
	ErrFailedToParseClaims        = errors.New("failed to read claims from jwt token")
	ErrNoGroups                   = errors.New("jwt token does not have any groups")
)

// CognitoRSAParser parses JWT tokens that have an RSA encrypted signature, and contain AWS cognito specific claims.
type CognitoRSAParser struct {
	publicKey *rsa.PublicKey
	jwtParser *jwt.Parser
}

// NewCognitoRSAParser creates a new instance of CognitoRSAParser using the given public key value.
func NewCognitoRSAParser(base64EncodedPublicKey string) (*CognitoRSAParser, error) {
	publicKey, err := parsePublicKey(base64EncodedPublicKey)
	if err != nil {
		return nil, ErrFailedToParsePublicKey
	}

	jwtParser := &jwt.Parser{
		UseJSONNumber: true,
	}

	return &CognitoRSAParser{
		publicKey: publicKey,
		jwtParser: jwtParser,
	}, nil
}

// Parse and verify the given JWT token, and return the EntityData contained within the JWT (user ID and groups list)
func (p CognitoRSAParser) Parse(tokenString string) (*permissions.EntityData, error) {
	token, err := p.jwtParser.Parse(tokenString, p.getKey)
	if err != nil {
		err = determineErrorType(err)
		return nil, err
	}
	if !token.Valid {
		return nil, ErrTokenInvalid
	}

	entityData, err := getEntityData(token)
	if err != nil {
		return nil, err
	}

	return entityData, nil
}

// getEntityData takes a jwt token and reads its claims to determine the entity data (user ID and groups)
func getEntityData(token *jwt.Token) (*permissions.EntityData, error) {
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, ErrFailedToParseClaims
	}

	userID, ok := claims["username"].(string)
	if !ok {
		return nil, ErrNoUserID
	}

	jwtGroups, ok := claims["cognito:groups"].([]interface{})
	if !ok {
		return nil, ErrNoGroups
	}

	groups := mapToStringArray(jwtGroups)
	entityData := &permissions.EntityData{
		UserID: userID,
		Groups: groups,
	}
	return entityData, nil
}

// determineErrorType attempts to cast an error to the JWT libraries error type,
// allowing the specific error type to be determined.
func determineErrorType(err error) error {
	validationErr, ok := err.(*jwt.ValidationError)
	if ok {
		if validationErr.Errors&jwt.ValidationErrorMalformed != 0 {
			return ErrTokenMalformed
		}
		if validationErr.Errors&jwt.ValidationErrorSignatureInvalid != 0 {
			return ErrInvalidSignature
		}
		if validationErr.Errors&(jwt.ValidationErrorExpired) != 0 {
			return ErrTokenExpired
		}
	}
	return err
}

// mapToStringArray maps from []interface{} to []string
func mapToStringArray(jwtGroups []interface{}) []string {
	var groups []string
	for _, cognitoGroup := range jwtGroups {
		group, ok := cognitoGroup.(string)
		if ok {
			groups = append(groups, group)
		}
	}
	return groups
}

// parsePublicKey takes the raw base64 encoded public key value and creates an instance of rsa.PublicKey
func parsePublicKey(base64EncodedPublicKey string) (*rsa.PublicKey, error) {
	pubKeyBytes, err := base64.StdEncoding.DecodeString(base64EncodedPublicKey)
	if err != nil {
		return nil, err
	}

	publicKey, err := x509.ParsePKIXPublicKey(pubKeyBytes)
	if err != nil {
		return nil, err
	}

	rsaPublicKey, ok := publicKey.(*rsa.PublicKey)
	if !ok {
		return nil, ErrUnexpectedKeyType
	}

	return rsaPublicKey, nil
}

// This function signature is required by the JWT library. The function is passed as a parameter to jwt.Parse
func (p CognitoRSAParser) getKey(token *jwt.Token) (interface{}, error) {
	// check for expected signing method on the token, before trying to verify it.
	if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
		return nil, ErrTokenUnsupportedEncryption
	}

	return p.publicKey, nil
}
