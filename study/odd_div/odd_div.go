package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)

	sc.Scan()
	name := sc.Text()

	rs := []rune(name)
	for i := 0; i < len(name); i++ {
		if i%2 == 0 {
			fmt.Print(string(rs[i]))
		}
	}
}
