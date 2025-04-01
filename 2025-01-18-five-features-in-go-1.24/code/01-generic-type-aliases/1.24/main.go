package main

import (
	"fmt"
	"reflect"
)

type StatusDefinition[T comparable] []T

type StatusAlias[T comparable] = []T

func main() {
	var def StatusDefinition[int]
	def = append(def, 1)
	fmt.Println("Type of StatusDefinition: ", reflect.TypeOf(def))

	var alias StatusAlias[int]
	alias = append(alias, 1)
	fmt.Println("Type of StatusAlias: ", reflect.TypeOf(alias))
}
