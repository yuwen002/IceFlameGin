package config

type DatabaseConfig struct {
	Host     string
	Port     int
	Username string
	Password string
	DBName   string
}

var DBConfig = DatabaseConfig{
	Host:     "localhost",
	Port:     3306,
	Username: "root",
	Password: "password",
	DBName:   "myapp",
}
