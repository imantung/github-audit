# GITHUB AUDIT

Generate csv report for audit github organization purpose

## Prerequisites

- [Golang](https://go.dev/)
- [Github Cli](https://cli.github.com/) (2.44.1)

Notes: make sure you have admin access to the organization

## Usage

```bash
GITHUB_ORG=YOUR-ORG make
```

Or

```makefile
go run cmd/repos/main.go $(GITHUB_ORG) 		output_$(GITHUB_ORG)_$$(date +"%m%d%Y")/1_repos.csv
go run cmd/teams/main.go $(GITHUB_ORG) 		output_$(GITHUB_ORG)_$$(date +"%m%d%Y")/2_teams.csv
go run cmd/contributors/main.go $(GITHUB_ORG) 	output_$(GITHUB_ORG)_$$(date +"%m%d%Y")/3_contributors.csv
go run cmd/collaborators/main.go $(GITHUB_ORG) 	output_$(GITHUB_ORG)_$$(date +"%m%d%Y")/4_collaborators.csv
go run cmd/team-members/main.go $(GITHUB_ORG) 	output_$(GITHUB_ORG)_$$(date +"%m%d%Y")/5_team-members.csv
go run cmd/team-repos/main.go $(GITHUB_ORG) 	output_$(GITHUB_ORG)_$$(date +"%m%d%Y")/6_team-repos.csv
go run cmd/actions/main.go $(GITHUB_ORG) 	output_$(GITHUB_ORG)_$$(date +"%m%d%Y")/7_actions.csv
```

## CSV Details

1. Repos


```
id,name,fullname,description,size,language,topics,open issue count,created at,updated at,pushed at,private,fork,archived,disabled,secret scanning,secret scanning push protection,dependabot security updates,secret scanning validity checks
...
```

2. Teams 

```
team name,parent,description,create at,members count,repos count
...
```

3. Contributors

```
repo,user name,user type,contributions
...
```

4. Collaborators

```
repo fullname,user name,user type,role
...
```

5. Team Members

```
team,member
...
```

6. Team Repos

```
team,repo
...
```

7. Actions

```
repo,artifact count,run count
...
```


## License

This project is licensed under the MIT License - see the [LICENSE.md](LICENSE.md) file for details