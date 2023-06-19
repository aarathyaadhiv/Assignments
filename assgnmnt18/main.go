package main

import "fmt"




func main(){
	var wt,le,as float32
	fmt.Println("enter the mark of written test,lab exam & assignments")
	fmt.Scanf("%f %f %f",&wt,&le,&as)
	grade:=((wt*70)/100)+((le*20)/100)+((as*10)/100)
	fmt.Println("overall grade is : ",grade)
}