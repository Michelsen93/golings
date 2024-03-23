// anonymous functions1
// Make me compile!

package main

import "fmt"

func main() {

    const llamo = "Svein"
	func(name string) {
		fmt.Printf("Hello %s", name)
	}(llamo)

}
