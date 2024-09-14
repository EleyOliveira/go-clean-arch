package configs

import "github.com/spf13/viper"

var cfg *conf

type conf struct {
	DBDriver             string `mapstructure:"DB_DRIVER"`
	DBHost               string `mapstructure:"DB_HOST"`
	DBPort               string `mapstructure:"DB_PORT"`
	DBUser               string `mapstructure:"DB_USER"`
	DBPassword           string `mapstructure:"DB_PASSWORD"`
	DBName               string `mapstructure:"DB_NAME"`
	WebServerPort        string `mapstructure:"WEB_SERVER_PORT"`
	GRPCServerPort       string `mapstructure:"GRPC_SERVER_PORT"`
	GraphQLServerPort    string `mapstructure:"GRAPHQL_SERVER_PORT"`
	RABBITMQ_USER        string `mapstructure:"RABBITMQ_USER"`
	RABBITMQ_PASSWORD    string `mapstructure:"RABBITMQ_PASSWORD"`
	RABBITMQ_HOST        string `mapstructure:"RABBITMQ_HOST"`
	RABBITMQ_SERVER_PORT string `mapstructure:"RABBITMQ_SERVER_PORT"`
}

func LoadConfig(path string) (*conf, error) {
	viper.SetConfigName("app_config")
	viper.SetConfigType("env")
	viper.AddConfigPath(path)
	viper.SetConfigFile(".env")
	viper.AutomaticEnv()
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
	err = viper.Unmarshal(&cfg)
	return cfg, err
}
