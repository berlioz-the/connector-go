package berlioz

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
)

func (x NativeResourceAccessor) getSession(peer CloudResourceModel) (*session.Session, error) {
	credentials := credentials.NewStaticCredentials(peer.Config.Credentials.AccessKeyID,
		peer.Config.Credentials.SecretAccessKey,
		"")
	sess, err := session.NewSession(&aws.Config{
		Region:      aws.String(peer.Config.Region),
		Credentials: credentials,
	})
	return sess, err
}
