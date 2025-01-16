package awsS3Object

import (
	"baize/app/utils/fileUtils"
	"context"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
	"mime/multipart"
)

type S3File struct {
	S3Config      *s3.Client
	PublicBucket  string
	PrivateBucket string
	DomainName    string
}

func (s *S3File) PublicUploadFile(ctx context.Context, file multipart.File, keyName string) (string, error) {
	ct := fileUtils.GetFileContentType(file)
	obj := &s3.PutObjectInput{
		Bucket:      aws.String(s.PublicBucket),
		Key:         aws.String(keyName),
		Body:        file,
		ContentType: aws.String(ct),
		ACL:         types.ObjectCannedACLPublicRead,
	}
	_, err := s.S3Config.PutObject(ctx, obj)
	if err != nil {
		return "", err
	}
	return s.DomainName + "/" + keyName, nil
}

func (s *S3File) PrivateUploadFile(ctx context.Context, file multipart.File, keyName string) (string, error) {
	ct := fileUtils.GetFileContentType(file)
	obj := &s3.PutObjectInput{
		Bucket:      aws.String(s.PrivateBucket),
		Key:         aws.String(keyName),
		Body:        file,
		ContentType: aws.String(ct),
		ACL:         types.ObjectCannedACLPrivate,
	}
	_, err := s.S3Config.PutObject(ctx, obj)
	if err != nil {
		return "", err
	}
	return keyName, nil
}
