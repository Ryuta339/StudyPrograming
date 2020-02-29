/**
 * Reference
 * https://astaxie.gitbooks.io/build-web-application-with-golang/ja/02.3.html
 */

package main

import "fmt"

func main () {
	if x := computedValue(); x > 10 {
		fmt.Println ("x is greater than 10");
	} else {
		fmt.Println ("x is not greater than 10");
	}
}
