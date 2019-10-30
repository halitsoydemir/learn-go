package intermediate

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"sync"
)
var counter int
var mutex *sync.Mutex = &sync.Mutex{}

func ServeWebServer() {
//	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
//		fmt.Fprintf(writer,"Hello, %q",html.EscapeString(request.URL.Path))
//	})

	http.HandleFunc("/hi", func(w http.ResponseWriter, r *http.Request){
		fmt.Fprintf(w, "Hi")
	})

	http.HandleFunc("/increment", incrementCounter)

//	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
//		http.ServeFile(w, r, r.URL.Path[1:])
//	})

	http.Handle("/", http.FileServer(http.Dir("./static")))

	log.Fatal(http.ListenAndServe(":8081", nil))
	//serve as https
	//http.ListenAndServeTLS(":443", "server.crt", "server.key", nil)
}

func incrementCounter(w http.ResponseWriter, r *http.Request) {
	mutex.Lock()
	counter++
	fmt.Fprintf(w, strconv.Itoa(counter))
	mutex.Unlock()
}
