package main

import "fmt"

func getarray(limit int)[][]int{
	array:= make([][]int,limit)
	for i:=0;i<limit;i++{
		array[i]= make([]int, limit)
	}
	for i:=0;i<limit;i++{
		for j:=0;j<limit;j++{
			fmt.Scanf("%d",&array[i][j])
		}
	}
	return array
}
func displayarray(array[][]int){
	for i:=0;i<len(array);i++{
		for j:=0;j<len(array);j++{
			fmt.Print(array[i][j],"\t")
		}
		fmt.Println()
	}
}

func main(){
	var limit int
	fmt.Println("enter the limit of the array")
	fmt.Scanf("%d",&limit)
	fmt.Println("enter the elements in array")
	array := getarray(limit)
	fmt.Println("array is: ")
	displayarray(array)

}