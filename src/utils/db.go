package utils

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/guregu/dynamo"
)

var FurizuDB *dynamo.DB
var AwsSession *session.Session

func init() {
	AwsSession = session.Must(session.NewSession())
	FurizuDB = dynamo.New(AwsSession, &aws.Config{Region: aws.String("us-west-2")})
}
