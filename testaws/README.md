# AWS test backend

This package implements creating an AWS test backend using [LocalStack](https://www.localstack.cloud/) and Docker. To use it, you will need a local Docker daemon running. You can use the backend as follows:

```go
package your_test

import (
	"bytes"
	"context"
	"testing"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/opentofu/tofutestutils/testaws"
)

func TestMyApp(t *testing.T) {
	awsBackend := testaws.New(t)

	s3Connection := s3.NewFromConfig(awsBackend.Config(), func(options *s3.Options) {
		options.UsePathStyle = awsBackend.S3UsePathStyle()
	})

	if _, err := s3Connection.PutObject(
		context.TODO(),
		&s3.PutObjectInput{
			Key:    aws.String("test.txt"),
			Body:   bytes.NewReader([]byte("Hello world!")),
			Bucket: aws.String(awsBackend.S3Bucket()),
		},
	); err != nil {
		t.Fatalf("❌ Failed to put object (%v)", err)
	}
}
```

> [!WARNING]
> Always use the [test context](../testcontext) for timeouts so the backend has time to clean up the test container.