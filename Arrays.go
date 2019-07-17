package main
import "fmt"
func main(){
	var FavNums2[5] float64
	FavNums2[0] = 163
	FavNums2[1] = 16.999
	FavNums2[2] = 16.2543
	FavNums2[3] = 16.8
	FavNums2[4] = 167
	fmt.Println(FavNums2[3])
	OtherFavNums := [5]int {1,2,9,4,5}
	fmt.Println(OtherFavNums[1])
	for i, value := range OtherFavNums{
		fmt.Println(value, i)
	}
}