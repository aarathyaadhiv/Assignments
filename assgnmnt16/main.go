package main

import "fmt"


func main(){
	var num int
	fmt.Println("enter the number")
	fmt.Scanf("%d",&num)
	flag:=0
	for i:=2;i<=num/2;i++{
		if num%i==0{
			flag=1
			break
		}
	}
	if flag==0 && num!=1{
		fmt.Println("number is prime")
	}else{
		fmt.Println("number is not a prime")
	}
}