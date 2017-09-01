package main

import (
	"flag"
	"fmt"
	"github.com/jD91mZM2/gtable"
	"os"
	"strings"
)

type StringArr []string

func (arr *StringArr) String() string {
	return "[" + strings.Join(*arr, " ") + "]"
}
func (arr *StringArr) Set(val string) error {
	*arr = append(*arr, val)
	return nil
}

var items StringArr
var paddingLeft int
var paddingRight int
var padding int
var noHeader bool
var center bool
var centerHeader bool
var round bool

func main() {
	flag.Var(&items, "i", "Adds item to the table. Use 'br' for new line.")
	flag.IntVar(&paddingLeft, "l", 0, "Specifies the left padding on items.")
	flag.IntVar(&paddingRight, "r", 0, "Specifies the right padding on items.")
	flag.IntVar(&padding, "p", 0, "Specifies common padding. Overrides -l and -r.")
	flag.BoolVar(&noHeader, "H", false, "Removes the separator between header and body.")
	flag.BoolVar(&center, "C", false, "Centralizes all columns.")
	flag.BoolVar(&centerHeader, "c", false, "Centralizes headers. Overrides -C")
	flag.BoolVar(&round, "R", false, "Makes round corners")

	flag.Parse()

	if len(items) <= 0 {
		fmt.Fprintln(os.Stderr, "No items supplied. Use --help for help.")
		return
	}

	table := gtable.NewStringTable()
	for _, item := range items {
		if strings.EqualFold(item, "br") {
			table.AddRow()
		} else {
			table.AddStrings(item)
		}
	}

	table.Each(func(t *gtable.TableItem) {
		if padding > 0 {
			t.Padding(padding)
		} else {
			if paddingLeft > 0 {
				t.PaddingLeft = paddingLeft
			}
			if paddingRight > 0 {
				t.PaddingRight = paddingRight
			}
		}

		t.Center = center && !centerHeader
	})

	if centerHeader {
		for _, t := range table.Rows()[0] {
			t.Center = true
		}
	}

	table.Header = !noHeader
	if round {
		table.Corner = gtable.CORNER_ROUND
	}

	fmt.Println(table.String())
}
