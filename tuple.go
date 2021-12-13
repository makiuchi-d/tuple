// This package defines n-tuple structs and alternative names.
package tuple

// Couple is a 2-tuple struct.
type Couple[T1, T2 any] struct {
	V1 T1
	V2 T2
}

// NewCouple returns a new Couple containing v1 and v2.
func NewCouple[T1, T2 any](v1 T1, v2 T2) Couple[T1, T2] {
	return Couple[T1, T2]{v1, v2}
}

// Pair is a alternative name of Couple.
type Pair[T1, T2 any] Couple[T1, T2]

// NewPair returns a new Pair containing v1 and v2.
func NewPair[T1, T2 any](v1 T1, v2 T2) Pair[T1, T2] {
	return Pair[T1, T2]{v1, v2}
}

// Dyad is a alternative name of Couple.
type Dyad[T1, T2 any] Couple[T1, T2]

// NewDyad returns a new Dyad containing v1 and v2.
func NewDyad[T1, T2 any](v1 T1, v2 T2) Dyad[T1, T2] {
	return Dyad[T1, T2]{v1, v2}
}

// Triple is a 3-tuple struct.
type Triple[T1, T2, T3 any] struct {
	V1 T1
	V2 T2
	V3 T3
}

// NewTriple returns a new Triple containing v1, v2 and v3.
func NewTriple[T1, T2, T3 any](v1 T1, v2 T2, v3 T3) Triple[T1, T2, T3] {
	return Triple[T1, T2, T3]{v1, v2, v3}
}

// Triplet is a alternative name of Triple.
type Triplet[T1, T2, T3 any] Triple[T1, T2, T3]

// NewTriplet returns a new Triplet containing v1, v2 and v3.
func NewTriplet[T1, T2, T3 any](v1 T1, v2 T2, v3 T3) Triplet[T1, T2, T3] {
	return Triplet[T1, T2, T3]{v1, v2, v3}
}

// Triad is a alternative name of Triple.
type Triad[T1, T2, T3 any] Triple[T1, T2, T3]

// NewTriad returns a new Triad containing v1, v2 and v3.
func NewTriad[T1, T2, T3 any](v1 T1, v2 T2, v3 T3) Triad[T1, T2, T3] {
	return Triad[T1, T2, T3]{v1, v2, v3}
}

// Quadruple is a 4-tuple struct.
type Quadruple[T1, T2, T3, T4 any] struct {
	V1 T1
	V2 T2
	V3 T3
	V4 T4
}

// NewQuadruple returns a new Quadruple containing v1, v2, v3 and v4.
func NewQuadruple[T1, T2, T3, T4 any](v1 T1, v2 T2, v3 T3, v4 T4) Quadruple[T1, T2, T3, T4] {
	return Quadruple[T1, T2, T3, T4]{v1, v2, v3, v4}
}

// Quartet is a alternative name of Quadruple.
type Quartet[T1, T2, T3, T4 any] Quadruple[T1, T2, T3, T4]

// NewQuartet returns a new Quartet containing v1, v2, v3 and v4.
func NewQuartet[T1, T2, T3, T4 any](v1 T1, v2 T2, v3 T3, v4 T4) Quartet[T1, T2, T3, T4] {
	return Quartet[T1, T2, T3, T4]{v1, v2, v3, v4}
}

// Quad is a alternative name of Quadruple.
type Quad[T1, T2, T3, T4 any] Quadruple[T1, T2, T3, T4]

// NewQuad returns a new Quad containing v1, v2, v3 and v4.
func NewQuad[T1, T2, T3, T4 any](v1 T1, v2 T2, v3 T3, v4 T4) Quad[T1, T2, T3, T4] {
	return Quad[T1, T2, T3, T4]{v1, v2, v3, v4}
}

// Tetrad is a alternative name of Quadruple.
type Tetrad[T1, T2, T3, T4 any] Quadruple[T1, T2, T3, T4]

