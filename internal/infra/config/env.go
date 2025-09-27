package config

import (
	"github.com/caarlos0/env/v11"
	_ "github.com/joho/godotenv/autoload"
)

type App struct {
	AppName string `env:"APP_NAME" envDefault:"Go App"`
	AppPort int    `env:"APP_PORT" envDefault:"8001"`
	AppEnv  string `env:"APP_ENV" envDefault:"development"`
}

type JWT struct {
	JWTSecret   string `env:"JWT_SECRET,required"`
	JWTDuration int    `env:"JWT_DURATION,required"`
}

type DB struct {
	DatabaseUsername string `env:"DATABASE_USERNAME,required"`
	DatabasePassword string `env:"DATABASE_PASSWORD,required"`
	DatabaseHost     string `env:"DATABASE_HOST,required"`
	DatabasePort     int    `env:"DATABASE_PORT,required"`
	DatabaseDB       string `env:"DATABASE_DB,required"`
	DatabaseSSL      string `env:"DATABASE_SSL,required"`
}

type Storage struct {
	S3AccessKey  string `env:"S3_ACCESS_KEY,required"`
	S3SecretKey  string `env:"S3_SECRET_KEY,required"`
	S3Region     string `env:"S3_REGION,required"`
	S3Endpoint   string `env:"S3_ENDPOINT,required"`
	S3BucketName string `env:"S3_BUCKET_NAME,required"`
}

type GRPC struct {
	GRPCHost string `env:"GRPC_HOST,required"`
	GRPCPort int    `env:"GRPC_PORT,required"`
}

type WEB3 struct {
	WEB3PrivateKey      string `env:"WEB3_PRIVATE_KEY,required"`
	WEB3InfuraRPC       string `env:"WEB3_INFURA_RPC,required"`
	WEB3ContractAddress string `env:"WEB3_CONTRACT_ADDRESS,required"`
}

type Env struct {
	App
	DB
	JWT
	Storage
	GRPC
	WEB3
}

func Load() (*Env, error) {
	_env := new(Env)

	if err := env.Parse(_env); err != nil {
		return nil, err
	}

	return _env, nil
}
