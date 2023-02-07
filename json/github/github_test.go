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

	numOfYear := make(map[string][]Issues)
	for _, item := range result.Items {
		//fmt.Printf("#%-5d %9.9s %.55s %v\n", item.Number, item.User.Login, item.Title, item.CreateAt)
		var days = item.CreatedAt.UnixMilli()
		if time.Now().UnixMilli()-days < 1000*60*60*24*30 {
			lt30 := numOfYear["<30"]
			lt30 = append(lt30, *item)
			numOfYear["<30"] = lt30
		} else if time.Now().UnixMilli()-days < 1000*60*60*24*365 {
			lt365 := numOfYear["<365"]
			lt365 = append(lt365, *item)
			numOfYear["<30"] = lt365
		} else {
			other := numOfYear["other"]
			other = append(other, *item)
			numOfYear["other"] = other
		}
	}
	fmt.Printf("Count:\n<30d <365d >365d\n")
	fmt.Printf("%d\t%d\t%d\n", len(numOfYear["<30"]), len(numOfYear["<365"]), len(numOfYear["other"]))
}

func TestDate(t *testing.T) {
	loc, _ := time.LoadLocation("GMT")
	var date = time.Date(2022, 12, 1, 0, 0, 0, 0, loc)
	fmt.Println(time.Now().UnixMilli()-date.UnixMilli() > 1000*60*60*24*30)
}
