package main

import (
	"fmt"
	
)


func main(){
	var num int
	fmt.Println("enter a number")
	fmt.Scanf("%d",&num)
	for i:=1;i<=10;i++{
		fmt.Println(i ," x ",num ,"=" ,i*num)
	}
}