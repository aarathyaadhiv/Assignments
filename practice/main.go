package main

import "fmt"

type student struct{
	name string
	age int
	grades []int
}
func(s student) average()float32{
	 sum:=0
	for i:=range s.grades{
		sum+=s.grades[i]
	}
	return float32(sum/len(s.grades))
}
func call(){
	for i:=0;i<5;i++{
		defer fmt.Println(i)
	}
}


func main(){
	var s student =student{"ach",10,[]int{1,2,3}}
	fmt.Println(s)
	defer fmt.Println(s.name)
	fmt.Println(s.average())
	x:=10
	defer fmt.Println(x)
	fmt.Println(&x)
	x=12
	var y *int=&x
	
	*y=20
	defer fmt.Println(y)
	
	fmt.Println(*y)
	fmt.Println(x)
	call()
	count:=0
	for count<10{
		
		count++
		if count==2{
			
			break
		}else if count==9{
			goto an
		}
		fmt.Println(count)
	}
	an:
	fmt.Println("hi aarathy")

}