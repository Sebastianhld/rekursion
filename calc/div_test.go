package calc

import "fmt"

func ExampleDiv() {
	fmt.Println(Div(8, 4))
	fmt.Println(Div(8, 3))
	fmt.Println(Div(0, 4))
	fmt.Println(Div(4, 0))

	//Output:
	//2
	//2.6666666666666665
	//0
	//-1
}

func ExampleDivWithRest() {
	fmt.Println(DivWithRest(9, 3))
	fmt.Println(DivWithRest(9, 4))

	//Output:
	//3 0
	//2 1
}
