package main

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"

	"github.com/conjurinc/secretless/internal/pkg/provider"
	_ "github.com/joho/godotenv/autoload"
)

func TestAWSSecrets_Provider(t *testing.T) {
	var err error

	name := "aws-secrets"

	provider, err := provider.NewAWSSecretsProvider(name)
	if err != nil {
		t.Fatal(err)
	}

	Convey("Has the expected provider name", t, func() {
		So(provider.Name(), ShouldEqual, "aws-secrets")
	})

}
