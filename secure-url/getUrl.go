package main

import (
	"os"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/endpoints"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

func getUrl(filename string) (string, error) {

	sess := session.Must(session.NewSession(&aws.Config{
		Region: aws.String(endpoints.EuWest3RegionID),
	}))

	bucket := os.Getenv("BUCKET_NAME")

	svc := s3.New(sess)
	key := filename

	req, _ := svc.PutObjectRequest(&s3.PutObjectInput{
		Bucket:      aws.String(bucket),
		Key:         aws.String(key),
		ContentType: aws.String("audio/mpeg"),
	})

	urlStr, err := req.Presign(15 * time.Minute)

	if err != nil {
		return "", err
	}

	return urlStr, nil

}
