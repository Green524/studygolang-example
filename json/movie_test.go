package json

import (
	"encoding/json"
	"fmt"
	"log"
	"reflect"
	"testing"
)

func TestPrint(t *testing.T) {
	fmt.Println(movies)
}
func TestMarshal(t *testing.T) {
	data, err := json.MarshalIndent(movies, "", "	")
	if err != nil {
		log.Fatalf("JSON marshaling faild:%v", err)
	}
	fmt.Printf("%s\n", data)

	var titles []struct{ Title string }
	fmt.Println(reflect.TypeOf(titles))
	if err := json.Unmarshal(data, &titles); err != nil {
		log.Fatalf("JSON unmarshaling failed:%v", err)
	}
	fmt.Println(titles)
}
