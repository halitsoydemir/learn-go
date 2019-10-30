package intermediate

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func EnvironmentVariableTest() {
	//run export DATABASE_PASS=unicorns
	var dbString string
	dbString = os.Getenv("DATABASE_PASS")
	fmt.Printf("database pass: %s\n",dbString)

	//bu update globali etkilemez, çünkü child process cannot change it's parent process env variables
	//yani tekrar çalıstırdıgımızda initial yine unicorns olarak gelecek
	err := os.Setenv("DATABASE_PASS", "newunicorns")
	if err != nil {
		fmt.Println(err)
	}
	dbString = os.Getenv("DATABASE_PASS")
	fmt.Printf("Database Password: %s\n", dbString)

	handleEnvRequests()
}
func homePageForEnvTest(w http.ResponseWriter, r *http.Request) {
	//should need restart app
	if os.Getenv("FEATURE_TOGGLE") == "TRUE" {
		fmt.Println(os.Getenv("FEATURE_TOGGLE"))
		fmt.Fprintf(w, "Exciting New Feature")
	} else {
		fmt.Println(os.Getenv("FEATURE_TOGGLE"))
		fmt.Fprintf(w, "existing boring feature")
	}
}

func handleEnvRequests() {
	http.HandleFunc("/", homePageForEnvTest)
	log.Fatal(http.ListenAndServe(":8081", nil))
}
