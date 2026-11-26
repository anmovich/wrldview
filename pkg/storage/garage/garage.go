package garage

import (
	"context"
	"io"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

type GarageS3 struct {
	client s3.Client
	bucket string
}

func NewGarage(endpoint string, accesedKey string, s3cr3tKey string, bucket string) *GarageS3 {
	client := s3.New(s3.Options{
		Region:       "garage",
		BaseEndpoint: aws.String(endpoint),
		Credentials:  credentials.NewStaticCredentialsProvider(accesedKey, s3cr3tKey, ""),
		UsePathStyle: true,
	})
	return &GarageS3{
		client: *client,
		bucket: bucket,
	}
}

func (g *GarageS3) Upload(ctx context.Context, name string, body io.Reader, contentType string) error {
	_, err := g.client.PutObject(
		ctx,
		&s3.PutObjectInput{
			Bucket:      aws.String(g.bucket),
			Key:         aws.String(name),
			Body:        body,
			ContentType: aws.String(contentType),
		},
	)
	return err
}
