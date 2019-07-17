package main
import "fmt"
func main(){
	i := 1
	for i <= 5{
		fmt.Println(i)
		i++

	}
	//OR
	for j := 1; j <= 5; j++{
		fmt.Println(j)
	}
	age := 16
	if age > 16{
		fmt.Println("You cant drive")
	} else if age <= 18{
		fmt.Println("You can vote")
	} else {
		fmt.Println(" you can drive")
	}
	switch age{
		case 14 : fmt.Println("your 14")
		case 16 : fmt.Println("your 16")
		default : fmt.Println("something")

	}
}