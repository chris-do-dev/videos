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
	for i := 0; i < union.Len(); i++ {
		fmt.Printf("- %v\n", union.Term(i).Type())
	}
}
