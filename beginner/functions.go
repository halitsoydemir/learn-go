package beginner

import "fmt"

func GetFullName(firstName string, lastName string) string{
	fullName := firstName + " " + lastName
	fmt.Println(fullName)
	return fullName
}

func GetFullNameWithErrorDetail(firstName string, lastName string) (string,error){
	fullName := firstName + " " + lastName
	fmt.Println(fullName)
	return fullName,nil
}

func AnonymousFunction() func() int {
	var x int
	return func() int {
		x++
		return x+1
	}

}

