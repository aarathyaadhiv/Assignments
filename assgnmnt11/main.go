package main

import "fmt"


func main(){
	var limit int
	var array[10] int

	fmt.Println("enter the size of the array")
	fmt.Scanf("%d",&limit)
	fmt.Println("enter the values in the array")
	for i:=0;i<limit;i++{
		fmt.Scanf("%d",&array[i])
	}
	count:=0
	for i:=0;i<limit;i++{
		if array[i]%2==0{
			count++
		}
	}
	fmt.Println("count of even numbers: ",count)
}