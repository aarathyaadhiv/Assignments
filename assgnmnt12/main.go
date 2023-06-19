package main

import "fmt"


func main(){
	var array[10] int
	var limit,temp int
	fmt.Println("enter the size of the array")
	fmt.Scanf("%d",&limit)
	fmt.Println("enter the values in the array")
	for i:=0;i<limit;i++{
		fmt.Scanf("%d",&array[i])
	}
	for i:=0;i<limit;i++{
		for j:=i+1;j<limit;j++{
			if array[i]<array[j]{
				temp=array[i]
				array[i]=array[j]
				array[j]=temp
			}
		}
	}
	fmt.Println("sorted array is")
	for i:=0;i<limit;i++{
		fmt.Print(array[i],"\t")
	}
}