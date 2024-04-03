package solid

import (
	"fmt"
	"os"
	"strings"
)

const newLine = "\n"

type Note struct {
	Lines    map[int]string
	lastLine int
}

func (note *Note) AddLine(line string) int {
	note.Lines[note.lastLine] = line
	note.lastLine++
	return len(note.Lines)
}

func (note *Note) DeleteLine(number int) int {
	delete(note.Lines, number)
	return len(note.Lines)
}

func (note *Note) UpdateLine(number int, correctedLine string) int {
	if _, exists := note.Lines[number]; exists {
		note.Lines[number] = correctedLine
	}
	return len(note.Lines)
}

func (note *Note) GetByLineNo(number int) string {
	if value, exists := note.Lines[number]; exists {
		return value
	}
	return ""
}

func (note *Note) PrintLines() string {
	linesWithNumbers := make([]string, 0)
	for key, value := range note.Lines {
		linesWithNumbers = append(linesWithNumbers, fmt.Sprintf("%d %s", key, value))
	}
	return strings.Join(linesWithNumbers, newLine)
}

// This is a violation of SRP , Persistance can be moved to another package or object
func (note *Note) PersistNote(fileName string) error {
	err := os.WriteFile(fileName, []byte(note.PrintLines()), 0644)
	return err
}
