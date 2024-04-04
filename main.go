package main

import (
	"fmt"
	solid "golang/designprinciplesandpatterns/solid/ocp"
)

//SRP
// func main() {
// 	fmt.Println("SRP implementation test")
// 	note := solid.NoteSRP{Lines: make(map[int]string), LastLine: 1}
// 	note.AddLine("I am lazy")
// 	note.AddLine("Oh no devin will replace me :-(")
// 	fmt.Println(note.PrintLines("\n"))

// 	noteFile := solid.NoteFile{LineSeparator: "\r\n"}
// 	noteFile.PersistNote(&note, "test.txt")
// }

// ******************************************OCP*******************************************************************
func main() {
	fmt.Println("--------------Failed ocp run first-----------------------------")
	shirts := []solid.Shirt{
		{Color: solid.Black, Size: solid.Large, Mrp: 1300.00},
		{Color: solid.Yellow, Size: solid.Large, Mrp: 1250.00},
		{Color: solid.Blue, Size: solid.Small, Mrp: 1200.00},
		{Color: solid.Black, Size: solid.Small, Mrp: 1300.00},
		{Color: solid.White, Size: solid.Medium, Mrp: 1500.00},
		{Color: solid.Green, Size: solid.Medium, Mrp: 1400.00},
	}

	filter := &solid.ShirtFilter{}

	fmt.Println("black shirts:")
	blackShirts := filter.FilterByColor(shirts, solid.Black)
	solid.PrintShirts(blackShirts)
	fmt.Println("small shirts:")
	smallShirts := filter.FilterBySize(shirts, solid.Small)
	solid.PrintShirts(smallShirts)

	fmt.Println("--------------Success ocp run first-----------------------------")

	// extended Predicate function for black shirt
	blackShirtsFilter := func(shirt solid.Shirt) bool { return shirt.Color == solid.Black }
	// extended Predicate function for large shirt
	largeShirtsFilter := func(shirt solid.Shirt) bool { return shirt.Size == solid.Large }
	// extended Predicate function for black large shirt
	largeBlackShirtFilter := func(shirt solid.Shirt) bool { return blackShirtsFilter(shirt) && largeShirtsFilter(shirt) }

	fmt.Println("Black shirts:")
	blackShirts2 := solid.Filter(shirts, blackShirtsFilter)
	solid.PrintShirtsOCP(blackShirts2)
	fmt.Println("Large shirts:")
	largeShirts2 := solid.Filter(shirts, largeShirtsFilter)
	solid.PrintShirtsOCP(largeShirts2)
	fmt.Println("Large black shirts:")
	largeBlackShirts := solid.Filter(shirts, largeBlackShirtFilter)
	solid.PrintShirtsOCP(largeBlackShirts)
}
