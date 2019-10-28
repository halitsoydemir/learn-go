package beginner

import (
	"fmt"
	"sort"
)

func SortingMethod() {
	fmt.Println("Go Sorting Tutorial")

	myInts := []int{1,3,2,6,3,4}
	fmt.Println(myInts)

	sort.Ints(myInts)
	fmt.Println(myInts)
}

func ComplexSortingMethod(){
	programmers := []Programmer{
		{Age:30},{Age:20},{Age:50},{Age:1000},
	}

	fmt.Println(programmers)
	sort.Sort(byAge(programmers))
	fmt.Println(programmers)
}

type Programmer struct {
	Age int
}

type byAge []Programmer

func (p byAge) Len() int{
	return len(p)
}

func (p byAge) Swap(i,j int){
	p[i],p[j] = p[j],p[i]
}

func (p byAge) Less(i,j int) bool {
	return p[i].Age<p[j].Age
}