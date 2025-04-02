package httpsql

import (
	"io"
	"fmt"
	"log"
	"net/http"
	//"sync"
	//"time"
)
//EOF - special error "END OF FILE"

func handler(r *http.Request) {
	fmt.Println("Request method: ", r.Method)
	fmt.Println("Request query: ", r.URL.Query().Get("id"))
	fmt.Println("Request URL: ", r.URL.Path)
}

func Les1() {
	resp, err := http.DefaultClient.Get("https://github.com/Shyyw1e/yells_post")
	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(body))

	// var r *http.Request

	// r.Method = "GET"
	// r.URL.Path = "https://github.com/Shyyw1e"
	
	// handler(r)


}