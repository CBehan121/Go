package main
import "fmt"
func main(){
	
	var age int = 40
	var floatyBoi float64 = 1.98495
	fmt.Println(age,floatyBoi)
	randNum := 1
	var randNum2 = 1
	fmt.Println(randNum2,randNum)
	const pi float64 = 3.14159265
	var(
		VarA = 2
		VarB = 3
		)
	var myName string = "Conner"
	fmt.Println(len(myName), VarA, VarB)
	fmt.Println(myName + " is a robot \n \n")
	//var Isover50 bool = true
	fmt.Printf("%.3f \n", pi)
	fmt.Printf("%T", myName)

}