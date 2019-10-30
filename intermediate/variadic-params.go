package intermediate

import "fmt"

func myVariadicFunction(args ...string) {
	fmt.Println(args)
}

func ParamsMethodTest() {
	//c# taki params aslında
	myVariadicFunction("hello", "world")
}