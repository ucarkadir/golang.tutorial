package main

import (
	"fmt"

	"tutorial.ground/arrays"
	"tutorial.ground/concurrencies"
	"tutorial.ground/constructors"
	"tutorial.ground/fieldsofastructures"
	"tutorial.ground/ifstatements"
	"tutorial.ground/interfaces"
	"tutorial.ground/maps"
	"tutorial.ground/shopping"
	"tutorial.ground/slices"
	"tutorial.ground/structures"
	"tutorial.ground/usefulinformations"
)

func main() {

	// Concurrencie - eş zamanlılık
	concurrencies.GoRoutineMutex()
	concurrencies.GoRoutineForLoop()

	// Userful Information Ex
	usefulinformations.FunctionType()
	usefulinformations.ByteArray()
	usefulinformations.DeferMethod()

	// Error Handling - Hata Yönetimi
	usefulinformations.EOFMethod()
	usefulinformations.Process(0)
	usefulinformations.ErrorHandling()

	// Interfaces - Arayüzler
	var logger interfaces.Logger
	logger.Log("Hata")
	interfaces.Process(logger)

	// Dosya yapısı
	fmt.Println(shopping.PriceCheck(4343))

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
