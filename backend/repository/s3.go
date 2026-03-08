package repository

import (
	"context"
	"io"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/s3/manager"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	awsHelper "github.com/chirag3003/go-backend-template/helpers/aws"
)

//go:generate mockgen -destination=mock/mock_s3.go -package=mock github.com/chirag3003/go-backend-template/repository S3Repository

// S3Repository defines the interface for S3 file operations.
type S3Repository interface {
	Upload(ctx context.Context, key string, file io.Reader) (*manager.UploadOutput, error)
}

type s3Repository struct {
	uploader *manager.Uploader
	bucket   string
}

// NewS3Repository creates a new S3Repository.
func NewS3Repository() S3Repository {
	return &s3Repository{
		uploader: awsHelper.GetS3Uploader(),
		bucket:   awsHelper.GetBucket(),
	}
}

func (r *s3Repository) Upload(ctx context.Context, key string, file io.Reader) (*manager.UploadOutput, error) {
	return r.uploader.Upload(ctx, &s3.PutObjectInput{
		Bucket: aws.String(r.bucket),
		Key:    aws.String(key),
		Body:   file,
		ACL:    "public-read",
	})
}
