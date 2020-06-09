package githubInfo

import (
	"context"
	"github.com/google/go-github/github"
	"github.com/sirupsen/logrus"
)

func getDetails(username string) UserInfo{
	client := github.NewClient(nil)
	ctx := context.Background()
	//list all organizations for a particular user
	orgs, _, err := client.Organizations.List(ctx, username, nil)
	if err != nil {
		logrus.Fatalln(err)
	}
	opt := &github.RepositoryListOptions{Type: "owner", Sort: "updated", Direction: "desc"}
	var allRepos []*github.Repository
	for {
		repos, resp, err := client.Repositories.List(ctx, username, opt)
		if err != nil {
			logrus.Fatalln(err)
		}
		allRepos = append(allRepos, repos...)
		if resp.NextPage == 0 {
			break
		}
		opt.Page = resp.NextPage
	}
	var details UserInfo
	details.Username = username
	details.NumberOfPublicRepos = len(allRepos)
	details.NumberOfJoinedOrganizations = len(orgs)
	return details
}