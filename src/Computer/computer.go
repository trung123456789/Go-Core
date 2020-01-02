package computer

// SpecExported struct
type SpecExported struct { //exported struct
	Maker string //exported field
	model string //unexported field
	Price int    //exported field
}
