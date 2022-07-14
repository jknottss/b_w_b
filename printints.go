package main

import (
	"bytes"
	"fmt"
)

//как fmt.Sprint() но добавляет запятые

func main() {
	fmt.Println(intsToString([]int{1, 23, 456, 789}))
}

func intsToString(val []int) string {
	var buf bytes.Buffer
	buf.WriteByte('[')
	for i, v := range val {
		if i > 0 {
			buf.WriteString(", ")
		}
		fmt.Fprintf(&buf, "%d", v)
	}
	buf.WriteByte(']')
	return buf.String()
}
