package server

type JobTemplateData struct {
	Name     string            `json:"name"`
	Metadata map[string]string `yaml:"metadata" json:"metadata"`
}

type Response struct {
	Status string      `json:"status"`
	Data   interface{} `json:"data"`
}
