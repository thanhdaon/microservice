package transport

import (
	"context"
	"encoding/json"
	"net/http"
	"os"
	"stringsvc/pkg/endpoint"
	"stringsvc/pkg/service"

	"github.com/go-kit/kit/log"
	httptransport "github.com/go-kit/kit/transport/http"
)

func SetupHTTPServer() {
	logger := log.NewLogfmtLogger(os.Stderr)
	svc := service.New()
	endpoints := endpoint.New(svc, logger)

	uppercaseHandler := httptransport.NewServer(
		endpoints.Uppercase,
		decodeUppercaseRequest,
		encodeResponse,
	)

	countHandler := httptransport.NewServer(
		endpoints.Count,
		decodeCountRequest,
		encodeResponse,
	)

	http.Handle("/uppercase", uppercaseHandler)
	http.Handle("/count", countHandler)

	logger.Log("msg", "HTTP", "addr", ":8080")
	logger.Log("err", http.ListenAndServe(":8080", nil))
}

func decodeUppercaseRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	var request endpoint.UppercaseRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

func decodeCountRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	var request endpoint.CountRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

func encodeResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}
