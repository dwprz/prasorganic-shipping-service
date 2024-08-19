package config

import (
	"os"

	"github.com/dwprz/prasorganic-shipping-service/src/common/log"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func setUpForDevelopment() *Config {
	err := os.Chdir(os.Getenv("PRASORGANIC_SHIPPING_SERVICE_WORKSPACE"))
	if err != nil {
		log.Logger.WithFields(logrus.Fields{"location": "config.setUpForDevelopment", "section": "os.Chdir"}).Fatal(err)
	}

	viper := viper.New()
	viper.SetConfigFile(".env")
	viper.AddConfigPath(".")

	err = viper.ReadInConfig()
	if err != nil {
		log.Logger.WithFields(logrus.Fields{"location": "config.setUpForDevelopment", "section": "viper.ReadInConfig"}).Fatal(err)
	}

	currentAppConf := new(currentApp)
	currentAppConf.RestfulAddress = viper.GetString("CURRENT_APP_RESTFUL_ADDRESS")
	currentAppConf.GrpcPort = viper.GetString("CURRENT_APP_GRPC_PORT")

	apiGatewayConf := new(apiGateway)
	apiGatewayConf.BaseUrl = viper.GetString("API_GATEWAY_BASE_URL")
	apiGatewayConf.BasicAuth = viper.GetString("API_GATEWAY_BASIC_AUTH")
	apiGatewayConf.BasicAuthUsername = viper.GetString("API_GATEWAY_BASIC_AUTH_USERNAME")
	apiGatewayConf.BasicAuthPassword = viper.GetString("API_GATEWAY_BASIC_AUTH_PASSWORD")

	shipperConf := new(shipper)
	shipperConf.BaseUrl = viper.GetString("SHIPPER_BASE_URL")
	shipperConf.ApiKey = viper.GetString("SHIPPER_API_KEY")

	jwtConf := new(jwt)
	jwtConf.PrivateKey = loadRSAPrivateKey(viper.GetString("JWT_PRIVATE_KEY"))
	jwtConf.PublicKey = loadRSAPublicKey(viper.GetString("JWT_PUBLIC_KEY"))

	redisConf := new(redis)
	redisConf.AddrNode1 = viper.GetString("REDIS_ADDR_NODE_1")
	redisConf.AddrNode2 = viper.GetString("REDIS_ADDR_NODE_2")
	redisConf.AddrNode3 = viper.GetString("REDIS_ADDR_NODE_3")
	redisConf.AddrNode4 = viper.GetString("REDIS_ADDR_NODE_4")
	redisConf.AddrNode5 = viper.GetString("REDIS_ADDR_NODE_5")
	redisConf.AddrNode6 = viper.GetString("REDIS_ADDR_NODE_6")
	redisConf.Password = viper.GetString("REDIS_PASSWORD")

	kafkaConf := new(kafka)
	kafkaConf.Addr1 = viper.GetString("KAFKA_ADDRESS_1")
	kafkaConf.Addr2 = viper.GetString("KAFKA_ADDRESS_2")
	kafkaConf.Addr3 = viper.GetString("KAFKA_ADDRESS_3")

	return &Config{
		CurrentApp: currentAppConf,
		ApiGateway: apiGatewayConf,
		Shipper:    shipperConf,
		Jwt:        jwtConf,
		Redis:      redisConf,
		Kafka:      kafkaConf,
	}
}
