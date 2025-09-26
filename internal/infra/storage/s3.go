package storage

import (
	"bytes"
	"context"
	"io"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/feature/s3/manager"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
)

type S3 struct {
	client     *s3.Client
	uploader   *manager.Uploader
	bucketName string
}

type IS3 interface {
	Upload(ctx context.Context, key string, file io.ReadSeeker, contentType string) error
	Delete(ctx context.Context, key string) error
}

func New(accessKey, secretKey, region, endpoint, bucketName string) IS3 {
	client := s3.New(s3.Options{
		BaseEndpoint: aws.String(endpoint),
		Region:       region,
		Credentials: aws.NewCredentialsCache(
			credentials.NewStaticCredentialsProvider(
				accessKey,
				secretKey,
				"",
			),
		),
		EndpointResolver: s3.EndpointResolverFromURL(endpoint),
		UsePathStyle:     true,
	})

	return &S3{
		client: client,
		uploader: manager.NewUploader(client, func(u *manager.Uploader) {
			u.PartSize = 100 * 1024 * 1024 // set max 100mb
		}),
		bucketName: bucketName,
	}
}

func (storage *S3) Upload(ctx context.Context, key string, file io.ReadSeeker, contentType string) error {
	buf, err := io.ReadAll(file)
	if err != nil {
		return err
	}

	reader := bytes.NewReader(buf)

	input := &s3.PutObjectInput{
		Bucket:             aws.String(storage.bucketName),
		Key:                aws.String(key),
		Body:               reader,
		ContentLength:      aws.Int64(int64(len(buf))),
		ACL:                types.ObjectCannedACLPublicRead,
		ContentType:        aws.String(contentType),
		ContentDisposition: aws.String("inline"),
	}

	_, err = storage.uploader.Upload(ctx, input)
	if err != nil {
		return err
	}

	return nil
}

func (storage *S3) Delete(ctx context.Context, key string) error {
	input := &s3.DeleteObjectInput{
		Bucket: aws.String(storage.bucketName),
		Key:    aws.String(key),
	}

	_, err := storage.client.DeleteObject(ctx, input)
	if err != nil {
		return err
	}

	return nil
}