// NewTetrad returns a new Tetrad containing v1, v2, v3 and v4.
func NewTetrad[T1, T2, T3, T4 any](v1 T1, v2 T2, v3 T3, v4 T4) Tetrad[T1, T2, T3, T4] {
	return Tetrad[T1, T2, T3, T4]{v1, v2, v3, v4}
}

// Quintuple is a 5-tuple struct.
type Quintuple[T1, T2, T3, T4, T5 any] struct {
	V1 T1
	V2 T2
	V3 T3
	V4 T4
	V5 T5
}

// NewQuintuple returns a new Quintuple containing v1, v2, v3, v4 and v5.
func NewQuintuple[T1, T2, T3, T4, T5 any](v1 T1, v2 T2, v3 T3, v4 T4, v5 T5) Quintuple[T1, T2, T3, T4, T5] {
	return Quintuple[T1, T2, T3, T4, T5]{v1, v2, v3, v4, v5}
}

// Quint is a alternative name of Quintuple.
type Quint[T1, T2, T3, T4, T5 any] Quintuple[T1, T2, T3, T4, T5]

// NewQuint returns a new Quint containing v1, v2, v3, v4 and v5.
func NewQuint[T1, T2, T3, T4, T5 any](v1 T1, v2 T2, v3 T3, v4 T4, v5 T5) Quint[T1, T2, T3, T4, T5] {
	return Quint[T1, T2, T3, T4, T5]{v1, v2, v3, v4, v5}
}

// Pentuple is a alternative name of Quintuple.
type Pentuple[T1, T2, T3, T4, T5 any] Quintuple[T1, T2, T3, T4, T5]

// NewPentuple returns a new Pentuple containing v1, v2, v3, v4 and v5.
func NewPentuple[T1, T2, T3, T4, T5 any](v1 T1, v2 T2, v3 T3, v4 T4, v5 T5) Pentuple[T1, T2, T3, T4, T5] {
	return Pentuple[T1, T2, T3, T4, T5]{v1, v2, v3, v4, v5}
}

// Pentad is a alternative name of Quintuple.
type Pentad[T1, T2, T3, T4, T5 any] Quintuple[T1, T2, T3, T4, T5]

// NewPentad returns a new Pentad containing v1, v2, v3, v4 and v5.
func NewPentad[T1, T2, T3, T4, T5 any](v1 T1, v2 T2, v3 T3, v4 T4, v5 T5) Pentad[T1, T2, T3, T4, T5] {
	return Pentad[T1, T2, T3, T4, T5]{v1, v2, v3, v4, v5}
}

// Sextuple is a 6-tuple struct.
type Sextuple[T1, T2, T3, T4, T5, T6 any] struct {
	V1 T1
	V2 T2
	V3 T3
	V4 T4
	V5 T5
	V6 T6
}

// NewSextuple returns a new Sextuple containing v1, v2, v3, v4, v5 and v6.
func NewSextuple[T1, T2, T3, T4, T5, T6 any](v1 T1, v2 T2, v3 T3, v4 T4, v5 T5, v6 T6) Sextuple[T1, T2, T3, T4, T5, T6] {
	return Sextuple[T1, T2, T3, T4, T5, T6]{v1, v2, v3, v4, v5, v6}
}

// Hextuple is a alternative name of Sextuple.
type Hextuple[T1, T2, T3, T4, T5, T6 any] Sextuple[T1, T2, T3, T4, T5, T6]

// NewHextuple returns a new Hextuple containing v1, v2, v3, v4, v5 and v6.
func NewHextuple[T1, T2, T3, T4, T5, T6 any](v1 T1, v2 T2, v3 T3, v4 T4, v5 T5, v6 T6) Hextuple[T1, T2, T3, T4, T5, T6] {
	return Hextuple[T1, T2, T3, T4, T5, T6]{v1, v2, v3, v4, v5, v6}
}

