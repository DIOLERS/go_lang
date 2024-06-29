package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)

	var name string
	var num int

	sc.Scan()
	num, _ = strconv.Atoi(sc.Text())

	for i := 0; i < num; i++ {
		sc.Scan()
		name = sc.Text()
		fmt.Println(name)
	}
}
