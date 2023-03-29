package count

import (
	"fmt"
	"log"
	"net/http"
)

type dollars float32

func (d dollars) String() string {
	return fmt.Sprintf("$%.2f", d)
}

type database map[string]dollars

func (db database) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	switch req.URL.Path {
	case "/list":
		for item, price := range db {
			fmt.Fprintf(w, "%s: %s\n", item, price)
		}
	case "/price":
		item := req.URL.Query().Get("item")
		fmt.Fprintf(w, "%s: %s", item, db[item])
	default:
		//w.WriteHeader(http.StatusNotFound) //这一步应该在response之前执行，不然会返回200
		//fmt.Fprintf(w, "not such page:%s", req.URL)
		msg := fmt.Sprintf("not such page:%s", req.URL)
		http.Error(w, msg, http.StatusNotFound)
	}
}

func (db database) list(w http.ResponseWriter, req *http.Request) {
	for item, price := range db {
		fmt.Fprintf(w, "%s: %s \n", item, price)
	}
}

func (db database) price(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	price, ok := db[item]
	if !ok {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "not such page: %q", item)
		return
	}
	fmt.Fprintf(w, "%s \n", price)
}
func main() {
	db := database{"shoes": 50, "socks": 5}
	//log.Fatal(http.ListenAndServe("localhost:8000", db))

	//mux := http.NewServeMux()
	//mux.Handle("/list", http.HandlerFunc(db.list)) //http.HandlerFunc(db.list) 是一个转换，不是一个方法调用
	//mux.Handle("/price", http.HandlerFunc(db.price))
	//简化
	//mux.HandleFunc("/list", db.list)
	//mux.HandleFunc("/price", db.price)
	//全局的ServeMux，使用DefaultServeMux
	http.HandleFunc("/list", db.list)
	http.HandleFunc("/price", db.price)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}
