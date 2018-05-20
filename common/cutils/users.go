package cutils

import (
	"encoding/json"
	"errors"

	"github.com/aws/aws-lambda-go/events"

	"github.com/TsvetanMilanov/go-lambda-workflow/workflow"

	"github.com/TsvetanMilanov/tasker-common/common/cconstants"
	"github.com/TsvetanMilanov/tasker-common/common/ctypes"
)

// GetAuthorizerUserFromContext returns the user passed from the custom authorizer in the request
// context. If no user is found in the context, the method will return error.
func GetAuthorizerUserFromContext(ctx workflow.Context) (*ctypes.Auth0UserInfoResponse, error) {
	evt := new(events.APIGatewayProxyRequest)
	err := ctx.GetLambdaEvent(evt)
	if err != nil {
		return nil, err
	}

	userInfoValue, ok := evt.RequestContext.Authorizer[cconstants.UserInfoContextKey]
	if !ok {
		return nil, errors.New("unable to find user in the authorizer context")
	}

	res := new(ctypes.Auth0UserInfoResponse)
	unmarshalErr := json.Unmarshal([]byte(userInfoValue.(string)), res)
	if unmarshalErr != nil {
		return nil, unmarshalErr
	}

	return res, nil
}
