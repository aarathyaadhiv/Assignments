package main

import "fmt"


func main(){
	var markpercentage float64
	fmt.Println("enter the total mark percentage")

	fmt.Scanf("%f",&markpercentage)
	if markpercentage>=90{
		fmt.Println("your grade is A")
	}else if markpercentage>=80{
		fmt.Println("your grade is B")
	}else if markpercentage>=70{
		fmt.Println("your grade is C")
	}else if markpercentage>=60{
		fmt.Println("your grade is D")
	}else if markpercentage>=50{
		fmt.Println("your grade is E")
	}else{
		fmt.Println("failed")
	}
}