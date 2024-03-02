package gh

import (
	"encoding/json"
	"errors"
	"os/exec"
	"strings"
)

// https://docs.github.com/rest

// Equivalent with: gh api orgs/ORG/repos --paginate“
func RetrieveRepos(org string) ([]Repo, error) {
	b, err := exec.Command("gh", "api", "orgs/"+org+"/repos", "--paginate").CombinedOutput()
	if err != nil {
		return nil, errors.New(string(b))
	}
	var repos []Repo
	if err := json.Unmarshal(b, &repos); err != nil {
		return nil, err
	}
	return repos, nil
}

// Equivalent with: gh repo list ORG --visibility VISIBILITY --json=nameWithOwner --jq='.[].nameWithOwner' -L 500
func RetrieveRepoNames(org, visibility string) ([]string, error) {
	b, err := exec.Command("gh", "repo", "list", org, "--visibility", visibility, "--json=nameWithOwner", "--jq=.[].nameWithOwner", "-L=500").CombinedOutput()
	if err != nil {
		return nil, errors.New(string(b))
	}
	names := strings.Split(strings.TrimSpace(string(b)), "\n")
	return names, nil
}

// Equivalent with: gh api repos/REPO_FULL_NAME/contributors --paginate
func RetrieveContributors(repoName string) ([]Contributor, error) {
	b, err := exec.Command("gh", "api", "repos/"+repoName+"/contributors", "--paginate").CombinedOutput()
	if err != nil {
		return nil, errors.New(string(b))
	}
	var contributors []Contributor
	if err := json.Unmarshal(b, &contributors); err != nil {
		return nil, err
	}
	return contributors, nil
}

// Equivalent with: gh api repos/REPO_FULL_NAME/collaborators --paginate
func RetrieveCollaborators(repoName string) ([]Collaborator, error) {
	b, err := exec.Command("gh", "api", "repos/"+repoName+"/collaborators", "--paginate").CombinedOutput()
	if err != nil {
		return nil, errors.New(string(b))
	}
	var collaborator []Collaborator
	if err := json.Unmarshal(b, &collaborator); err != nil {
		return nil, err
	}
	return collaborator, nil
}

// Equivalent with: gh api orgs/ion-mobility/members --jq '.[].login'
func RetrieveTeams(org string) ([]string, error) {
	b, err := exec.Command("gh", "api", "orgs/"+org+"/teams", "--jq", ".[].name", "--paginate").CombinedOutput()
	if err != nil {
		return nil, errors.New(string(b))
	}
	names := strings.Split(strings.TrimSpace(string(b)), "\n")
	return names, nil
}

func RetrieveTeamMembers(org, team string) ([]string, error) {
	b, err := exec.Command("gh", "api", "orgs/"+org+"/teams/"+team+"/members", "--jq", ".[].login", "--paginate").CombinedOutput()
	if err != nil {
		return nil, errors.New(string(b))
	}
	names := strings.Split(strings.TrimSpace(string(b)), "\n")
	return names, nil
}

func RetrieveTeamRepos(org, team string) ([]string, error) {
	b, err := exec.Command("gh", "api", "orgs/"+org+"/teams/"+team+"/repos", "--jq", ".[].full_name", "--paginate").CombinedOutput()
	if err != nil {
		return nil, errors.New(string(b))
	}
	names := strings.Split(strings.TrimSpace(string(b)), "\n")
	return names, nil
}
