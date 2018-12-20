package dynamo

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	db "github.com/guregu/dynamo"
)

// New creates a new DynamoDB client
func New() (*db.DB, error) {
	sess, err := session.NewSession(&aws.Config{
		Endpoint: aws.String("http://localhost:4569"),
		Region:   aws.String("us-west-2"),
	})
	if err != nil {
		return nil, err
	}
	return db.New(sess), nil
}
