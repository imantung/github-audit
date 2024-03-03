generates: 
	mkdir output
	go run cmd/repos/main.go $(GITHUB_ORG) 			output_$(GITHUB_ORG)_$$(date +"%m%d%Y")/repos.csv
	go run cmd/contributors/main.go $(GITHUB_ORG) 	output_$(GITHUB_ORG)_$$(date +"%m%d%Y")/contributors.csv
	go run cmd/collaborators/main.go $(GITHUB_ORG) 	output_$(GITHUB_ORG)_$$(date +"%m%d%Y")/collaborators.csv
	go run cmd/team-members/main.go $(GITHUB_ORG) 	output_$(GITHUB_ORG)_$$(date +"%m%d%Y")/team-members.csv
	go run cmd/team-repos/main.go $(GITHUB_ORG) 	output_$(GITHUB_ORG)_$$(date +"%m%d%Y")/team-repos.csv
	go run cmd/actions/main.go $(GITHUB_ORG) 		output_$(GITHUB_ORG)_$$(date +"%m%d%Y")/actions.csv
