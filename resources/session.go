package resources

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
)
const AWS_REGION = "ap-southeast-2"

var Session = session.New()

var Service = ec2.New(Session, &aws.Config{
	Region: aws.String(AWS_REGION),
})

