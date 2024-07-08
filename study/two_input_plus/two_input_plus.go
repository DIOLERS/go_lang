package main

import (
	"fmt"
)

func main() {
var a, b int                   // int型の変数を宣言
fmt.Scanf("%d %d", &a, &b) // %dでint型を代入
fmt.Println(a+b)
}