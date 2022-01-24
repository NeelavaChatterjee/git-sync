package utilities

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/google/go-github/v42/github"
	"github.com/joho/godotenv"
	"golang.org/x/oauth2"
)

// contains methods and functions required to operate with the github apis
// and getting the stuff like authentication done
// that might be too big for a single controller to handle.

type TokenSource struct {
	AccessToken string
}

func (t *TokenSource) Token() (*oauth2.Token, error) {
	token := &oauth2.Token{
		AccessToken: t.AccessToken,
	}
	return token, nil
}

func init() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	ctx := context.Background()
	tokenSource := &TokenSource{
		AccessToken: os.Getenv("GITHUB_PAT"),
	}
	oauthClient := oauth2.NewClient(ctx, tokenSource)
	client := github.NewClient(oauthClient)

	commitInfo, _, err := client.Repositories.ListCommits(ctx, "NeelavaChatterjee", "mytracker", nil)
	if err != nil {
		fmt.Printf("Problem in commit information %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("%+v\n", commitInfo[0])
}

func GetRepositories() {
	fmt.Println("repos")
}
