package intermediate

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"net/http"
)

type User struct {
	gorm.Model
	Name  string
	Email string
}

func CrudWithGormSqlite() {
	fmt.Println("Go ORM Tutorial")
	initDb()
	handleRequestsForGorm()
}

func initDb(){
	db,err := gorm.Open("sqlite3","user.db")
	if err != nil {
		fmt.Println(err.Error())
		panic("db connection failed")
	}
	defer db.Close()

	db.AutoMigrate(&User{})
}

func allUsers(w http.ResponseWriter, r *http.Request) {
	db,err := gorm.Open("sqlite3","user.db")
	if err != nil {
		fmt.Println(err.Error())
		panic("db connection failed")
	}
	defer db.Close()

	var users []User
	db.Find(&users)
	fmt.Println("{}",users)
	json.NewEncoder(w).Encode(users)

}

func newUser(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "New User Endpoint Hit")
	db,err := gorm.Open("sqlite3","user.db")
	if err != nil {
		fmt.Println(err.Error())
		panic("db connection failed")
	}
	defer db.Close()

	var user User
	vars := mux.Vars(r)
	user.Name = vars["name"]
	user.Email = vars["email"]

	db.Create(&user)

	fmt.Println("{}",user)
	json.NewEncoder(w).Encode(user)
}

func deleteUser(w http.ResponseWriter, r *http.Request) {
	db,err := gorm.Open("sqlite3","user.db")
	if err != nil {
		fmt.Println(err.Error())
		panic("db connection failed")
	}
	defer db.Close()


	vars := mux.Vars(r)
	name := vars["name"]

	var user User
	db.Where("name = ?",name).Find(&user)
	db.Delete(&user)
}

func updateUser(w http.ResponseWriter, r *http.Request) {
	db,err := gorm.Open("sqlite3","user.db")
	if err != nil {
		fmt.Println(err.Error())
		panic("db connection failed")
	}
	defer db.Close()


	vars := mux.Vars(r)
	name := vars["name"]
	email := vars["email"]

	var user User
	db.Where("name = ?",name).Find(&user)
	user.Email = email
	db.Save(&user)
	json.NewEncoder(w).Encode(user)
}

func handleRequestsForGorm() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/users", allUsers).Methods("GET")
	myRouter.HandleFunc("/user/{name}", deleteUser).Methods("DELETE")
	myRouter.HandleFunc("/user/{name}/{email}", updateUser).Methods("PUT")
	myRouter.HandleFunc("/user/{name}/{email}", newUser).Methods("POST")
	log.Fatal(http.ListenAndServe(":8081", myRouter))
}