package config

import "os"

type Config struct {
	Host   string
	Port   string
	DBHost string
	DBPort string
	DBUser string
	DBPass string
}

const (
	defaultHost   = `localhost`
	defaultPort   = `8080`
	defaultDBHost = `localhost`
	defaultDBPort = `5432`
	defaultDBUser = `blog_user`
	defaultDBPass = `devpass`
)

var (
	AppConfig Config
)

func Bootstrap() {

	if host := os.Getenv("GB_HOST"); host != `` {
		AppConfig.Host = host
	} else {
		AppConfig.Host = defaultHost
	}

	if port := os.Getenv("GB_PORT"); port != `` {
		AppConfig.Port = port
	} else {
		AppConfig.Port = defaultPort
	}

	if dbHost := os.Getenv("GB_DB_HOST"); dbHost != `` {
		AppConfig.DBHost = dbHost
	} else {
		AppConfig.DBHost = defaultDBHost
	}

	if dbPort := os.Getenv("GB_DB_PORT"); dbPort != `` {
		AppConfig.DBPort = dbPort
	} else {
		AppConfig.DBPort = defaultDBPort
	}

	if dbUser := os.Getenv("GB_DB_USER"); dbUser != `` {
		AppConfig.DBUser = dbUser
	} else {
		AppConfig.DBUser = defaultDBUser
	}

	if dbPass := os.Getenv("GB_DB_PASS"); dbPass != `` {
		AppConfig.DBPass = dbPass
	} else {
		AppConfig.DBPass = defaultDBPass
	}
}