// Hexad is a alternative name of Sextuple.
type Hexad[T1, T2, T3, T4, T5, T6 any] Sextuple[T1, T2, T3, T4, T5, T6]

// NewHexad returns a new Hexad containing v1, v2, v3, v4, v5 and v6.
func NewHexad[T1, T2, T3, T4, T5, T6 any](v1 T1, v2 T2, v3 T3, v4 T4, v5 T5, v6 T6) Hexad[T1, T2, T3, T4, T5, T6] {
	return Hexad[T1, T2, T3, T4, T5, T6]{v1, v2, v3, v4, v5, v6}
}

// Septuple is a 7-tuple struct.
type Septuple[T1, T2, T3, T4, T5, T6, T7 any] struct {
	V1 T1
	V2 T2
	V3 T3
	V4 T4
	V5 T5
	V6 T6
	V7 T7
}

// NewSeptuple returns a new Septuple containing v1, v2, v3, v4, v5, v6 and v7.
func NewSeptuple[T1, T2, T3, T4, T5, T6, T7 any](v1 T1, v2 T2, v3 T3, v4 T4, v5 T5, v6 T6, v7 T7) Septuple[T1, T2, T3, T4, T5, T6, T7] {
	return Septuple[T1, T2, T3, T4, T5, T6, T7]{v1, v2, v3, v4, v5, v6, v7}
}

// Heptuple is a alternative name of Septuple.
type Heptuple[T1, T2, T3, T4, T5, T6, T7 any] Septuple[T1, T2, T3, T4, T5, T6, T7]

// NewHeptuple returns a new Heptuple containing v1, v2, v3, v4, v5, v6 and v7.
func NewHeptuple[T1, T2, T3, T4, T5, T6, T7 any](v1 T1, v2 T2, v3 T3, v4 T4, v5 T5, v6 T6, v7 T7) Heptuple[T1, T2, T3, T4, T5, T6, T7] {
	return Heptuple[T1, T2, T3, T4, T5, T6, T7]{v1, v2, v3, v4, v5, v6, v7}
}

// Heptad is a alternative name of Septuple.
type Heptad[T1, T2, T3, T4, T5, T6, T7 any] Septuple[T1, T2, T3, T4, T5, T6, T7]

// NewHeptad returns a new Heptad containing v1, v2, v3, v4, v5, v6 and v7.
func NewHeptad[T1, T2, T3, T4, T5, T6, T7 any](v1 T1, v2 T2, v3 T3, v4 T4, v5 T5, v6 T6, v7 T7) Heptad[T1, T2, T3, T4, T5, T6, T7] {
	return Heptad[T1, T2, T3, T4, T5, T6, T7]{v1, v2, v3, v4, v5, v6, v7}
}

// Octuple is a 8-tuple struct.
type Octuple[T1, T2, T3, T4, T5, T6, T7, T8 any] struct {
	V1 T1
	V2 T2
	V3 T3
	V4 T4
	V5 T5
	V6 T6
	V7 T7
	V8 T8
}

// NewOctuple returns a new Octuple containing v1, v2, v3, v4, v5, v6, v7 and v8.
func NewOctuple[T1, T2, T3, T4, T5, T6, T7, T8 any](v1 T1, v2 T2, v3 T3, v4 T4, v5 T5, v6 T6, v7 T7, v8 T8) Octuple[T1, T2, T3, T4, T5, T6, T7, T8] {
	return Octuple[T1, T2, T3, T4, T5, T6, T7, T8]{v1, v2, v3, v4, v5, v6, v7, v8}
}

// Octet is a alternative name of Octuple.
type Octet[T1, T2, T3, T4, T5, T6, T7, T8 any] Octuple[T1, T2, T3, T4, T5, T6, T7, T8]

// NewOctet returns a new Octet containing v1, v2, v3, v4, v5, v6, v7 and v8.
func NewOctet[T1, T2, T3, T4, T5, T6, T7, T8 any](v1 T1, v2 T2, v3 T3, v4 T4, v5 T5, v6 T6, v7 T7, v8 T8) Octet[T1, T2, T3, T4, T5, T6, T7, T8] {
	return Octet[T1, T2, T3, T4, T5, T6, T7, T8]{v1, v2, v3, v4, v5, v6, v7, v8}
}

