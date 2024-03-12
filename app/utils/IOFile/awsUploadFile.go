package IOFile

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
	"mime/multipart"
)

type s3IOFile struct {
	s3Config      *s3.Client
	publicBucket  string
	privateBucket string
	domainName    string
}

func (s *s3IOFile) PublicUploadFile(ctx context.Context, file multipart.File, keyName string) (string, error) {
	ct := GetFileContentType(file)
	obj := &s3.PutObjectInput{
		Bucket:      aws.String(s.publicBucket),
		Key:         aws.String(keyName),
		Body:        file,
		ContentType: aws.String(ct),
		ACL:         types.ObjectCannedACLPublicRead,
	}
	_, err := s.s3Config.PutObject(ctx, obj)
	if err != nil {
		return "", err
	}
	return s.domainName + "/" + keyName, nil
}

func (s *s3IOFile) PrivateUploadFile(ctx context.Context, file multipart.File, keyName string) (string, error) {
	ct := GetFileContentType(file)
	obj := &s3.PutObjectInput{
		Bucket:      aws.String(s.privateBucket),
		Key:         aws.String(keyName),
		Body:        file,
		ContentType: aws.String(ct),
		ACL:         types.ObjectCannedACLPrivate,
	}
	_, err := s.s3Config.PutObject(ctx, obj)
	if err != nil {
		return "", err
	}
	return keyName, nil
}
