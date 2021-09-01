package server

type JobTemplate struct {
	Name     string                   `yaml:"name" json:"name"`
	Metadata map[string]string        `yaml:"metadata" json:"metadata"`
	Inputs   []map[string]interface{} `yaml:"inputs" json:"inputs"`
	Template interface{}              `yaml:"template" json:"template"`
}

type Response struct {
	Status string      `json:"status"`
	Data   interface{} `json:"data"`
}
