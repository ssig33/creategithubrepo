package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
	"os"
	"strings"
)

func getClient(token string) (*github.Client, context.Context) {
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(

		&oauth2.Token{AccessToken: token},
	)
	tc := oauth2.NewClient(ctx, ts)
	return github.NewClient(tc), ctx
}

func createRepo(nameFlag, token string) {
	fmt.Println("Tring to create repository:", nameFlag)
	client, ctx := getClient(token)

	repo := github.Repository{
		Name: &nameFlag,
	}

	client.Repositories.Create(ctx, "", &repo)

	fmt.Println("Done")
}

func main() {
	token := os.Getenv("GITHUB_TOKEN")
	if token != "" {
		var nameFlag string
		flag.StringVar(&nameFlag, "name", "blank", "string flag")

		if nameFlag == "blank" {
			path, _ := os.Getwd()
			arr := strings.Split(path, "/")
			nameFlag = arr[len(arr)-1]
			fmt.Println(nameFlag)
		}
		createRepo(nameFlag, token)
	} else {
		fmt.Println("Please set Environment Value: GITHUB_TOKEN")
	}

}
