package config

type DatabaseConfig struct {
	Host     string
	Port     int
	Username string
	Password string
	DBName   string
}

var DBConfig = DatabaseConfig{
	Host:     "82.157.248.230",
	Port:     3306,
	Username: "go_test",
	Password: "XLx8EhRTEkF5fLLr111111",
	DBName:   "go_test",
}
