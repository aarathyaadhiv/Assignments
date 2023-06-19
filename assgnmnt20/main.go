package main

import "fmt"



func main(){
	k:=1
	for i:=1;i<=4;i++{
		for j:=1;j<=i;j++{
			fmt.Print(k," ")
			k++
		}
		fmt.Println()
	}
}