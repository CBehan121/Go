package main
import "fmt"
func main(){
	nameAge := make(map[string] int)
	nameAge["Conner"] = 20
	fmt.Println(nameAge["Conner"])
	fmt.Println(len(nameAge))
	delete(nameAge, "Conner")
}