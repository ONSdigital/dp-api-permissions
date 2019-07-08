package authv2

import (
	"context"
	"net/http"

	"github.com/ONSdigital/go-ns/common"
	"github.com/ONSdigital/log.go/log"
)

func RequireReadDatasetPermission(endpoint http.HandlerFunc) http.HandlerFunc {
	return RequireDatasetPermissions(&Permissions{Read: true}, endpoint)
}

func RequireDatasetPermissions(requiredPermissions *Permissions, wrappedHandler http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		logD := log.Data{"requested_uri": req.URL.RequestURI()}

		parameters, err := extractAuthorisationParameters(req)
		if err != nil {
			handleAuthoriseError(req.Context(), err, w, logD)
			return
		}

		callerPermissions, err := permissionsCli.GetCallerPermissions(parameters)
		if err != nil {
			handleAuthoriseError(req.Context(), err, w, logD)
			return
		}

		err = permissionsVerifier.CheckPermissionsRequirementsSatisfied(callerPermissions, requiredPermissions)
		if err != nil {
			handleAuthoriseError(req.Context(), err, w, logD)
			return
		}

		log.Event(req.Context(), "caller authorised to perform requested action", logD)
		wrappedHandler(w, req)
	})
}

func extractAuthorisationParameters(req *http.Request) (*Parameters, error) {
	userAuthToken := req.Header.Get(common.FlorenceHeaderKey)
	serviceAuthToken := req.Header.Get(common.AuthHeaderKey)
	collectionID := req.Header.Get(CollectionIDHeader)
	datasetID := getRequestVars(req)[datasetIDKey]

	if userAuthToken != "" {
		return newUserParameters(userAuthToken, collectionID, datasetID), nil
	}

	if serviceAuthToken != "" {
		return newServiceParameters(serviceAuthToken, datasetID), nil
	}

	return nil, noUserOrServiceAuthTokenProvidedError
}

func handleAuthoriseError(ctx context.Context, err error, w http.ResponseWriter, logD log.Data) {
	permErr, ok := err.(Error)
	if ok {
		writeErr(ctx, w, permErr.Status, permErr.Message, logD)
		return
	}
	writeErr(ctx, w, 500, "internal server error", logD)
}

func writeErr(ctx context.Context, w http.ResponseWriter, status int, body string, logD log.Data) {
	w.WriteHeader(status)
	b := []byte(body)
	_, wErr := w.Write(b)
	if wErr != nil {
		w.WriteHeader(500)
		logD["original_err_body"] = body
		logD["original_err_status"] = status

		log.Event(ctx, "internal server error failed writing permissions error to response", log.Error(wErr), logD)
		return
	}
}
