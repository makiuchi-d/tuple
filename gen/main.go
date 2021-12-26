package main

import (
	"bytes"
	"fmt"
	"go/format"
	"os"
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

	w := bytes.NewBuffer(nil)

	fmt.Fprintln(w, "// This package defines n-tuple structs and alternative names.")
	fmt.Fprintln(w, "package tuple")

	var pn int
	var pname string
	for _, d := range defs {
		fmt.Fprintln(w, "")
		types := gen1(d.n, "T%d", ", ")

		if d.n != pn {
			pn = d.n
			pname = d.name
			fmt.Fprintf(w, "// %s is a %d-tuple struct.\n", d.name, d.n)
			fmt.Fprintf(w, "type %s[%s any] struct {\n", d.name, types)
			for i := 1; i <= d.n; i++ {
				fmt.Fprintf(w, "\tV%d T%d\n", i, i)
			}
			fmt.Fprintln(w, "}")

		} else {
			fmt.Fprintf(w, "// %s is a alternative name of %s.\n", d.name, pname)
			fmt.Fprintf(w, "type %s[%s any] %s[%s]\n", d.name, types, pname, types)
		}

		params := gen2(d.n, "v%d T%d", ", ")
		vals := gen1(d.n, "v%d", ", ")
		fmt.Fprintln(w, "")
		fmt.Fprintf(w, "// New%s returns a new %s containing %s and v%d.\n", d.name, d.name, gen1(d.n-1, "v%d", ", "), d.n)
		fmt.Fprintf(w, "func New%s[%s any](%s) %s[%s] {\n", d.name, types, params, d.name, types)
		fmt.Fprintf(w, "\treturn %s[%s]{%s}\n", d.name, types, vals)
		fmt.Fprintln(w, "}")
	}

	out, err := format.Source(w.Bytes())
	if err != nil {
		panic(err)
	}

	os.Stdout.Write(out)
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
