package common

import (
	"net/http"

	"github.com/TsvetanMilanov/go-lambda-workflow/workflow"
	"github.com/go-playground/validator"
)

// ValidateRequest validates the request object in the context and sets the
// response to bad request if there are any errors.
func ValidateRequest(ctx workflow.Context) error {
	req := ctx.GetRequest()
	vErrs := ValidateStruct(req)
	if len(vErrs) > 0 {
		resBody := []string{}
		for _, e := range vErrs {
			resBody = append(resBody, e.Translate(nil))
		}

		ctx.SetResponse(resBody).SetResponseStatusCode(http.StatusBadRequest)
	}

	return nil
}

func ValidateStruct(item interface{}) validator.ValidationErrors {
	validate := validator.New()
	errs := validate.Struct(item)
	if errs != nil {
		return errs.(validator.ValidationErrors)
	}

	return nil
}
