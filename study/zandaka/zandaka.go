package main
import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)
func main(){
    sc := bufio.NewScanner(os.Stdin)
    var num[5] int
	for i := 0; i < 2; i++ {
		sc.Scan()
		num[i], _ = strconv.Atoi(sc.Text())
	}
	if(num[0] < num[1]){
	  fmt.Println("error")
	}else {
	  num[2] = num[0]-num[1]
	  fmt.Println(num[2])
	}
}