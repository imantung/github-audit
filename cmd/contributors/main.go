package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gocarina/gocsv"
	"github.com/imantung/github-audit/internal/gh"
)

type (
	Row struct {
		Repo          string `csv:"repo"`
		UserName      string `csv:"user name"`
		UserType      string `csv:"user type"`
		Contributions int    `csv:"contributions"`
	}
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("Missing args[1]: Github Org")
	}
	if len(os.Args) < 3 {
		log.Fatal("Missing args[2]: Target File")
	}
	org := os.Args[1]

	fmt.Println("Prepare target file: " + os.Args[2])
	targetFile, err := os.OpenFile(os.Args[2], os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		log.Fatal(err)
	}
	defer targetFile.Close()

	fmt.Println("Retrive private repos from: " + org)
	repoNames, err := gh.RetrievePrivateRepoNames(org)
	if err != nil {
		log.Fatal(err)
	}

	var rows []Row
	for _, repoName := range repoNames {
		fmt.Println("Fetching contributor from: " + repoName)
		contributors, err := gh.RetrieveContributors(repoName)
		if err != nil {
			rows = append(rows, Row{
				Repo:          repoName,
				UserName:      err.Error(),
				UserType:      "",
				Contributions: -1,
			})
		}
		for _, contributor := range contributors {
			rows = append(rows, Row{
				Repo:          repoName,
				UserName:      contributor.Login,
				UserType:      contributor.Type,
				Contributions: contributor.Contributions,
			})
		}
	}
	fmt.Println("Wrapping up to CSV")
	if err := gocsv.MarshalFile(&rows, targetFile); err != nil {
		log.Fatal(err)
	}
	fmt.Println("Done")
}
