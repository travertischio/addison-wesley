// Issues prints a table of GitHub issues matching search terms
package main

import (
  "fmt"
  "log"
  "os"
  "time"

  "github.com/travertischio/addison-wesley/github"
)

func main() {
  now := time.Now()
  lastMonth := now.AddDate(0, -1, 0)
  lastYear := now.AddDate(-1, 0, 0)

  var thisMonth = make([]*github.Issue, 0)
  var thisYear = make([]*github.Issue, 0)
  var older = make([]*github.Issue, 0)

  result, err := github.SearchIssues(os.Args[1:])
  if err != nil {
    log.Fatal(err)
  }

  for _, item := range result.Items {
    if item.CreatedAt.Before(lastYear) {
      older = append(older, item)
    } else if item.CreatedAt.Before(lastMonth) {
      thisYear = append(thisYear, item)
    } else {
      thisMonth = append(thisMonth, item)
    }
  }

  fmt.Printf("%d issues:\n", result.TotalCount)
  fmt.Println("Less than a month old")
  for _, item := range thisMonth {
    fmt.Printf("#%-5d %9.9s %.55s\n", item.Number, item.User.Login, item.Title)
  }
  fmt.Println("\nLess than a year old")
  for _, item := range thisYear {
    fmt.Printf("#%-5d %9.9s %.55s\n", item.Number, item.User.Login, item.Title)
  }
  fmt.Println("\nOlder than one year")
  for _, item := range older {
    fmt.Printf("#%-5d %9.9s %.55s\n", item.Number, item.User.Login, item.Title)
  }
}
