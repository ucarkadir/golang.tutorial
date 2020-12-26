package main

import (
	"tutorial.ground/arrays"
	"tutorial.ground/constructors"
	"tutorial.ground/fieldsofastructures"
	"tutorial.ground/ifstatements"
	"tutorial.ground/maps"
	"tutorial.ground/slices"
	"tutorial.ground/structures"
)

func main() {

	// maps - Eşlemeler
	maps.MapExOne()

	// Slices - Dilimler
	slices.SliceMethodZero()
	slices.SliceMethodOne()
	slices.SliceMethodTwo()
	slices.SliceMethodThree([]*slices.Saiyan{})
	slices.SliceMethodFour()
	slices.SliceMethodFive()

	// Diziler
	arrays.ArrayMethodOne()

	// Yapıların Alanları
	fieldsofastructures.FieldsOfaStructuresMethodOne()

	// Constructor Kurucular
	constructors.ConstructorMethodOne()

	// Structures Yapılar
	structures.StructureMethodOne()

	// if statements
	ifstatements.IfStatementMethodOne()

}
