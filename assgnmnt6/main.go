package main

import "fmt"



func main(){

	var number int
	fmt.Println("1 for sunday\n2 for monday \n3 for tuesday\n4 for wednesday\n5 for thursday\n6 for friday\n7 for saturday\nselect a number")
	fmt.Scanf("%d",&number)

	switch number{
	case 1:
		fmt.Println("sunday")
	case 2:
		fmt.Println("monday")
	case 3:
		fmt.Println("tuesday")
	case 4:
		fmt.Println("wednesday")
	case 5:
		fmt.Println("thursday")
	case 6:
		fmt.Println("friday")
	case 7:
		fmt.Println("saturday")
	default:
		fmt.Println("invalid entry")
	}
}