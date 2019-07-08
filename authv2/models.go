package authv2

import (
	"context"
	"fmt"
	"net/http"

	"github.com/ONSdigital/go-ns/common"
)

var (
	hostRequiredButEmptyError = Error{
		Status:  500,
		Message: "error creating get permissions request host required but was empty",
	}

	noUserOrServiceAuthTokenProvidedError = Error{
		Message: "invalid request require user or service auth token but none provide",
		Status:  400,
	}

	callerUnauthorisedError = Error{
		Message: "access denied caller does not have the required permissions to perform this action",
		Status:  403,
	}

	userDatasetPermissionsURL    = "%s/userDatasetPermissions?dataset_id=%s&collection_id=%s"
	serviceDatasetPermissionsURL = "%s/serviceDatasetPermissions?dataset_id=%s"
)

type Error struct {
	Status  int
	Message string
	Cause   error
}

// Permissions is a struct
type Permissions struct {
	Create bool
	Read   bool
	Update bool
	Delete bool
}

// Parameters is a model encapsulating details about the authorisation request.
type Parameters struct {
	UserToken    string
	ServiceToken string
	CollectionID string
	DatasetID    string
}

type HTTPClienter interface {
	Do(ctx context.Context, req *http.Request) (*http.Response, error)
}

type PermissionsClient struct {
	host string
	cli  HTTPClienter
}

// newUserParameters is a constructor function for creating a new Parameters object for a user auth request.
func newUserParameters(userToken string, collectionID string, datasetID string) *Parameters {
	return &Parameters{
		UserToken:    userToken,
		CollectionID: collectionID,
		DatasetID:    datasetID,
	}
}

// newUserParameters is a constructor function for creating a new Parameters object for a service account auth request.
func newServiceParameters(serviceToken string, datasetID string) *Parameters {
	return &Parameters{
		ServiceToken: serviceToken,
		DatasetID:    datasetID,
	}
}

func (e Error) Error() string {
	if e.Cause != nil {
		return e.Cause.Error()
	}
	return e.Message
}

// createUserDatasetPermissionsRequest create a new get user dataset permissions HTTP request.
func (params *Parameters) createUserDatasetPermissionsRequest(host string) (*http.Request, error) {
	if host == "" {
		return nil, hostRequiredButEmptyError
	}

	url := fmt.Sprintf(userDatasetPermissionsURL, host, params.DatasetID, params.CollectionID)
	httpRequest, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, Error{
			Cause:   err,
			Status:  500,
			Message: "error creating new get permissions http request",
		}
	}

	httpRequest.Header.Set(common.FlorenceHeaderKey, params.UserToken)
	return httpRequest, nil
}

// createServiceDatasetPermissionsRequest create a new get service account dataset permissions HTTP request.
func (params *Parameters) createServiceDatasetPermissionsRequest(host string) (*http.Request, error) {
	if host == "" {
		return nil, hostRequiredButEmptyError
	}

	url := fmt.Sprintf(serviceDatasetPermissionsURL, host, params.DatasetID)
	r, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, Error{
			Cause:   err,
			Status:  500,
			Message: "error making get permissions http request",
		}
	}

	r.Header.Set(common.AuthHeaderKey, params.ServiceToken)
	return r, nil
}
