package main

import (
	"log"
	"os"
	"strings"
	"time"

	"github.com/gocarina/gocsv"
	"github.com/imantung/github-audit/internal/gh"
)

type (
	Row struct {
		ID          int    `csv:"id"`
		Name        string `csv:"name"`
		FullName    string `csv:"fullname"`
		Description string `csv:"description"`
		Size        int    `csv:"size"`
		Language    string `csv:"language"`
		Topics      string `csv:"topics"`

		OpenIssuesCount int `csv:"open issue count"`

		CreatedAt string `csv:"created at"`
		UpdatedAt string `csv:"updated at"`
		PushedAt  string `csv:"pushed at"`

		Private                      bool `csv:"private"`
		Fork                         bool `csv:"fork"`
		Archived                     bool `csv:"archived"`
		Disabled                     bool `csv:"disabled"`
		SecretScanning               bool `csv:"secret scanning"`
		SecretScanningPushProtection bool `csv:"secret scanning push protection"`
		DependabotSecurityUpdates    bool `csv:"dependabot security updates"`
		SecretScanningValidityChecks bool `csv:"secret scanning validity checks"`
	}
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("Missing args[1]: Github Org")
	}
	repos, err := gh.Repos(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	var rows []Row
	for _, e := range repos {
		rows = append(rows, Row{
			ID:                           e.ID,
			Name:                         e.Name,
			FullName:                     e.FullName,
			Private:                      e.Private,
			Description:                  e.Description,
			Fork:                         e.Fork,
			CreatedAt:                    e.CreatedAt.Format(time.DateOnly),
			UpdatedAt:                    e.UpdatedAt.Format(time.DateOnly),
			PushedAt:                     e.PushedAt.Format(time.DateOnly),
			Size:                         e.Size,
			Language:                     e.Language,
			Archived:                     e.Archived,
			Disabled:                     e.Disabled,
			OpenIssuesCount:              e.OpenIssuesCount,
			Topics:                       strings.Join(e.Topics, ", "),
			SecretScanning:               e.SecurityAndAnalytics.SecretScanning.Status == "enabled",
			SecretScanningPushProtection: e.SecurityAndAnalytics.SecretScanningPushProtection.Status == "enabled",
			DependabotSecurityUpdates:    e.SecurityAndAnalytics.DependabotSecurityUpdates.Status == "enabled",
			SecretScanningValidityChecks: e.SecurityAndAnalytics.SecretScanningValidityChecks.Status == "enabled",
		})
	}
	if err := gocsv.Marshal(&rows, os.Stdout); err != nil {
		log.Fatal(err)
	}
}
