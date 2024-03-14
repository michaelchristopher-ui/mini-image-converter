package conf

type Config struct {
	Server Server `yaml:"server"`
}

type Server struct {
	ReadTimeout  int    `yaml:"read_timeout"`
	WriteTimeout int    `yaml:"write_timeout"`
	Port         string `yaml:"port"`
}
