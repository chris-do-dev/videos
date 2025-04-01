package main

import (
	"fmt"
	"go/types"
)

func main() {
	terms := []*types.Term{
		types.NewTerm(false, types.Typ[types.Int]),
		types.NewTerm(false, types.Typ[types.String]),
		types.NewTerm(false, types.Typ[types.Bool]),
	}

	union := types.NewUnion(terms)

	fmt.Println("Union Terms:")
	for term := range union.Terms() {
		fmt.Printf("- %v\n", term.Type())
	}
}
