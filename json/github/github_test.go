package github

import (
	"fmt"
	"log"
	"testing"
	"time"
)

func TestJSONDecode(t *testing.T) {
	issue := []string{"repo:golang/go", "is:open", "json", "decoder"}
	result, err := SearchIssues(issue)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%d issues:\n", result.TotalCount)
	//class := make(map[string]struct{ Issues })
	for _, item := range result.Items {
		fmt.Printf("#%-5d %9.9s %.55s\n", item.Number, item.User.Login, item.Title)
		//switch item.CreateAt {
		//case
		//}
	}
}

func TestDay(t *testing.T) {
	local, _ := time.LoadLocation("GMT")
	var time = time.Date(2022, 1, 1, 1, 1, 1, 1, local)
	fmt.Println(day(time))
}
func day(time time.Time) int64 {
	return time.UnixMilli()

}
