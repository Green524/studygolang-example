package wait

import (
	"log"
	"testing"
)

func TestWaitForServer(t *testing.T) {
	log.SetPrefix("wait:")
	log.SetFlags(0)
	if err := WaitForServer("https://golang.org"); err != nil {
		//fmt.Fprintf(os.Stderr, "Site is down:%v\n", err)
		//os.Exit(1)

		//等价上面两行代码
		log.Fatalf("Site is down:%v", err)

	}

	//if err := Ping(); err != nil {
	//	log.Printf("ping failed: %v; networking disabled", err)
	//}
}
