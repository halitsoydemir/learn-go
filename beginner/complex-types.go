package beginner

import "fmt"

func ArraysMethod() {
	var days []string
	days = []string{"tuesday", "wednesday"}
	fmt.Println(days[0])
}

func SliceMethod(){
	days := [...]string{"monday","tuesday", "wednesday","thursday","friday","saturday","sunday"}
	//that means starts from 0 (0 is included), end with 5 (5 is excluded)
	weekdays := days[0:5]
	fmt.Println(weekdays)
}

func MapMethod(){
	subscribers := map[string]int{
		"hltsydmr" : 5,
		"sıla" : 10,
	}

	// returns 0
	//fmt.Println(subscribers["hltsydmr2"])

	// returns 5
	fmt.Println(subscribers["hltsydmr"])
}

func StructMethod(){

	type Person struct {
		Name string
		Age int
	}

	var p1 Person
	p1.Name = "test name without new"
	fmt.Println(p1)

	elliot := Person{Name:"Elliot",Age:123}
	fmt.Println(elliot)

	elliot.Age = 18
	fmt.Println(elliot)

}



func NestedStructMethod(){
	type Person struct {
		name string
		age int
	}

	type Team struct {
		name string
		players [2]Person
	}

	var myTeam Team
	fmt.Println(myTeam)

	players := [...]Person{
		{
		name: "firstPerson",
		age:  1,
		},
		{
			name: "secondPerson",
			age:  2,
		},
	}

	celtic := Team{
		name:    "Celtic FC",
		players: players,
	}

	fmt.Println(celtic)
}