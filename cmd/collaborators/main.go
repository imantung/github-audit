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
	org := os.Args[1]
	targetFilename := os.Args[2]

	fmt.Println("Open target file: " + targetFilename)
	targetFile, err := os.OpenFile(targetFilename, os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		log.Fatal(err)
	}
	defer targetFile.Close()

	fmt.Println("Retrieve all repos from: " + org)
	repoNames, err := gh.RetrieveRepoNames(org)
	if err != nil {
		log.Fatal(err)
	}

	var rows []Row
	for _, repoName := range repoNames {
		fmt.Println("Retrieve collaborator from: " + repoName)
		collaborators, err := gh.RetrieveCollaborators(repoName)
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