// Octad is a alternative name of Octuple.
type Octad[T1, T2, T3, T4, T5, T6, T7, T8 any] Octuple[T1, T2, T3, T4, T5, T6, T7, T8]

// NewOctad returns a new Octad containing v1, v2, v3, v4, v5, v6, v7 and v8.
func NewOctad[T1, T2, T3, T4, T5, T6, T7, T8 any](v1 T1, v2 T2, v3 T3, v4 T4, v5 T5, v6 T6, v7 T7, v8 T8) Octad[T1, T2, T3, T4, T5, T6, T7, T8] {
	return Octad[T1, T2, T3, T4, T5, T6, T7, T8]{v1, v2, v3, v4, v5, v6, v7, v8}
}

// Nonuple is a 9-tuple struct.
type Nonuple[T1, T2, T3, T4, T5, T6, T7, T8, T9 any] struct {
	V1 T1
	V2 T2
	V3 T3
	V4 T4
	V5 T5
	V6 T6
	V7 T7
	V8 T8
	V9 T9
}

// NewNonuple returns a new Nonuple containing v1, v2, v3, v4, v5, v6, v7, v8 and v9.
func NewNonuple[T1, T2, T3, T4, T5, T6, T7, T8, T9 any](v1 T1, v2 T2, v3 T3, v4 T4, v5 T5, v6 T6, v7 T7, v8 T8, v9 T9) Nonuple[T1, T2, T3, T4, T5, T6, T7, T8, T9] {
	return Nonuple[T1, T2, T3, T4, T5, T6, T7, T8, T9]{v1, v2, v3, v4, v5, v6, v7, v8, v9}
}

// Nonad is a alternative name of Nonuple.
type Nonad[T1, T2, T3, T4, T5, T6, T7, T8, T9 any] Nonuple[T1, T2, T3, T4, T5, T6, T7, T8, T9]

// NewNonad returns a new Nonad containing v1, v2, v3, v4, v5, v6, v7, v8 and v9.
func NewNonad[T1, T2, T3, T4, T5, T6, T7, T8, T9 any](v1 T1, v2 T2, v3 T3, v4 T4, v5 T5, v6 T6, v7 T7, v8 T8, v9 T9) Nonad[T1, T2, T3, T4, T5, T6, T7, T8, T9] {
	return Nonad[T1, T2, T3, T4, T5, T6, T7, T8, T9]{v1, v2, v3, v4, v5, v6, v7, v8, v9}
}

// Decuple is a 10-tuple struct.
type Decuple[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10 any] struct {
	V1  T1
	V2  T2
	V3  T3
	V4  T4
	V5  T5
	V6  T6
	V7  T7
	V8  T8
	V9  T9
	V10 T10
}

// NewDecuple returns a new Decuple containing v1, v2, v3, v4, v5, v6, v7, v8, v9 and v10.
func NewDecuple[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10 any](v1 T1, v2 T2, v3 T3, v4 T4, v5 T5, v6 T6, v7 T7, v8 T8, v9 T9, v10 T10) Decuple[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10] {
	return Decuple[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10]{v1, v2, v3, v4, v5, v6, v7, v8, v9, v10}
}

// Decad is a alternative name of Decuple.
type Decad[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10 any] Decuple[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10]

// NewDecad returns a new Decad containing v1, v2, v3, v4, v5, v6, v7, v8, v9 and v10.
func NewDecad[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10 any](v1 T1, v2 T2, v3 T3, v4 T4, v5 T5, v6 T6, v7 T7, v8 T8, v9 T9, v10 T10) Decad[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10] {
	return Decad[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10]{v1, v2, v3, v4, v5, v6, v7, v8, v9, v10}
}
