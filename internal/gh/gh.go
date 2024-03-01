package gh

import (
	"encoding/json"
	"os/exec"
)

// https://docs.github.com/rest

// Repos equivalent with `gh api orgs/ORG/repos --paginate“
func Repos(org string) ([]Repo, error) {
	b, err := exec.Command("gh", "api", "orgs/"+org+"/repos", "--paginate").Output()
	if err != nil {
		return nil, err
	}

	var repos []Repo
	if err := json.Unmarshal(b, &repos); err != nil {
		return nil, err
	}

	return repos, nil
}

// gh api orgs/ion-mobility/repos --jq '.[].full_name' --paginate

// gh api repos/ion-mobility/hmi/collaborators
// gh api repos/ion-mobility/hmi/contributors

// gh api orgs/ion-mobility/memberships/imantung

// gh api users/ion-mobility/members
