package model

type Config struct {
	Database struct {
		Host   string `yaml:"host" env:"HOST" env-default:"localhost"`
		Source string `yaml:"source" env:"SOURCE" env-default:"mygram_db"`
	} `yaml:"database"`
	Server struct {
		Port string `yaml:"port" env:"PORT" env-default:"8080"`
	} `yaml:"server"`
	JWT struct {
		Secret string `yaml:"secret" env:"SECRET" env-default:"G0l4ng*#f1n@lPr0j3T?"`
	} `yaml:"jwt"`
}
