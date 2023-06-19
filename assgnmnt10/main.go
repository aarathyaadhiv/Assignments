package main

import "fmt"



func main(){
	var array1 [10] int
	var array2 [10] int
	var limit,temp int

	fmt.Println("enter the limit")
	fmt.Scanf("%d",&limit)
	fmt.Println("enter the elements in array1")
	for i:=0;i<limit;i++{
		fmt.Scanf("%d",&array1[i])
	}
	fmt.Println("enter the elements in array2")
	for i:=0;i<limit;i++{
		fmt.Scanf("%d",&array2[i])
	}
	for i:=0;i<limit;i++{
		temp=array1[i]
		array1[i]=array2[i]
		array2[i]=temp
	}
	fmt.Println("arrays after swapping")
	fmt.Println("array1: ")
	for i:=0;i<limit;i++{
		fmt.Print(array1[i],"\t")
	}
	fmt.Println("\narray2: ")
	for i:=0;i<limit;i++{
		fmt.Print(array2[i],"\t")
	}
}