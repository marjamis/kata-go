package main

import (
	"fmt"
	"os"
	"time"

	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/credentials/endpointcreds"
	"github.com/aws/aws-sdk-go/aws/defaults"
	"github.com/aws/aws-sdk-go/aws/session"
)

func credDetails(typeOfCheck string, creds *credentials.Credentials) {
	credValue, err := creds.Get()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Printf("%s: %+v\n", typeOfCheck, credValue.AccessKeyID)
}

func customOrder() {
	ds := defaults.Get()
	creds := credentials.NewChainCredentials(
		[]credentials.Provider{
			defaults.RemoteCredProvider(*ds.Config, ds.Handlers),
			&credentials.EnvProvider{},
		})
	credDetails("customOrder", creds)
}

func onlyTaskRoles() {
	const (
		ECSContainerCredentialsURI         = "http://169.254.170.2"
		HttpProviderAuthorizationEnvVar    = "AWS_CONTAINER_AUTHORIZATION_TOKEN"
		ECSContainerCredentialsRelativeURI = "AWS_CONTAINER_CREDENTIALS_RELATIVE_URI"
	)

	credsEndpoint := fmt.Sprintf("%s%s", ECSContainerCredentialsURI,
		os.Getenv(ECSContainerCredentialsRelativeURI))

	ds := defaults.Get()

	// credsProvider :=
	creds := credentials.NewChainCredentials(
		[]credentials.Provider{
			endpointcreds.NewProviderClient(*ds.Config, ds.Handlers, credsEndpoint,
				func(p *endpointcreds.Provider) {
					p.ExpiryWindow = 5 * time.Minute
					p.AuthorizationToken = os.Getenv(HttpProviderAuthorizationEnvVar)
				}),
		})

	// creds := credentials.NewCredentials(credsProvider)
	credDetails("onlyTaskRoles", creds)
}

func normal() {
	sess, _ := session.NewSession()
	credValue, err := sess.Config.Credentials.Get()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Printf("Credentials: %+v\n", credValue.AccessKeyID)
}

func main() {
	normal()
	customOrder()
	onlyTaskRoles()
}
