package main

import "fmt"


func main(){
	var limit int
    var array [10]int
	var array2[10] int
	fmt.Println("enter the array limit")
	fmt.Scanf("%d",&limit)
	fmt.Println("enter the values in array")
	for i:=0;i<limit;i++{
		fmt.Scanf("%d",&array[i])
	}
	for i:=0;i<limit;i++{
		array2[i]=array[i] * array[i+1]
	}
	fmt.Println("output")
	for i:=0;i<limit-1;i++{
		fmt.Print(array2[i],"\t")
	}
}