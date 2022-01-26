package utilities

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/google/go-github/v42/github"
	"github.com/joho/godotenv"
	"golang.org/x/oauth2"
)

// contains methods and functions required to operate with the github apis
// and getting the stuff like authentication done
// that might be too big for a single controller to handle.

var githubClient *github.Client

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
	githubClient = github.NewClient(oauthClient)

	commitInfo, _, err := githubClient.Repositories.ListCommits(ctx, "NeelavaChatterjee", "mytracker", nil)
	if err != nil {
		fmt.Printf("Problem in commit information %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("%+v\n", commitInfo[0])
}

// Does not work becauseyou need to authenticate with installation access token
func GetRepositories() {
	repos, _, err := githubClient.Apps.ListRepos(context.Background(), nil)
	if err != nil {
		fmt.Println("Couldn't fetch the list of Repositories. ", err)
		os.Exit(1)
	}

	fmt.Println(repos.TotalCount)
	for i, repo := range repos.Repositories {
		fmt.Println(i, repo.Name)
	}
}

// Function to fetch all the branches of a repository.
// ref: https://pkg.go.dev/github.com/google/go-github/v42@v42.0.0/github#RepositoriesService.ListBranches
// TODO Send/Return these branch names back as the application requires and not only print them.
func GetBranches(owner string, repo string) {
	branches, _, err := githubClient.Repositories.ListBranches(context.Background(), owner, repo, nil)
	if err != nil {
		fmt.Println("Couldn't fetch the list of Repositories. ", err)
		os.Exit(1)
	}

	for i, branch := range branches {
		fmt.Println(i, branch.GetName())
	}
}

// Function to fetch commits from a repo
// ref: https://pkg.go.dev/github.com/google/go-github/v42/github#RepositoriesService.ListCommits
// TODO Get better data formats
func GetCommits(owner string, repo string) {
	commits, _, err := githubClient.Repositories.ListCommits(
		context.Background(),
		owner,
		repo,
		&github.CommitsListOptions{
			SHA:   "atherton",
			Since: time.Date(2022, 1, 1, 0, 0, 0, 0, time.Local),
			Until: time.Date(2022, 1, 15, 0, 0, 0, 0, time.Local),
		},
	)
	if err != nil {
		fmt.Println("Couldn't fetch the commits. ", err)
		os.Exit(1)
	}

	for i, commit := range commits {
		fmt.Println(
			i,
			commit.GetSHA(),
			/* commit.Commit.GetMessage(), */
			commit.Commit.GetAuthor().GetName(),
			commit.GetAuthor().GetLogin(),
		)
	}
}

func GetCommit(owner string, repo string, sha string) {
	commit, _, err := githubClient.Repositories.GetCommit(
		context.Background(),
		owner,
		repo,
		sha,
		nil,
	)

	if err != nil {
		fmt.Println("Couldn't get the commit", err)
	}

	fmt.Println(commit.GetSHA())
	for i, file := range commit.Files {
		fmt.Println(
			i,
			file.GetFilename(),
			file.GetBlobURL(),
			file.GetContentsURL(),
			file.GetRawURL(),
		)
	}
}
