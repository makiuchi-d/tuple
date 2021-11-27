// This package defines n-tuple structs and alternative names.
package tuple

// Couple is a 2-tuple struct.
type Couple[T1, T2 any] struct {
	V1 T1
	V2 T2
}

// Pair is a alternative name of Couple.
type Pair[T1, T2 any] Couple[T1, T2]

// Dyad is a alternative name of Couple.
type Dyad[T1, T2 any] Pair[T1, T2]

// Triple is a 3-tuple struct.
type Triple[T1, T2, T3 any] struct {
	V1 T1
	V2 T2
	V3 T3
}

// Triplet is a alternative name of Triple.
type Triplet[T1, T2, T3 any] Triple[T1, T2, T3]

// Triad is a alternative name of Triple.
type Triad[T1, T2, T3 any] Triple[T1, T2, T3]

// Quadruple is a 4-tuple struct.
type Quadruple[T1, T2, T3, T4 any] struct {
	V1 T1
	V2 T2
	V3 T3
	V4 T4
}

// Quartet is a alternative name of Quadruple.
type Quartet[T1, T2, T3, T4 any] Quadruple[T1, T2, T3, T4]

// Quad is a alternative name of Quadruple.
type Quad[T1, T2, T3, T4 any] Quadruple[T1, T2, T3, T4]

// Tetrad is a alternative name of Quadruple.
type Tetrad[T1, T2, T3, T4 any] Quadruple[T1, T2, T3, T4]

// Quintuple is a 5-tuple struct.
type Quintuple[T1, T2, T3, T4, T5 any] struct {
	V1 T1
	V2 T2
	V3 T3
	V4 T4
	V5 T5
}

// Quint is a alternative name of Quintuple
type Quint[T1, T2, T3, T4, T5 any] Quintuple[T1, T2, T3, T4, T5]

// Pentuple is a alternative name of Quintuple
type Pentuple[T1, T2, T3, T4, T5 any] Quintuple[T1, T2, T3, T4, T5]

// Pentad is a alternative name of Quintuple
type Pentad[T1, T2, T3, T4, T5 any] Quintuple[T1, T2, T3, T4, T5]

// Sextuple is a 6-tuple struct.
type Sextuple[T1, T2, T3, T4, T5, T6 any] struct {
	V1 T1
	V2 T2
	V3 T3
	V4 T4
	V5 T5
	V6 T6
}

// Hextuple is a alternative name of Sextuple
type Hextuple[T1, T2, T3, T4, T5, T6 any] Sextuple[T1, T2, T3, T4, T5, T6]

// Hexad is a alternative name of Sextuple
type Hexad[T1, T2, T3, T4, T5, T6 any] Sextuple[T1, T2, T3, T4, T5, T6]

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

// Heptuple is a alternative name of Septuple
type Heptuple[T1, T2, T3, T4, T5, T6, T7 any] Septuple[T1, T2, T3, T4, T5, T6, T7]

// Heptad is a alternative name of Septuple
type Heptad[T1, T2, T3, T4, T5, T6, T7 any] Septuple[T1, T2, T3, T4, T5, T6, T7]

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

// Octet is a alternative name of Octuple
type Octet[T1, T2, T3, T4, T5, T6, T7, T8 any] Octuple[T1, T2, T3, T4, T5, T6, T7, T8]

// Octad is a alternative name of Octuple
type Octad[T1, T2, T3, T4, T5, T6, T7, T8 any] Octuple[T1, T2, T3, T4, T5, T6, T7, T8]

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

// Nonad is a alternative name of Nonuple
type Nonad[T1, T2, T3, T4, T5, T6, T7, T8, T9 any] Nonuple[T1, T2, T3, T4, T5, T6, T7, T8, T9]

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

// Decad is a alternative name of Decuple
type Decad[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10 any] Decuple[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10]
