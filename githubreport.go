package main

import (
	"fmt"
	"strconv"
	github "github.com/wesovi/githubreport/github"
	domain "github.com/wesovi/githubreport/domain"
)

func main() {
	developer := domain.Developer{}
	github.PopulateDeveloperDetails(&developer)
	fmt.Println("User Github id "+strconv.Itoa(developer.Id))
	fmt.Println("User Github login "+developer.Login)
}
