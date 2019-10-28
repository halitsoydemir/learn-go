package beginner

import "fmt"

func TestInterfaceMethod() {
	println(25)
}

func print(input interface{}){
	fmt.Println(input)
}

func TestCustomInterface(){
	var player BaseGuitarist
	player.Name = "Paul"
	player.PlayGuitar()

	var player2 AcousticGuitarist
	player2.Name = "Ringo"
	player2.PlayGuitar()

	var guitarists []IGuitarist
	//interface üzeridneki contractı sağlıyorsa implement ediodur demek
	guitarists = append(guitarists, player)
	guitarists = append(guitarists, player2)
}

type IGuitarist interface {
	PlayGuitar()
}

type BaseGuitarist struct {
	Name string
}

type AcousticGuitarist struct {
	Name string
}

func(b BaseGuitarist) PlayGuitar(){
	fmt.Printf("%s BassGuitar",b.Name)
}

func(a AcousticGuitarist) PlayGuitar(){
	fmt.Printf("%s AcousticGuitar",a.Name)
}


//---

type Employee interface {
	Language() string
	Age() int
	Random() (string, error)
}

type Engineer struct {
	Name string
}

func (e *Engineer) Language() string {
	return e.Name + " programs in Go"
}

func ComplexInterface() {
	//var programmers []Employee
	//elliot := Engineer{Name: "Elliot"}
	// Engineer does not implement the Employee interface
	// you'll need to implement Age() and Random()
	//programmers = append(programmers, elliot)
}