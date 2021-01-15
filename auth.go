package main

import (
	"context"
	"log"
	"os"

	// "github.com/AzureAD/microsoft-authentication-library-for-go/msal"
	msal "github.com/AzureAD/microsoft-authentication-library-for-go/apps/confidential"
)

const cogSvcScope = "https://cognitiveservices.azure.com/.default"
const clientID = "91cbc6d5-9a68-4437-848a-f17f16dbf7ea"

func getSecret() string {
	sec := os.Getenv("MSAL_SECRET")
	if sec == "" {
		log.Fatalf("MSAL_SECRET is missing")
	}
	return sec
}

type tokenProvider struct{}

func (t *tokenProvider) OAuthToken() string {

	clientCredential, err := msal.NewCredFromSecret(getSecret())

	if err != nil {
		log.Fatalf("Couldn't create client app (%s)", err)
	}
	app, err := msal.New(clientID, clientCredential, func(o *msal.Options) {
		o.Authority = "https://login.windows.net/72f988bf-86f1-41af-91ab-2d7cd011db47"
	})
	if err != nil {
		log.Fatalf("Couldn't get application (%s)", err)
	}

	ctx := context.Background()
	token, err := app.AcquireTokenByCredential(ctx, []string{
		cogSvcScope,
	})
	if err != nil {
		log.Fatalf("Error getting token (%s)", err)
	}
	return token.AccessToken
}
