package create_new_item

import (
	"go-shopping/lambler"
	lTesting "go-shopping/lambler/testing"
	"go-shopping/tests"
	"math/rand"
)

const Url = "/inventory/v1/items"

func doValidRequest(handler lambler.Handler) (any, error) {
	payload := validPayload(tests.RandomString(rand.Intn(16)), tests.RandomString(rand.Intn(64)), rand.Int())
	return doValidRequestWithPayload(handler, payload)
}

func doValidRequestWithPayload(handler lambler.Handler, payload lambler.Json) (any, error) {
	return handler(lTesting.NewContext(), lTesting.HttpPostEventWithJsonBody(Url, payload))
}

func validPayload(name string, description string, quantity int) lambler.Json {
	return lambler.Json{
		"name":        name,
		"description": description,
		"quantity":    quantity,
	}
}
