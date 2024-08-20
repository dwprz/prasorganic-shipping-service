package config

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"os"

	"github.com/dwprz/prasorganic-shipping-service/src/common/log"
	"github.com/sirupsen/logrus"
)

type currentApp struct {
	RestfulAddress string
	GrpcPort       string
}

type apiGateway struct {
	BaseUrl           string
	BasicAuth         string
	BasicAuthUsername string
	BasicAuthPassword string
}

type shipper struct {
	BaseUrl string
	ApiKey  string
}

type jwt struct {
	PrivateKey *rsa.PrivateKey
	PublicKey  *rsa.PublicKey
}

type redis struct {
	AddrNode1 string
	AddrNode2 string
	AddrNode3 string
	AddrNode4 string
	AddrNode5 string
	AddrNode6 string
	Password  string
}

type kafka struct {
	Addr1 string
	Addr2 string
	Addr3 string
}

type Config struct {
	CurrentApp *currentApp
	ApiGateway *apiGateway
	Shipper    *shipper
	Jwt        *jwt
	Redis      *redis
	Kafka      *kafka
}

var Conf *Config

func init() {
	appStatus := os.Getenv("PRASORGANIC_APP_STATUS")

	if appStatus == "DEVELOPMENT" {

		Conf = setUpForDevelopment()
		return
	}

	Conf = setUpForNonDevelopment(appStatus)
}

func loadRSAPrivateKey(privateKey string) *rsa.PrivateKey {
	block, _ := pem.Decode([]byte(privateKey))
	if block == nil {
		log.Logger.WithFields(logrus.Fields{
			"location": "config.loadRSAPrivateKey",
			"section":  "pem.Decode",
		}).Fatal("failed to parse pem block containing the key")
	}

	rsaPrivateKey, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		log.Logger.WithFields(logrus.Fields{
			"location": "config.loadRSAPrivateKey",
			"section":  "x509.ParsePKCS1PrivateKey",
		}).Fatalf("failed to parse rsa private key: %v", err)
	}

	return rsaPrivateKey
}

func loadRSAPublicKey(publicKey string) *rsa.PublicKey {
	block, _ := pem.Decode([]byte(publicKey))
	if block == nil {
		log.Logger.WithFields(logrus.Fields{
			"location": "config.loadRSAPrivateKey",
			"section":  "pem.Decode",
		}).Fatal("failed to parse pem block containing the key")
	}

	pub, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		log.Logger.WithFields(logrus.Fields{
			"location": "loadRSAPublicKey",
			"section":  "x509.ParsePKCS1PublicKey",
		}).Fatalf("failed to parse rsa public key: %v", err)
	}

	rsaPublicKey, ok := pub.(*rsa.PublicKey)
	if !ok {
		log.Logger.WithFields(logrus.Fields{
			"location": "loadRSAPublicKey",
			"section":  "type assertion",
		}).Fatal("failed to assert type to *rsa.PublicKey")
	}

	return rsaPublicKey
}
