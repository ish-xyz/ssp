package ui

type Request struct {
	Status string      `json:"status"`
	Data   interface{} `json:"data"`
}
