package objectFile

import (
	"baize/app/datasource/objectFile/awsS3Object"
	"baize/app/datasource/objectFile/localhostObject"
	"context"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/spf13/viper"
	"mime/multipart"
)

const (
	awsS3     = "s3"
	yiDong    = "eos"
	localhost = "localhost"
)

type ObjectFile interface {
	PublicUploadFile(ctx context.Context, file multipart.File, keyName string) (string, error)
	PrivateUploadFile(ctx context.Context, file multipart.File, keyName string) (string, error)
}

func NewConfig() (of ObjectFile) {
	type UploadFile struct {
		Type       string `mapstructure:"type"`
		DomainName string `mapstructure:"domain_name"`
	}
	var uf UploadFile
	if err := viper.UnmarshalKey("upload_file", &uf); err != nil {
		panic(err)
	}
	switch uf.Type {
	case awsS3:
		type S3 struct {
			Region            string `mapstructure:"region"`
			PublicBucketName  string `mapstructure:"public_bucket_name"`
			PrivateBucketName string `mapstructure:"private_bucket_name"`
		}
		var s3s S3
		if err := viper.UnmarshalKey("upload_file.s3", &s3s); err != nil {
			panic(err)
		}
		config := aws.Config{
			//Credentials: credentials.NewStaticCredentialsProvider(, ""),
			Region: s3s.Region,
		}
		s := new(awsS3Object.S3File)
		s.S3Config = s3.NewFromConfig(config)
		s.PublicBucket = s3s.PublicBucketName
		s.PrivateBucket = s3s.PrivateBucketName
		s.DomainName = uf.DomainName
		of = s
	case yiDong:
		type Eos struct {
			AccessKeyId       string `mapstructure:"access_key_id"`
			SecretAccessKey   string `mapstructure:"Secret_access_key"`
			Url               string `mapstructure:"url"`
			PublicBucketName  string `mapstructure:"public_bucket_name"`
			PrivateBucketName string `mapstructure:"private_bucket_name"`
		}
		var eos Eos
		if err := viper.UnmarshalKey("upload_file.eos", &eos); err != nil {
			panic(err)
		}
		config := aws.Config{
			Credentials: credentials.NewStaticCredentialsProvider(eos.AccessKeyId, eos.SecretAccessKey, ""),
			EndpointResolverWithOptions: aws.EndpointResolverWithOptionsFunc(func(service, region string, options ...interface{}) (aws.Endpoint, error) {
				return aws.Endpoint{
					PartitionID: "aws",
					URL:         eos.Url,
				}, nil
			}),
		}
		s := new(awsS3Object.S3File)
		s.S3Config = s3.NewFromConfig(config)
		s.PublicBucket = eos.PublicBucketName
		s.PrivateBucket = eos.PrivateBucketName
		s.DomainName = uf.DomainName
		of = s
	default:
		type Localhost struct {
			PublicPath  string `mapstructure:"public_path"`
			PrivatePath string `mapstructure:"private_path"`
		}
		var lh Localhost
		if err := viper.UnmarshalKey("upload_file.localhost", &lh); err != nil {
			panic(err)
		}
		l := new(localhostObject.LocalHostFile)
		l.DomainName = uf.DomainName
		l.PublicPath = lh.PublicPath
		l.PrivatePath = lh.PrivatePath
		of = l
	}
	return of
}
