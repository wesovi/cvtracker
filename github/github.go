package github

import (
	"fmt"
	restClient "github.com/jmcvetta/napping"	
	domain "github.com/wesovi/githubreport/domain"
)

const githubApiUrL = "https://api.github.com"

func PopulateDeveloperDetails(developer *domain.Developer)(){
	resp, err := restClient.Get(githubApiUrL+"/users/ivancorrales", nil, &developer, nil)
	if err != nil {
		fmt.Println("Something went wrong...")
		fmt.Println(err)
		panic(err)
	}
	fmt.Println("response Status:", resp.Status())
	if resp.Status() != 200 {
		fmt.Println(err)
		panic(err)
	}
	
}