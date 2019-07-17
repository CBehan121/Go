package main
import "fmt"
func main(){
	numSlice := []int {5,4,3,2,1,1,1,1}
	numSlice2 := numSlice[3:5]
	fmt.Println(numSlice2[0])
	numSlice3 := append(numSlice,-1)
	fmt.Println(numSlice3)
}



