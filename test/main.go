package main

import "fmt"

func main(){
	// go lang output statement 
	fmt.Print("hello world")
	fmt.Println("this is new line")
	// go variables 

	var age = 90
	var name = "ABDURHMAN"

	fmt.Println("my name is ", name , "and i am ", age , "years old")


	// var fName string
	// var price int 
	// fName = "Redi"
	// price = 50

	// fmt.Println("my father name is ", fName)
	// fmt.Println("monthly price is", price)

	//go receive input from the user  
	// var fName string
	// var price int 

	// fmt.Println("Please input your father name")
	// fmt.Scan(&fName)
	// fmt.Println("your father name is ", fName)


	// go arithmetic operations

	// var num1 int
	// var num2 int
	// var result int 
	// fmt.Println("Arithimetic Operations")
	// fmt.Println("Input the first number ")
	// fmt.Scan(&num1)
	// fmt.Println("Input the second number")
	// fmt.Scan(&num2)
	// result = num1 + num2
	// fmt.Println("The sum is ", result)


	//go Array 
	var nums [5] int
	var i int
	for  i=0;i<5;i++{
		fmt.Scan(&nums[i])
	}
	fmt.Println("the numbers are")
	for i=0;i<5;i++{
		fmt.Println(nums[i])
	}
	
	// nums[0] = 5
	// fmt.Println("the first index is", nums[0])


}