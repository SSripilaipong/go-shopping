package http

import (
	"go-shopping/lambler"
	"net/http"
)

type Error interface {
	Error() string
	ToResponse() LambdaResponseBuilder
}

type requestError struct {
	message        string
	httpStatusCode int
}

func (r *requestError) Error() string {
	return r.message
}

func (r *requestError) ToResponse() LambdaResponseBuilder {
	return JsonResponse{
		StatusCode: r.httpStatusCode,
		Body:       lambler.Json{"message": r.message},
	}
}

func UnprocessableEntityResponse() Error {
	return &requestError{httpStatusCode: http.StatusUnprocessableEntity, message: "unprocessable entity"}
}

func BadRequestResponse() Error {
	return &requestError{httpStatusCode: http.StatusBadRequest, message: "bad request"}
}
