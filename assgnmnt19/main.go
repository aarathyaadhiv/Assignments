package main

import "fmt"



func main(){
	var income float32
	var tax float32
  fmt.Println("enter your annual income")
  fmt.Scanf("%f",&income)
  if income<=250000{
	tax=0
  }else if income<500000{
	tax=(income*5)/100
  }else if income<1000000{
	tax=(income*20)/100
  }else if income<5000000{
	tax=(income*30)/100
  }else{
	fmt.Println("no tax structure for your income")
  }
  fmt.Println("your tax is : ",tax)
}