// slices2
// Make me compile!

package main

import "fmt"

func main() {
	var names = [4]string{"John", "Maria", "Carl", "Peter"}
    lastTwoNames := names[2:4] // after figuring out the answer, try with other low/high bounds
	fmt.Println(lastTwoNames)
}
