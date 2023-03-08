package main

import (
	"CLPSS/server"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

var createServer = server.CreateServer

func main() {
	//java -Djava.library.path=./DynamoDBLocal_lib -jar DynamoDBLocal.jar -sharedDb
	// Place for the config files
	//configFilePath := "./resources/"

	// Load Configs
	//config.Config.SetConfigs(configFilePath, configFilePath)

	//psConfig.LoadKubeSecrets()

	// Add local secret files
	//_ = config.Config.MergeConfig(".env")

	// Init Variables
	//handler.InitGlobalVariables(viper.GetViper())

	// Generate Server
	router := createServer(viper.GetViper())

	logrus.Info("Listening on 8080")
	_ = router.Run(":8080")

	//Init

}
