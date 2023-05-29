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

	country := c.Param("country")

	//if country != "USA" && country != "India" {
	//	logrus.Error("Failed to get the users since Country is wrong: ", country)
	//	c.JSON(http.StatusInternalServerError, "Invalid Request")
	//}

	params := &dynamodb.QueryInput{
		TableName:              aws.String("user_data"),
		KeyConditionExpression: aws.String("#country = :countryValue"),
		ExpressionAttributeNames: map[string]*string{
			"#country": aws.String("country"),
		},
		ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
			":countryValue": {
				S: aws.String(country),
			},
		},
	}
	result, err := dynamoDB.Query(params)
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
