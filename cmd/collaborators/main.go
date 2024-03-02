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
		RepoName string `csv:"repo fullname"`
		UserName string `csv:"user name"`
		UserType string `csv:"user type"`
		Role     string `csv:"role"`
	}
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("Missing args[1]: Github Org")
	}
	if len(os.Args) < 3 {
		log.Fatal("Missing args[2]: Target File")
	}
	fmt.Println("Prepare target file: " + os.Args[2])
	targetFile, err := os.OpenFile(os.Args[2], os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		log.Fatal(err)
	}
	defer targetFile.Close()
	fmt.Println("Fetching private repo list from: " + os.Args[1])
	repoNames, err := gh.RepoNames(os.Args[1], "private")
	if err != nil {
		log.Fatal(err)
	}
	var rows []Row
	for _, repoName := range repoNames {
		fmt.Println("Fetching contributor from: " + repoName)
		collaborators, err := gh.Collaborators(repoName)
		if err != nil {
			rows = append(rows, Row{
				RepoName: repoName,
				UserName: err.Error(),
				UserType: "",
			})
		}
		for _, collaborator := range collaborators {
			rows = append(rows, Row{
				RepoName: repoName,
				UserName: collaborator.Login,
				UserType: collaborator.Type,
				Role:     collaborator.RoleName,
			})
		}
	}
	fmt.Println("Wrapping up to CSV")
	if err := gocsv.MarshalFile(&rows, targetFile); err != nil {
		log.Fatal(err)
	}
	fmt.Println("Done")
}
