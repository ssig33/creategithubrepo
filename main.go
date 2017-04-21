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

func createRepo(name, token string, private bool) {
	fmt.Println("Tring to create repository:", name)
	fmt.Println("Private:", private)
	client, ctx := getClient(token)

	repo := github.Repository{
		Name:    &name,
		Private: &private,
	}

	client.Repositories.Create(ctx, "", &repo)

	fmt.Println("Done")
}

func main() {
	token := os.Getenv("GITHUB_TOKEN")
	if token != "" {
		var name string
		var private bool
		flag.StringVar(&name, "name", "blank", "name of repository")
		flag.BoolVar(&private, "private", false, "private mode?")
		flag.BoolVar(&private, "p", false, "private mode?")
		flag.Parse()

		if name == "blank" {
			path, _ := os.Getwd()
			arr := strings.Split(path, "/")
			name = arr[len(arr)-1]
		}
		createRepo(name, token, private)
	} else {
		fmt.Println("Please set Environment Value: GITHUB_TOKEN")
	}

}
