package beginner

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func Read() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Simple shell ...")

	for {
		fmt.Println("->")
		text,_ := reader.ReadString('\n')
		text = strings.Replace(text,"\n","",-1)

		if strings.Compare("hi",text) == 0  {
			fmt.Println("hello!")
		}
	}

}
