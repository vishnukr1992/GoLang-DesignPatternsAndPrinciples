package main

import (
	"fmt"
	solid "golang/designprinciplesandpatterns/solid/srp"
)

func main() {
	fmt.Println("SRP implementation test")
	note := solid.NoteSRP{Lines: make(map[int]string), LastLine: 1}
	note.AddLine("I am lazy")
	note.AddLine("Oh no devin will replace me :-(")
	fmt.Println(note.PrintLines("\n"))

	noteFile := solid.NoteFile{LineSeparator: "\r\n"}
	noteFile.PersistNote(&note, "test.txt")
}
