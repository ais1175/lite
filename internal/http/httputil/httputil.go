package httputil

type response struct {
	Status string `json:"status"`
	Data   any    `json:"data"`
}

func Response(data any) *response {
	return &response{
		Status: "ok",
		Data:   data,
	}
}
