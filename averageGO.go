//print average number
package main
import "fmt"
func main(){
	Single := 0
	Log := 0
	Vari  := [6]int {1,2,3,4,5,6}
	for i, Value := range Vari{
		Single += Value
		Log = i
	}
	fmt.Println(Single/Log)
	
}