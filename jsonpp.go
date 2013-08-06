package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"flag"
	"io"
	"log"
	"os"
	"strings"
)

var help = flag.Bool("help", false, "Print this help message")
var tabs = flag.Bool("tabs", false, "Use tabs to indent")
var indentSize = flag.Int("indent", 2, "Number of spaces to use for indent. Use 0 to print JSON on single line")

var indent string

func main() {
	flag.Parse()

	if *help {
		flag.PrintDefaults()
		os.Exit(0)
	}

	var out = bufio.NewWriter(os.Stdout)
	defer out.Flush()

	if *tabs {
		indent = "\t"
	} else {
		indent = strings.Repeat(" ", *indentSize)
	}

	if len(flag.Args()) == 0 {
		formatJson(bufio.NewReader(os.Stdin), out)
	} else {
		for _, filename := range flag.Args() {
			f, err := os.Open(filename)
			defer f.Close()

			if err != nil {
				log.Println(err)
				continue
			}

			formatJson(bufio.NewReader(f), out)

			out.WriteRune('\n')
		}
	}
}

func formatJson(in io.Reader, out io.Writer) {
	reader := new(bytes.Buffer)

	_, err := reader.ReadFrom(in)
	if err != nil {
		log.Print(err)
	}

	b := bytes.TrimSpace(reader.Bytes())

	w := new(bytes.Buffer)

	if indent == "" {
		err = json.Compact(w, b)
	} else {
		err = json.Indent(w, b, "", indent)
	}

	if err != nil {
		log.Print(err)
		return
	}

	_, err = w.WriteTo(out)
	if err != nil {
		log.Print(err)
		return
	}
}
