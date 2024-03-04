generates: 
	mkdir output_$(GITHUB_ORG)_$$(date +"%m%d%Y")
	go run cmd/repos/main.go $(GITHUB_ORG) 		output_$(GITHUB_ORG)_$$(date +"%m%d%Y")/1_repos.csv
	go run cmd/teams/main.go $(GITHUB_ORG) 		output_$(GITHUB_ORG)_$$(date +"%m%d%Y")/2_teams.csv
	go run cmd/contributors/main.go $(GITHUB_ORG) 	output_$(GITHUB_ORG)_$$(date +"%m%d%Y")/3_contributors.csv
	go run cmd/collaborators/main.go $(GITHUB_ORG) 	output_$(GITHUB_ORG)_$$(date +"%m%d%Y")/4_collaborators.csv
	go run cmd/team-members/main.go $(GITHUB_ORG) 	output_$(GITHUB_ORG)_$$(date +"%m%d%Y")/5_team-members.csv
	go run cmd/team-repos/main.go $(GITHUB_ORG) 	output_$(GITHUB_ORG)_$$(date +"%m%d%Y")/6_team-repos.csv
	go run cmd/actions/main.go $(GITHUB_ORG) 	output_$(GITHUB_ORG)_$$(date +"%m%d%Y")/7_actions.csv
