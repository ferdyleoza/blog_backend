package config

var AllowedOrigins = []string{
	"http://localhost:6969",
	"http://localhost:5174",
}

func GetAllowedOrigins() []string {
	return AllowedOrigins
}