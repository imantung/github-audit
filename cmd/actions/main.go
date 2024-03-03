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
		ArtifactCount string `csv:"artifact count"`
		RunCount      string `csv:"run count"`
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
		fmt.Println("Retrieve artifact total count from: " + repoName)
		artifactCount, err := gh.RetrieveArtifactTotalCount(repoName)
		if err != nil {
			artifactCount = err.Error()
		}

		fmt.Println("Retrieve run total count from: " + repoName)
		runCount, err := gh.RetrieveRunTotalCount(repoName)
		if err != nil {
			runCount = err.Error()
		}

		rows = append(rows, Row{
			Repo:          repoName,
			ArtifactCount: artifactCount,
			RunCount:      runCount,
		})

	}
	fmt.Println("Wrapping up to CSV")
	if err := gocsv.MarshalFile(&rows, targetFile); err != nil {
		log.Fatal(err)
	}
	fmt.Println("Done")
}
