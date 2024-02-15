package setting

type UploadFile struct {
	Type       string `mapstructure:"type"`
	DomainName string `mapstructure:"domain_name"`
	*S3        `mapstructure:"s3"`
	*Eos       `mapstructure:"eos"`
	*Localhost `mapstructure:"localhost"`
}
type S3 struct {
	AccessKeyId       string `mapstructure:"access_key_id"`
	SecretAccessKey   string `mapstructure:"Secret_access_key"`
	Region            string `mapstructure:"region"`
	PublicBucketName  string `mapstructure:"public_bucket_name"`
	PrivateBucketName string `mapstructure:"private_bucket_name"`
}
type Eos struct {
	AccessKeyId       string `mapstructure:"access_key_id"`
	SecretAccessKey   string `mapstructure:"Secret_access_key"`
	Url               string `mapstructure:"url"`
	PublicBucketName  string `mapstructure:"public_bucket_name"`
	PrivateBucketName string `mapstructure:"private_bucket_name"`
}
type Localhost struct {
	PublicPath  string `mapstructure:"public_path"`
	PrivatePath string `mapstructure:"private_path"`
}
