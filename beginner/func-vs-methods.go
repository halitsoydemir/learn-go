package beginner

import "fmt"

func FuncVsMethod() {

	g := Guitarist{
		Name: "Halit",
		Age:  1,
	}

	fmt.Println(g)

	UpdateGuitarist(&g,2)
	fmt.Println(g)

	g.Update(3)
	fmt.Println(g)


}

type Guitarist struct {
	Name string
	Age int
}

//function
func UpdateGuitarist(guitarist *Guitarist,age int){
	guitarist.Age = age
}

//method
func (g *Guitarist) Update(age int){
	g.Age = age
}
