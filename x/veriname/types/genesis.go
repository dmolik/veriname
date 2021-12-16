package types

import (
	"fmt"
)

// DefaultIndex is the default capability global index
const DefaultIndex uint64 = 1

// DefaultGenesis returns the default Capability genesis state
func DefaultGenesis() *GenesisState {
	return &GenesisState{
		IdentList: []Ident{},
		// this line is used by starport scaffolding # genesis/types/default
	}
}

// Validate performs basic genesis state validation returning an error upon any
// failure.
func (gs GenesisState) Validate() error {
	// Check for duplicated index in ident
	identIndexMap := make(map[string]struct{})

	for _, elem := range gs.IdentList {
		index := string(IdentKey(elem.Index))
		if _, ok := identIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for ident")
		}
		identIndexMap[index] = struct{}{}
	}
	// this line is used by starport scaffolding # genesis/types/validate

	return nil
}
