package aws

import (
	"github.com/aws/aws-sdk-go-v2/feature/s3/manager"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

var (
	s3Client     *s3.Client
	s3Uploader   *manager.Uploader
	s3Downloader *manager.Downloader
	bucket       string
)

func setupS3() {
	client := s3.NewFromConfig(cfg)
	s3Client = client
	s3Uploader = manager.NewUploader(client)
	s3Downloader = manager.NewDownloader(client)
}

// GetS3Client returns the S3 client.
func GetS3Client() *s3.Client {
	return s3Client
}

// GetS3Uploader returns the S3 uploader.
func GetS3Uploader() *manager.Uploader {
	return s3Uploader
}

// GetS3Downloader returns the S3 downloader.
func GetS3Downloader() *manager.Downloader {
	return s3Downloader
}

// GetBucket returns the configured S3 bucket name.
func GetBucket() string {
	return bucket
}
