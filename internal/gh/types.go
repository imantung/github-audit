package gh

import "time"

type (
	Repo struct {
		ID                   int                  `json:"id"`
		Name                 string               `json:"name"`
		FullName             string               `json:"full_name"`
		Private              bool                 `json:"private"`
		Description          string               `json:"description"`
		Fork                 bool                 `json:"fork"`
		CreatedAt            time.Time            `json:"created_at"`
		UpdatedAt            time.Time            `json:"updated_at"`
		PushedAt             time.Time            `json:"pushed_at"`
		Size                 int                  `json:"size"`
		Language             string               `json:"language"`
		Archived             bool                 `json:"archived"`
		Disabled             bool                 `json:"disabled"`
		OpenIssuesCount      int                  `json:"open_issues_count"`
		Topics               []string             `json:"topics"`
		SecurityAndAnalytics SecurityAndAnalytics `json:"security_and_analysis"`
	}
	SecurityAndAnalytics struct {
		SecretScanning               Status `json:"secret_scanning"`
		SecretScanningPushProtection Status `json:"secret_scanning_push_protection"`
		DependabotSecurityUpdates    Status `json:"dependabot_security_updates"`
		SecretScanningValidityChecks Status `json:"secret_scanning_validity_checks"`
	}
	Status struct {
		Status string `json:"status"`
	}
	Contributor struct {
		Login         string `json:"login"`
		Type          string `json:"type"`
		Contributions int    `json:"contributions"`
	}
	Collaborator struct {
		Login    string `json:"login"`
		Type     string `json:"type"`
		RoleName string `json:"role_name"`
	}
	Team struct {
		Name         string    `json:"name"`
		Description  string    `json:"description"`
		CreatedAt    time.Time `json:"created_at"`
		MembersCount int       `json:"members_count"`
		ReposCount   int       `json:"repos_count"`
		Parent       *Team     `json:"parent"`
	}
)
