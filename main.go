package main

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

// CreateSessionForLocalstack Create a session for localstack
func CreateSessionForLocalstack(region string, endpoint string) *session.Session {
	conf := aws.Config{
		Region:           aws.String(region),
		Endpoint:         aws.String(endpoint),
		S3ForcePathStyle: aws.Bool(true),
	}
	sess, _ := session.NewSessionWithOptions(session.Options{
		Profile:           "localstack",
		Config:            conf,
		SharedConfigState: session.SharedConfigEnable,
	})
	return sess
}

func main() {
	region := "ap-northeast-1"
	endpoint := "http://localstack:4572"
	bucket := "sample"

	sess := CreateSessionForLocalstack(
		region,
		endpoint,
	)

	svc := s3.New(sess)

	svc.CreateBucket(&s3.CreateBucketInput{
		Bucket: aws.String(bucket),
	})

	result, _ := svc.ListBuckets(nil)
	fmt.Println("Buckets:")

	for _, b := range result.Buckets {
		fmt.Printf("* %s created on %s\n", aws.StringValue(b.Name), aws.TimeValue(b.CreationDate))
	}
}
