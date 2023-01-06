package model

type Position string

const (
	Volante   Position = "volante"
	Delantero Position = "delantero"
	Defensor  Position = "defensor"
)

type Rank int

const (
	Five  Rank = 5
	Four  Rank = 4
	Three Rank = 3
	Two   Rank = 2
	One   Rank = 1
)
