package main

import (
	"errors"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/gonyyi/afmt"
	"time"
)

func splitAwsPath(path string) (bucket string, filename string) {
	out := afmt.Extract(path, "^s3://([^/]+)/(.+)$", 2)
	return out[0], out[1]
}

func awsRm(bucket, region, filename string) (int64, error) {
	if filename == "" {
		return -1, errors.New("missing filename")
	}
	sess, err := session.NewSession(&aws.Config{
		Region: &region},
	)
	if err != nil {
		return -1, err
	}
	svc := s3.New(sess)

	obj, err := svc.GetObject(&s3.GetObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(filename),
	})
	if err != nil {
		return -1, err
	}

	// If size of object returned, delete it.
	objSize := aws.Int64Value(obj.ContentLength)
	if objSize < 1 {
		return -1, errors.New("obj not exist")
	}

	// Create S3 service client
	_, err = svc.DeleteObject(&s3.DeleteObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(filename),
	})
	if err != nil {
		return -1, err
	}

	return objSize, nil
}

func awsSize(bucket, region, filename string) (size int64, err error) {
	if filename == "" {
		return -1, errors.New("missing filename")
	}
	sess, err := session.NewSession(&aws.Config{
		Region: &region},
	)
	if err != nil {
		return -1, err
	}

	obj, err := s3.New(sess).GetObject(&s3.GetObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(filename),
	})
	if err != nil {
		return -1, err
	}

	return aws.Int64Value(obj.ContentLength), nil
}

type awsFiles struct {
	Name         string
	LastModified time.Time
	Size         int64
	StorageClass string
}

func (f awsFiles) String() string {
	return fmt.Sprintf("File: <%s>\n\t%12s: %s\n\t%12s: %d\n\t%12s: %s\n",
		f.Name,
		"LastModified", f.LastModified.String(),
		"Size", f.Size,
		"StorageClass", f.StorageClass)
}

func awsLs(bucket, region, prefix string) ([]awsFiles, error) {
	sess, err := session.NewSession(&aws.Config{
		Region: &region},
	)

	// Create S3 service client
	svc := s3.New(sess)

	resp, err := svc.ListObjectsV2(&s3.ListObjectsV2Input{Bucket: aws.String(bucket), Prefix: &prefix})
	if err != nil {
		return nil, err
	}

	var out []awsFiles

	for _, item := range resp.Contents {
		out = append(out, awsFiles{
			Name:         *item.Key,
			LastModified: *item.LastModified,
			Size:         *item.Size,
			StorageClass: *item.StorageClass,
		})
	}
	return out, nil
}
