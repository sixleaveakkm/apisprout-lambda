package main

import (
	"errors"
	"io/ioutil"
	"net/url"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

func loadSwaggerFromS3(bucket string) ([]byte, error) {
	sess, err := session.NewSession()
	if err != nil {
		return nil, err
	}

	s3Svc := s3.New(sess)
	if !strings.HasPrefix(bucket, "s3://") {
		return nil, errors.New("invalid bucket name, missing prefix")
	}

	s3Uri, err := url.Parse(bucket)
	if err != nil {
		return nil, err
	}
	s3ObjectData, err := s3Svc.GetObject(&s3.GetObjectInput{
		Bucket: aws.String(s3Uri.Host),
		Key:    aws.String(s3Uri.Path),
	})

	if err != nil {
		return nil, err
	}

	return ioutil.ReadAll(s3ObjectData.Body)
}
