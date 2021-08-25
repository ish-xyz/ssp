package config

type Config struct {
	ServerAddr       string
	ServerPort       int
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
