package enviroment

import (
	"log"

	"github.com/spf13/viper"
)

type Enviroment struct {
	AppEnv                 string `mapstructure:"APP_ENV"`
	ServerAddress          string `mapstructure:"SERVER_ADDRESS"`
	DBHost                 string `mapstructure:"DB_HOST"`
	DBPort                 string `mapstructure:"DB_PORT"`
	DBUser                 string `mapstructure:"MYSQL_ROOT_PASSWORD"`
	DBPass                 string `mapstructure:"MYSQL_ROOT_USER"`
	DBName                 string `mapstructure:"MYSQL_DATABASE"`
	AccessTokenExpiryHour  int    `mapstructure:"ACCESS_TOKEN_EXPIRY_HOUR"`
	RefreshTokenExpiryHour int    `mapstructure:"REFRESH_TOKEN_EXPIRY_HOUR"`
	AccessTokenSecret      string `mapstructure:"ACCESS_TOKEN_SECRET"`
	RefreshTokenSecret     string `mapstructure:"REFRESH_TOKEN_SECRET"`
	SecretKey              string `mapstructure:"SECRET_KEY"`
}

func EnviromentInit() *Enviroment {
	env := Enviroment{}

	viper.SetConfigType("env")
	viper.SetConfigName(".env.development")
	viper.AddConfigPath("./secrets")

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal("Can't find the file .env.development : ", err)
		panic(err)
	}

	err = viper.Unmarshal(&env)
	if err != nil {
		log.Fatal("Environment can't be loaded: ", err)
		panic(err)
	}

	if env.AppEnv == "development" {
		log.Println("The App is running in development env")
	}

	if env.AppEnv == "production" {
		log.Println("The App is running in production env")
	}

	if env.AppEnv == "staging" {
		log.Println("The App is running in staging env")
	}

	return &env
}
