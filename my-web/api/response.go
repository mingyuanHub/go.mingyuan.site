package api

type Response struct {
	Status  int                    `json:"status"`
	Message string                 `json:"message"`
	Result  map[string]interface{} `json:"result"`
}
