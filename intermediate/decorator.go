package intermediate

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

func printFunction() {
	fmt.Println("Hello World")
}

func withLogging(a func()) {
	fmt.Println("Method starting...")
	a()
	fmt.Println("Method ended")
}

func FuncAsParameterTest() {
	fmt.Printf("Type: %T\n", withLogging)
	withLogging(printFunction)
}

func DecoratorWithRealWorldExample(){
	handleRequestsDec()
}

func homePageDec(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: homePage")
	fmt.Fprintf(w, "Welcome to the HomePage!")
}

func handleRequestsDec() {

	http.Handle("/", isAuthorized(homePageDec))
	log.Fatal(http.ListenAndServe(":8081", nil))
}

func isAuthorized(endpoint func(http.ResponseWriter, *http.Request)) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		fmt.Println("Checking to see if Authorized header set...")

		if val, ok := r.Header["Authorization"]; ok {
			fmt.Println(val)
			if len(val[0])>6 && strings.ToLower(val[0][0:6]) == "bearer" {
				fmt.Println("Header is set! We can serve content!")
				endpoint(w, r)
			}else{
				fmt.Println("Not Authorized!!")
				fmt.Fprintf(w, "Not Authorized!!")
			}
		} else {
			fmt.Println("Not Authorized!!")
			fmt.Fprintf(w, "Not Authorized!!")
		}
	})
}

func handlePanic(endpoint func(http.ResponseWriter, *http.Request)) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		fmt.Println("Checking to see if Authorized header set...")
		endpoint(w,r)
		if r := recover(); r != nil {
			http.Error(w,"Beklenmeyen hata oluştu! ",500)
		}
	})
}