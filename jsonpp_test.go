package main

import (
	"bytes"
	"fmt"
)

func ExampleCompact() {
	in := bytes.NewReader([]byte(`{   "hello":"world", "count":    123 }`))
	out := bytes.NewBuffer([]byte(""))

	formatJson(in, out)

	fmt.Print(out)

	// Output:
	// {"hello":"world","count":123}
}

func ExampleIndent2() {
	in := bytes.NewReader([]byte(`{   "hello":"world", "count":    123 }`))
	out := bytes.NewBuffer([]byte(""))
	indent = "  "

	formatJson(in, out)

	fmt.Print(out)

	// Output:
	// {
	//   "hello": "world",
	//   "count": 123
	// }
}

func ExampleIndent4() {
	in := bytes.NewReader([]byte(`{   "count":    123,  "hello":"world" }`))
	out := bytes.NewBuffer([]byte(""))
	indent = "    "

	formatJson(in, out)

	fmt.Print(out)

	// Output:
	// {
	//     "count": 123,
	//     "hello": "world"
	// }
}

func ExampleIndentTabs() {
	in := bytes.NewReader([]byte(`{   "count":    123,  "hello":"world" }`))
	out := bytes.NewBuffer([]byte(""))
	indent = "\t"

	formatJson(in, out)

	fmt.Print(out)

	// Output:
	// {
	// 	"count": 123,
	// 	"hello": "world"
	// }
}
