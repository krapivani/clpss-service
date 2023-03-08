package dbConfig

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/sirupsen/logrus"
)

func ConnectDb() *dynamodb.DynamoDB {
	sess, err := session.NewSession(&aws.Config{
		Region:      aws.String("us-east-2"),
		Credentials: credentials.NewSharedCredentials("/Users/BlackJack/.aws/config", "default"),
	})
	if err != nil {
		logrus.Error("Error connecting to DB", err)
	}
	svc := dynamodb.New(sess)
	return svc
}
