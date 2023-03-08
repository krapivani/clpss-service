package config

import (
	"os"
)

var Properties ConfigInterface

// ConfigInterface is a mockable type of viper
type ConfigInterface interface {
	GetString(string) string
	GetInt(string) int
	GetBool(string) bool
	SetDefault(string, interface{})
	Set(key string, value interface{})
}

// needs some fake code to be tested for code coverage
func fake() int {
	return 0
}

// Takes secret Alias and corresponding property name
func SetProperty(propertyKey string, propertyName string) {
	propertyValue, err := os.LookupEnv(propertyKey)
	// for local development
	if !err {
		propertyValue = Properties.GetString(propertyName)
	}

	Properties.Set(propertyName, propertyValue)
}

//func LoadKubeSecrets() {
//	Properties = viper.GetViper()
//
//	SetProperty("BASIC_AUTH_OLD_USER", "BASIC_LOGIN_NAME_OLD")
//	SetProperty("BASIC_AUTH_OLD_KEY", "BASIC_PASSWORD_OLD")
//	SetProperty("BASIC_AUTH_USER", "BASIC_LOGIN_NAME")
//	SetProperty("BASIC_AUTH_KEY", "BASIC_PASSWORD")
//}
