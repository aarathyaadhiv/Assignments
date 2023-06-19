package main

import "fmt"


func main(){
	var mark float64
	fmt.Println("enter the mark")

	fmt.Scanf("%f",&mark)

	if mark<50 {
		fmt.Println("failed")
	}else{
		fmt.Println("passed")
	}
}