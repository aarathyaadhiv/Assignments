package main

import "fmt"



type operations struct{
	x int
	y int
}
func (o operations) addition()int{
	return o.x+o.y
}
func (o operations) subtraction()int{
	return o.x-o.y
}
func (o operations) multiplication()int{
   return o.x*o.y
}
func (o operations) division()float32{
	return float32(o.x)/float32(o.y)
}



func main(){
 var o operations
 var num1,num2,choice int
 fmt.Println("enter 2 numbers")
 fmt.Scanf("%d %d",&num1,&num2)
 o.x=num1
 o.y=num2
 fmt.Println("1 for addition\n2 for subtraction\n3 for multiplication\n4 for division\nselect your choice")
 fmt.Scanf("%d",&choice)
 switch choice{
 case 1:
	fmt.Println("result is : ",o.addition())
 case 2:
	fmt.Println("result is :",o.subtraction())
 case 3:
	fmt.Println("result is:",o.multiplication())
 case 4:
	fmt.Println("result is : ",o.division())
 default:
	fmt.Println("enter a valid choice")


 }

 
}