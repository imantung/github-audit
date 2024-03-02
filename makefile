generate: 
	go run cmd/repos/main.go $(GITHUB_ORG) repos_$(GITHUB_ORG)_$$(date +"%m%d%Y").csv
	go run cmd/contributors/main.go $(GITHUB_ORG) contributors_$(GITHUB_ORG)_$$(date +"%m%d%Y").csv
	go run cmd/collaborators/main.go $(GITHUB_ORG) collaborators_$(GITHUB_ORG)_$$(date +"%m%d%Y").csv