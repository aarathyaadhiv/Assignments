package main

import "fmt"



func main(){
	var limit,sum int
	fmt.Println("enter the limit")
	fmt.Scanf("%d",&limit)
	for i:=1;i<=limit;i++{
		if i%2==1{
			sum+=i
		}
	}
	fmt.Println("sum is",sum)
}