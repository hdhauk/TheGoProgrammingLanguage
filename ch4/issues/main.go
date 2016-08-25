// Issues prints a table of GitHub issues matching the search terms.
package main

import (
	"fmt"
	"log"
	"os"

	"github.com/hdhauk/TheGoProgrammingLanguage/ch4/github"
)

func main() {
	result, err := github.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%d issues:\n", result.TotalCount)
	for _, item := range result.Items {
		fmt.Printf("#%-6d %9.9s %13.55s %50.55s\n", item.Number, item.User.Login, item.Title, item.REPOURL)
	}
}
