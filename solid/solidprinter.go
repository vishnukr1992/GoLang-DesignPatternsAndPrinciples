package solid

import (
	"fmt"
	isp "golang/designprinciplesandpatterns/solid/isp"
	lsp "golang/designprinciplesandpatterns/solid/lsp"
	ocp "golang/designprinciplesandpatterns/solid/ocp"
	srp "golang/designprinciplesandpatterns/solid/srp"
	"time"
)

func PrintSRP() {
	fmt.Println("SRP implementation test")
	note := srp.NoteSRP{Lines: make(map[int]string), LastLine: 1}
	note.AddLine("I am lazy")
	note.AddLine("Oh no devin will replace me :-(")
	fmt.Println(note.PrintLines("\n"))

	noteFile := srp.NoteFile{LineSeparator: "\r\n"}
	noteFile.PersistNote(&note, "test.txt")
}

func PrintOCP() {
	fmt.Println("--------------Failed ocp run first-----------------------------")
	shirts := []ocp.Shirt{
		{Color: ocp.Black, Size: ocp.Large, Mrp: 1300.00},
		{Color: ocp.Yellow, Size: ocp.Large, Mrp: 1250.00},
		{Color: ocp.Blue, Size: ocp.Small, Mrp: 1200.00},
		{Color: ocp.Black, Size: ocp.Small, Mrp: 1300.00},
		{Color: ocp.White, Size: ocp.Medium, Mrp: 1500.00},
		{Color: ocp.Green, Size: ocp.Medium, Mrp: 1400.00},
	}

	filter := &ocp.ShirtFilter{}

	fmt.Println("black shirts:")
	blackShirts := filter.FilterByColor(shirts, ocp.Black)
	ocp.PrintShirts(blackShirts)
	fmt.Println("small shirts:")
	smallShirts := filter.FilterBySize(shirts, ocp.Small)
	ocp.PrintShirts(smallShirts)

	fmt.Println("--------------Success ocp run first-----------------------------")

	// extended Predicate function for black shirt
	blackShirtsFilter := func(shirt ocp.Shirt) bool { return shirt.Color == ocp.Black }
	// extended Predicate function for large shirt
	largeShirtsFilter := func(shirt ocp.Shirt) bool { return shirt.Size == ocp.Large }
	// extended Predicate function for black large shirt
	largeBlackShirtFilter := func(shirt ocp.Shirt) bool { return blackShirtsFilter(shirt) && largeShirtsFilter(shirt) }

	fmt.Println("Black shirts:")
	blackShirts2 := ocp.Filter(shirts, blackShirtsFilter)
	ocp.PrintShirtsOCP(blackShirts2)
	fmt.Println("Large shirts:")
	largeShirts2 := ocp.Filter(shirts, largeShirtsFilter)
	ocp.PrintShirtsOCP(largeShirts2)
	fmt.Println("Large black shirts:")
	largeBlackShirts := ocp.Filter(shirts, largeBlackShirtFilter)
	ocp.PrintShirtsOCP(largeBlackShirts)
}

func PrintLSP() {
	rc := &lsp.Rectangle{Width: 2, Height: 5}
	lsp.UseIt(rc)

	sq := lsp.NewSquare(5)
	lsp.UseIt(sq)
}

func PrintIsp() {
	mobile := &isp.MobileIsp{}
	mobile.SetAlarm(time.Now().Local().UTC().Sub(time.Time{}))
	mobile.DisableAlarm(time.Now().Local().UTC().Sub(time.Time{}))
	mobile.ConnectCall(9400976122)
	mobile.DisconnectCall(9400976122)
	mobile.PlayGame("PubG Mobile")
}
