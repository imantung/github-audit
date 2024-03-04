package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/gocarina/gocsv"
	"github.com/imantung/github-audit/internal/gh"
)

type (
	Row struct {
		TeamName     string `csv:"team name"`
		Parent       string `csv:"parent"`
		Description  string `csv:"description"`
		CreatedAt    string `csv:"create at"`
		MembersCount int    `csv:"members count"`
		ReposCount   int    `csv:"repos count"`
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
	for _, teamName := range teams {
		fmt.Println("Retrieve team detail from: " + teamName)
		team, err := gh.RetrieveTeamDetails(org, teamName)
		if err != nil {
			rows = append(rows, Row{
				TeamName:    teamName,
				Description: err.Error(),
			})
		} else {
			parent := ""
			if team.Parent != nil {
				parent = team.Parent.Name
			}
			rows = append(rows, Row{
				TeamName:     teamName,
				Parent:       parent,
				Description:  team.Description,
				CreatedAt:    team.CreatedAt.Format(time.DateOnly),
				MembersCount: team.MembersCount,
				ReposCount:   team.ReposCount,
			})
		}
	}
	fmt.Println("Wrapping up to CSV")
	if err := gocsv.MarshalFile(&rows, targetFile); err != nil {
		log.Fatal(err)
	}
	fmt.Println("Done")
}
