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
		Team   string `csv:"team"`
		Member string `csv:"member"`
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

	fmt.Println("Prepare target file: " + targetFilename)
	targetFile, err := os.OpenFile(targetFilename, os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		log.Fatal(err)
	}
	defer targetFile.Close()

	fmt.Println("Retrieve teams from: " + org)
	teams, err := gh.RetrieveTeams(org)
	if err != nil {
		log.Fatal(err)
	}
	var rows []Row
	for _, team := range teams {
		members, err := gh.RetrieveTeamMembers(org, team)
		if err != nil {
			rows = append(rows, Row{
				Team:   team,
				Member: err.Error(),
			})
		}
		for _, member := range members {
			rows = append(rows, Row{
				Team:   team,
				Member: member,
			})
		}
	}

	fmt.Println("Wrapping up to CSV")
	if err := gocsv.MarshalFile(&rows, targetFile); err != nil {
		log.Fatal(err)
	}
	fmt.Println("Done")
}
