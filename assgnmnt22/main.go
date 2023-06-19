
package main

import "fmt"

func getarray(limit int) [][]int {

	array := make([][]int,limit)
	for i:=0;i<limit;i++{
		array[i] = make([]int, limit)
	}
	for i := 0; i < limit; i++ {
		for j := 0; j < limit; j++ {
			fmt.Scanf("%d", &array[i][j])
		}

	}
	return array
}
func addarray(array [][]int, array2 [][]int) [][]int {
	add := make([][]int, len(array))
	for i:=0;i<len(array);i++{
		add[i]=make([]int, len(array))
	}
	for i := 0; i < len(array); i++ {
		for j := 0; j < len(array); j++ {
			add[i][j] = array[i][j] + array2[i][j]
		}
	}
	return add
}
func displayarray(array [][]int) {
	for i := 0; i < len(array); i++ {
		for j := 0; j < len(array); j++ {
			fmt.Print(array[i][j],"\t")
		}
		fmt.Println()
	}
}

func main() {
	var limit int
	fmt.Println("enter the size of array")
	fmt.Scanf("%d", &limit)
	fmt.Println("enter the elements in array1")
	array := getarray(limit)
	fmt.Println("enter the elements in array2")
	array2 := getarray(limit)
	add := addarray(array, array2)
	fmt.Println("result array is : ")
	displayarray(add)

}
