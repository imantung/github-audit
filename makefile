GITHUB_ORG = ion-mobility

generate: 
	go run cmd/repos/main.go $(GITHUB_ORG) > repos_$(GITHUB_ORG)_$$(date +"%m%d%Y").csv