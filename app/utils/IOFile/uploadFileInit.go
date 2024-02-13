package IOFile

import (
	"baize/app/setting"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"io"
)

type fileParams struct {
	keyName     string
	contentType string
	data        io.Reader
}

const (
	awsS3     = "s3"
	yiDong    = "eos"
	localhost = "localhost"
)

type IOFile interface {
	PublicUploadFile(file *fileParams) (string, error)
	privateUploadFile(file *fileParams) (string, error)
}

var ioFile IOFile

func GetConfig() IOFile {
	return ioFile
}

func init() {
	switch setting.Conf.UploadFile.Type {
	case awsS3:
		config := aws.Config{
			Credentials: credentials.NewStaticCredentialsProvider(setting.Conf.UploadFile.S3.AccessKeyId, setting.Conf.UploadFile.S3.SecretAccessKey, ""),
			Region:      setting.Conf.UploadFile.S3.Region,
		}
		s := new(s3IOFile)
		s.s3Config = s3.NewFromConfig(config)
		s.bucket = setting.Conf.UploadFile.S3.BucketName
		s.domainName = setting.Conf.UploadFile.DomainName
		ioFile = s
	case yiDong:
		config := aws.Config{
			Credentials: credentials.NewStaticCredentialsProvider(setting.Conf.UploadFile.Eos.AccessKeyId, setting.Conf.UploadFile.Eos.SecretAccessKey, ""),
			EndpointResolverWithOptions: aws.EndpointResolverWithOptionsFunc(func(service, region string, options ...interface{}) (aws.Endpoint, error) {
				return aws.Endpoint{
					PartitionID: "aws",
					URL:         setting.Conf.UploadFile.Eos.Url,
				}, nil
			}),
		}
		s := new(s3IOFile)
		s.s3Config = s3.NewFromConfig(config)
		s.bucket = setting.Conf.UploadFile.Eos.BucketName
		s.domainName = setting.Conf.UploadFile.DomainName
		ioFile = s
	default:
		l := new(localHostIOFile)
		l.domainName = setting.Conf.UploadFile.DomainName
		pubPath := setting.Conf.UploadFile.Localhost.PublicPath
		l.publicPath = pubPath
		priPath := setting.Conf.UploadFile.Localhost.PrivatePath
		l.privatePath = priPath
		ioFile = l
	}
}
