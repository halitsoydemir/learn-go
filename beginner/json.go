package beginner

import (
	"encoding/json"
	"fmt"
)

type Book struct {
	Title string `json:"title"`
	Author Author `json:"author"`
}

type Author struct {
	Sales int `json:"book_sales"`
	Age int `json:"age"`
	Developer bool `json:"is_developer"`
}

func MarshallTest() {
	author := Author{Sales: 3, Age: 25, Developer: true}
	book := Book{Title: "Learning Concurrency in Python", Author: author}

	// without indendation
	byteArray, err := json.Marshal(book)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(book)
	fmt.Println(string(byteArray))

	// with indendation
	byteArray, err = json.MarshalIndent(book,"", "  ")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(book)
	fmt.Println(string(byteArray))
}

func UnMarshallTest()  {
	jsonString := `{"name": "battery sensor", "capacity": 40, "time": "2019-01-21T19:07:28Z"}`

	//typed unmarshall
	var reading SensorReading
	err := json.Unmarshal([]byte(jsonString), &reading)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%+v\n", reading)
	fmt.Println(reading)

	//generic unmarshall
	var reading2 map[string]interface{}
	err = json.Unmarshal([]byte(jsonString), &reading2)
	fmt.Println(reading)
}

type SensorReading struct {
	Name string `json:"name"`
	Capacity int `json:"capacity"`
	Time string `json:"time"`
}