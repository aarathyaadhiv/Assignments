package main

import "fmt"


func main(){
	var number1 int
	var number2 float32
	fmt.Println("enter 2 numbers")
	

	fmt.Scanf("%d",&number1)
	fmt.Scanf("%f",&number2)
	
	sum := number2+float32(number1)

	fmt.Println(sum)

}