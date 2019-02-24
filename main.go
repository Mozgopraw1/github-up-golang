package main

import (
	"context"
	"log"

	"github.com/google/go-github/v24/github"
)

func main() {
	client := github.NewClient(nil)
	opt := &github.RepositoryListOptions{}
	repos, _, err := client.Repositories.List(context.Background(), "Mozgopraw1", opt)

	if err != nil {
		log.Fatalf("failed to fetch repos: %s", err)
	}

	for _, repo := range repos {
		log.Printf("%s\t%s\n", repo.GetUpdatedAt().Format("02 Jan 15:04:05"), repo.GetName())
	}
}
