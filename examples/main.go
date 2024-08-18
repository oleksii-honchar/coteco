package main

import (
	"fmt"

	c "github.com/oleksii-honchar/coteco"
)

var f = fmt.Sprintf

func main() {

	data := struct {
		key1 string
		key2 string
	}{
		key1: "value1",
		key2: "value2",
	}

	c.UseColors = true
	fmt.Println(f("Colored struct message: %s%+v%s", c.Yellow(), data, c.Reset()))
	fmt.Println(f("%sStruct message: %+v%s", c.Green(), data, c.Reset()))
	fmt.Println(f("%sStruct message: %#v%s", c.Blue(), data, c.Reset()))
	fmt.Println(c.WithGreenCyan49("green-cyan text"))

	c.UseColors = false
	fmt.Println(f("%sColors turned off%s", c.Yellow(), c.Reset()))
	fmt.Println(f("%sStruct message: %+v%s", c.Green(), data, c.Reset()))
	fmt.Println(f("%sStruct message: %#v%s", c.Blue(), data, c.Reset()))
	fmt.Println(c.WithGreenCyan49("green-cyan text"))
}
