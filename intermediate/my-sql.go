package intermediate

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

func ConnectToMySql() {
	fmt.Println("Go MySql Tutorial")

	db,err := sql.Open("mysql","root:my-secret-pw@tcp(127.0.0.1:3308)/test")
	handleError(err)

	insert,err := db.Query("INSERT INTO test VALUES (2,'TEST')")
	handleError(err)

	defer insert.Close()
	defer db.Close()

	//QUERY ALL
	results,err := db.Query("SELECT * FROM test")
	handleError(err)

	for results.Next() {
		var tag Tag
		err = results.Scan(&tag.ID,&tag.Name)
		handleError(err)

		log.Printf(tag.Name)
	}


	//QUERY SINGLE ROW
	log.Println("QUERY SINGLE ROW")
	var singleTag Tag
	errr := db.QueryRow("SELECT * FROM test WHERE name = ?","TEST").Scan(&singleTag.ID,&singleTag.Name)
	handleError(errr)

	log.Println(singleTag.Name,singleTag.ID)
}

func handleError(err error){
	if err != nil {
		panic(err.Error())
	}
}

type Tag struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}
