package endpoint

type UppercaseRequest struct {
	S string `json:"s"`
}

type UppercaseResponse struct {
	V     string `json:"v"`
	Error string `json:"error,omitempty"`
}

type CountRequest struct {
	S string `json:"s"`
}

type CountResponse struct {
	V int `json:"v"`
}
