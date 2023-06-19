package main

import "fmt"



func main(){
	var limit int
	var array1[10][10]int 
	var array2[10][10] int
	var sum[10][10] int
	fmt.Println("enter the size of the arrays")
	fmt.Scanf("%d",&limit)
	fmt.Println("enter the elements in first array")
	for i:=0;i<limit;i++{
		for j:=0;j<limit;j++{
			fmt.Scanf("%d",&array1[i][j])
		}
	}
	
	fmt.Println("enter the elements in the array2")
	for i:=0;i<limit;i++{
		for j:=0;j<limit;j++{
			fmt.Scanf("%d",&array2[i][j])
		}
	}
	
	for i:=0;i<limit;i++{
		for j:=0;j<limit;j++{
			sum[i][j]=array1[i][j]+array2[i][j]
		}
	}
	for i:=0;i<limit;i++{
		for j:=0;j<limit;j++{
			fmt.Print(sum[i][j],"\t")
		}
		fmt.Println()
	}
}