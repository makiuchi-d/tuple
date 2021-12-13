package main

import (
	"fmt"
	"strings"
)

func main() {
	defs := []struct {
		n    int
		name string
	}{
		{2, "Couple"},
		{2, "Pair"},
		{2, "Dyad"},
		{3, "Triple"},
		{3, "Triplet"},
		{3, "Triad"},
		{4, "Quadruple"},
		{4, "Quartet"},
		{4, "Quad"},
		{4, "Tetrad"},
		{5, "Quintuple"},
		{5, "Quint"},
		{5, "Pentuple"},
		{5, "Pentad"},
		{6, "Sextuple"},
		{6, "Hextuple"},
		{6, "Hexad"},
		{7, "Septuple"},
		{7, "Heptuple"},
		{7, "Heptad"},
		{8, "Octuple"},
		{8, "Octet"},
		{8, "Octad"},
		{9, "Nonuple"},
		{9, "Nonad"},
		{10, "Decuple"},
		{10, "Decad"},
	}

	fmt.Println("// This package defines n-tuple structs and alternative names.")
	fmt.Println("package tuple")

	var pn int
	var pname string
	for _, d := range defs {
		fmt.Println("")
		types := gen1(d.n, "T%d", ", ")

		if d.n != pn {
			pn = d.n
			pname = d.name
			fmt.Printf("// %s is a %d-tuple struct.\n", d.name, d.n)
			fmt.Printf("type %s[%s any] struct {\n", d.name, types)
			for i := 1; i <= d.n; i++ {
				fmt.Printf("\tV%d T%d\n", i, i)
			}
			fmt.Println("}")

		} else {
			fmt.Printf("// %s is a alternative name of %s.\n", d.name, pname)
			fmt.Printf("type %s[%s any] %s[%s]\n", d.name, types, pname, types)
		}

		params := gen2(d.n, "v%d T%d", ", ")
		vals := gen1(d.n, "v%d", ", ")
		fmt.Println("")
		fmt.Printf("// New%s returns a new %s containing %s and v%d.\n", d.name, d.name, gen1(d.n-1, "v%d", ", "), d.n)
		fmt.Printf("func New%s[%s any](%s) %s[%s] {\n", d.name, types, params, d.name, types)
		fmt.Printf("\treturn %s[%s]{%s}\n", d.name, types, vals)
		fmt.Println("}")
	}
}

func gen1(n int, pat, sep string) string {
	elems := make([]string, n)
	for i := 0; i < n; i++ {
		elems[i] = fmt.Sprintf(pat, i+1)
	}
	return strings.Join(elems, sep)
}

func gen2(n int, pat, sep string) string {
	elems := make([]string, n)
	for i := 0; i < n; i++ {
		elems[i] = fmt.Sprintf(pat, i+1, i+1)
	}
	return strings.Join(elems, sep)
}
