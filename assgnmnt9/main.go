package main

import "fmt"



func main(){
	
	for i:=1;i<=5;i++{
		k:=1
		for j:=1;j<=i;j++{
			fmt.Print(k)
			k++
		}
		fmt.Println()
	}
}