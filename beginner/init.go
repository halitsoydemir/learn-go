package beginner

import "fmt"

var name string

//paket içerisinde her hangi bi yer için tüm paketteki initler çalışır
func init() {
	fmt.Println("This will get called on main initialization")
	name = "Elliot"
}

func TestInit() {
	fmt.Println("My Wonderful Go Program")
	fmt.Printf("Name: %s\n", name)
}
