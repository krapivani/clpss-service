package handler

import (
	"CLPSS/dbConfig"
	"CLPSS/structs"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
)

var dynamoDB = dbConfig.ConnectDb()

type HealthResponse struct {
	Message string `json:"Message,omitempty"`
	Status  string `json:"Status,omitempty"`
}

func GetHealth(c *gin.Context) {
	response := HealthResponse{Status: "OK"}
	c.JSON(http.StatusOK, response)
}

func GetInfo(c *gin.Context) {
	response := HealthResponse{Message: "GetInfo.", Status: "OK"}
	c.JSON(http.StatusOK, response)
}

func GetUsers(c *gin.Context) {

	params := &dynamodb.ScanInput{
		TableName: aws.String("user_data"),
	}
	result, err := dynamoDB.Scan(params)
	if err != nil {
		logrus.Error("Error fetching data ", err)
		c.JSON(http.StatusInternalServerError, err)
	}
	var Users []structs.User
	err = dynamodbattribute.UnmarshalListOfMaps(result.Items, &Users)
	if err != nil {
		logrus.Error("Failed to unmarshal Query result items", err)
		c.JSON(http.StatusInternalServerError, err)
	}

	response := structs.HttpResponse{
		Profile: Users,
		Status:  "ok",
	}
	c.JSON(http.StatusOK, response)
}
