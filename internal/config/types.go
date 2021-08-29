package config

type Config struct {
	BackendAddr      string
	FrontendAddr     string
	KubeConfigPath   string
	JobTemplatesPath string
	AllowedUsers     []string
	GitSync          struct {
		Repository string
		Protocol   string
		Username   string
		Password   string
	}
}
