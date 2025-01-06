package golang_saying

import "fmt"

func Saying(saying string, isAdmin bool) {
	fmt.Println("Hai kak,Hayo mau ngomong apa")
	fmt.Println(saying)

	fmt.Println(Filter(saying))
}
