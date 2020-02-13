package endpoint

import (
	"context"
	"stringsvc/pkg/service"

	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/log"
)

type Endpoints struct {
	Uppercase endpoint.Endpoint
	Count     endpoint.Endpoint
}

func New(svc service.StringService, logger log.Logger) Endpoints {
	var uppercase endpoint.Endpoint
	{
		uppercase = makeUppercaseEndpoint(svc)
		uppercase = loggingMiddleware(log.With(logger, "method", "uppercase"))(uppercase)
	}

	var count endpoint.Endpoint
	{
		count = makeCountEndpoint(svc)
		count = loggingMiddleware(log.With(logger, "count", "uppercase"))(count)
	}

	return Endpoints{
		Uppercase: uppercase,
		Count:     count,
	}
}

func makeUppercaseEndpoint(svc service.StringService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(UppercaseRequest)
		v, err := svc.Uppercase(req.S)
		if err != nil {
			return UppercaseResponse{v, err.Error()}, nil
		}
		return UppercaseResponse{v, ""}, nil
	}
}

func makeCountEndpoint(svc service.StringService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(CountRequest)
		return CountResponse{svc.Count(req.S)}, nil
	}
}
