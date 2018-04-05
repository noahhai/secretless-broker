package provider

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/secretsmanager"
	"github.com/aws/aws-sdk-go/aws"
)

// AWSSecretsProvider provides data values from AWS Secrets Manager.
type AWSSecretsProvider struct {
	name   string
	client *secretsmanager.SecretsManager
}

// NewAWSSecretsProvider constructs a AWSSecretsProvider. The API client is configured in a similar fashion to the AWS CLI.
// See https://docs.aws.amazon.com/cli/latest/userguide/cli-chap-getting-started.html
func NewAWSSecretsProvider(name string) (provider Provider, err error) {

	// All clients require a Session. The Session provides the client with
	// shared configuration such as region, endpoint, and credentials. A
	// Session should be shared where possible to take advantage of
	// configuration and credential caching.
	sess, err := session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	})
	if err != nil {
		return
	}
	// Create a new instance of the service's client with a Session.
	svc := secretsmanager.New(sess)

	provider = AWSSecretsProvider{name: name, client: svc}
	return
}

// Name returns the name of the provider
func (p AWSSecretsProvider) Name() string {
	return p.name
}

// Value obtains a secret value by id.
func (p AWSSecretsProvider) Value(id string) (value []byte, err error) {
	client := p.client

	req, resp := client.GetSecretValueRequest(&secretsmanager.GetSecretValueInput{
		SecretId: aws.String(id),
	})

	err = req.Send()
	if err != nil { // resp is now filled
		return
	}

	if resp.SecretString != nil {
		value = []byte(*resp.SecretString)
	} else {
		value = resp.SecretBinary
	}

	return
}
