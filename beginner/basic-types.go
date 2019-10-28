package beginner

import "fmt"

func ExceedLimit() {
	var myInt int8
	//build error
	//myInt = 2500
	fmt.Println(myInt)
	//max value can be 127, min is -127
	// so it makes panic like,  constant 2500 overflows int8
}

func Convert(){
	var men uint8
	men = 5

	var women uint16
	women = 6

	var people int
	// not compile
	//people = men+women

	people = int(men) + int(women)
	fmt.Println(people)
}

func MaxFloat(){
	var maxFloat32 float32
	maxFloat32 = 16777216
	fmt.Println(maxFloat32 == maxFloat32+10)
	// it returns false
	fmt.Println(maxFloat32+10) // 16777226
	fmt.Println(maxFloat32+2000000) // 18777216
}

func FloatToIntIntToFloat(){
	var myint int = 11
	myfloat := float64(myint)
	var myfloat2 float64 = 0.23
	myint2 := int(myfloat2)

	fmt.Println(myint)
	fmt.Println(myfloat)

	fmt.Println(myfloat2)
	fmt.Println(myint2)
}

func BoolControl(){
	var isTrue bool = true
	var isFalse bool = false


	if isTrue && isFalse {
		fmt.Println("Both Conditions need to be True")
	}


	if isTrue || isFalse {
		fmt.Println("Only one condition needs to be True")
	}
}

func StringAndConst(){
	var myName string
	myName = "Elliot Forbes"

	const meaningOfLife = 42
	//cannot do this because it is constant
	//meaningOfLife = 11

	fmt.Println(myName)
	fmt.Println(meaningOfLife)
}