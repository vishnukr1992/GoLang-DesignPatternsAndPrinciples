package solid

import (
	"fmt"
	"os"
	"strings"
)

type NoteSRP struct {
	Lines    map[int]string
	lastLine int
}

func (noteSrp *NoteSRP) AddLine(line string) int {
	noteSrp.Lines[noteSrp.lastLine] = line
	noteSrp.lastLine++
	return len(noteSrp.Lines)
}

func (noteSrp *NoteSRP) DeleteLine(number int) int {
	delete(noteSrp.Lines, number)
	return len(noteSrp.Lines)
}

func (noteSrp *NoteSRP) UpdateLine(number int, correctedLine string) int {
	if _, exists := noteSrp.Lines[number]; exists {
		noteSrp.Lines[number] = correctedLine
	}
	return len(noteSrp.Lines)
}

func (noteSrp *NoteSRP) GetByLineNo(number int) string {
	if value, exists := noteSrp.Lines[number]; exists {
		return value
	}
	return ""
}

func (noteSrp *NoteSRP) PrintLines(lineSeparator string) string {
	linesWithNumbers := make([]string, 0)
	for key, value := range noteSrp.Lines {
		linesWithNumbers = append(linesWithNumbers, fmt.Sprintf("%d %s", key, value))
	}
	return strings.Join(linesWithNumbers, newLine)
}

// We moved the persistance to different object, in real scenario this could be in a new file
type NoteFile struct {
	LineSeparator string
}

func (noteFile *NoteFile) PersistNote(note *NoteSRP, fileName string) error {
	err := os.WriteFile(fileName, []byte(note.PrintLines(noteFile.LineSeparator)), 0644)
	return err
}
