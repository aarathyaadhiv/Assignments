package main

import "fmt"


func main(){
	var p int
	var r float64
	var n float64
	fmt.Println("enter the principal amount, rate of interest, no of years")
	fmt.Scanf("%d %f %f",&p, &r, &n)
	
	si:= (float64(p)*r*n)/100

	fmt.Printf("simple interest is %f",si)
}