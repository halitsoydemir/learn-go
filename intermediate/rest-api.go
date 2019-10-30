package intermediate

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"net/http"
)
var Articles []Article

func init(){
	Articles = []Article{
		{Id:"1",Title: "Hello", Desc: "Article Description", Content: "Article Content"},
		{Id:"2",Title: "Hello 2", Desc: "Article Description", Content: "Article Content"},
	}
}

func homePage(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "Welcome to the HomePage!")
	fmt.Println("Endpoint Hit: homePage")
}

func handleRequests() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/articles", returnAllArticles)
	log.Fatal(http.ListenAndServe(":10000", nil))
}

func handleRequestsWithMux() {
	// creates a new instance of a mux router
	myRouter := mux.NewRouter().StrictSlash(true)
	// replace http.HandleFunc with myRouter.HandleFunc
	myRouter.HandleFunc("/", homePage)
	myRouter.HandleFunc("/articles", returnAllArticles).Methods("GET")
	myRouter.HandleFunc("/articles", createNewArticle).Methods("POST")
	myRouter.HandleFunc("/articles/{id}", returnSingleArticle).Methods("GET")
	myRouter.HandleFunc("/articles/{id}", removeArticle).Methods("DELETE")
	// finally, instead of passing in nil, we want
	// to pass in our newly created router as the second
	// argument
	log.Fatal(http.ListenAndServe(":10000", myRouter))
}

func ServeSampleApi() {
	handleRequestsWithMux()
}

type Article struct {
	Id      string `json:"Id"`
	Title   string `json:"Title"`
	Desc    string `json:"desc"`
	Content string `json:"content"`
}

func returnAllArticles(w http.ResponseWriter, r *http.Request){
	fmt.Println("Endpoint Hit: returnAllArticles")
	json.NewEncoder(w).Encode(Articles)
}

func returnSingleArticle(w http.ResponseWriter, r *http.Request){
	fmt.Println("Endpoint Hit: returnSingleArticle")
	vars := mux.Vars(r)
	key := vars["id"]

	isExist := false
	for _,item := range Articles  {
		if item.Id == key {
			json.NewEncoder(w).Encode(item)
			isExist = true
		}
	}

	if !isExist {
		http.Error(w,"Kayıt bulunamadı!",404)
	}
}

func createNewArticle(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)
	var article Article
	json.Unmarshal(reqBody, &article)
	Articles = append(Articles, article)
	json.NewEncoder(w).Encode(article)
}

func removeArticle(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["id"]

	for i,item := range Articles {
		if item.Id == key {
			Articles = append(Articles[:i],Articles[i+1:]...)
		}
	}
}