package config

type AppConfig struct {
	Server Server 
	Logger Logger 
}

type Server struct {
	Port string
}

type Logger struct {
	Type string
}