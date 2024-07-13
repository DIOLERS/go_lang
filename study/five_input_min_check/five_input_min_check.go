package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)

	var name[5] string
	var num[5] int
    var mnum int

   mnum = 99

	for i := 0; i < 5; i++ {
		sc.Scan()
		num[i], _ = strconv.Atoi(sc.Text())
		name[i] = sc.Text()
		if(i==0){
			mnum = num[i]
		}
		else if(mnum > num[i]){
		mnum = num[i]
		}
	}
			fmt.Println(mnum)
}