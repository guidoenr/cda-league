package model

import "encoding/json"

// Position ------------------------------------------
const (
	Volante   Position = "volante"
	Delantero Position = "delantero"
	Defensor  Position = "defensor"
)

type Position string

// UnmarshalJSON give us the capability of unmarshal a non-string field into or Player struct
func (r *Position) UnmarshalJSON(data []byte) error {
	// Parse the data as an integer
	var position string
	err := json.Unmarshal(data, &position)
	if err != nil {
		return err
	}

	// Set the value of r
	*r = Position(position)
	return nil
}

// Rank ------------------------------------------
const (
	Five  Rank = 5
	Four  Rank = 4
	Three Rank = 3
	Two   Rank = 2
	One   Rank = 1
)

type Rank int

// UnmarshalJSON give us the capability of unmarshal a non-int field into or Player struct
func (r *Rank) UnmarshalJSON(data []byte) error {
	// Parse the data as an integer
	var rank int
	err := json.Unmarshal(data, &rank)
	if err != nil {
		return err
	}

	// Set the value of r
	*r = Rank(rank)
	return nil
}
