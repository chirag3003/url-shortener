package aws

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	awsConfig "github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/chirag3003/go-backend-template/config"
)

var cfg aws.Config

// SetupAWS initializes the AWS SDK configuration.
func SetupAWS(appCfg *config.Config) error {
	c, err := awsConfig.LoadDefaultConfig(context.TODO(),
		awsConfig.WithCredentialsProvider(
			credentials.NewStaticCredentialsProvider(appCfg.S3AccessKey, appCfg.S3SecretKey, ""),
		),
		awsConfig.WithRegion(appCfg.S3Region),
	)
	if err != nil {
		return fmt.Errorf("loading AWS config: %w", err)
	}

	cfg = c
	bucket = appCfg.S3Bucket
	setupS3()
	return nil
}
