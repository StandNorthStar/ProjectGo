package main

import (
	"bytes"
	"fmt"
)

func main() {

	var buf bytes.Buffer

	buf.WriteString("test01")
	buf.WriteRune(';')
	buf.WriteString("t80-098;")
	fmt.Println(buf.String())
}
