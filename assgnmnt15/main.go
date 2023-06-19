package main

import "fmt"

func getArray(limit int)[]int{
	array:=make([]int,limit)
	fmt.Println("enter the elements in array")
	for i:=0;i<limit;i++{
		fmt.Scanf("%d", &array[i])
	}
	return array
}
func displayArray(array[]int,limit int){
  for i:=0;i<limit;i++{
	fmt.Print(array[i],"\t")
  }
}
func main(){
	
	var limit int
	fmt.Println("enter the size of array")
	fmt.Scanf("%d",&limit)
	array:= getArray(limit)
	displayArray(array,limit)
}