package Merger

import (
	"bufio"
	"bytes"
	"os"
	"slices"
	"strings"
)

type Fcfs struct {}

func (f Fcfs) MergeFiles(buffer *bytes.Buffer, filepaths ...string) error {
	// Map functions as a set (name -> quadrant)
	var set = make(map[string][]string)
	for _, filepath := range filepaths {
		file, err := os.Open(filepath)
		if err != nil {
			panic(err)
		}

		defer file.Close()
		f.scanFile(file, buffer, set)
	}
	return nil
}

func (f Fcfs) scanFile(file *os.File, buffer *bytes.Buffer, set map[string][]string) {
	scanner := bufio.NewScanner(file)

	// Skip header
	scanner.Scan()

	for scanner.Scan() {
		line := scanner.Text()
		// Faster than splitting
		// Panic handler
		name := ""
		index := strings.IndexByte(line, ',')
		if index != -1 {
			name = line[:index]
		}

		f.duplicateRemoval(name, line, buffer, set)
	}
}

func (f Fcfs) duplicateRemoval(name, line string, buffer *bytes.Buffer, set map[string][]string) error {
	//TODO: Unmarshal the json file (or some other file based solution) to get the alternative names
	// Or just use a baked in str read line by line or combination
	//os.Stat("./Dictionary/alt_names.txt")

	real_name := name
	if alt_names[name] != "" {
		//TODO: Figure out how to handle numbers in names
		name = alt_names[strings.ToLower(name)]
	}

	if set[name] != nil {
		// Skips the name + first comma and does the same forward search for next comma
		quadrant := line[len(real_name)+1 : strings.IndexByte(line[len(real_name)+1:], ',')+len(real_name)+1]
		if !(slices.Contains(set[name], quadrant)) {
			set[name] = append(set[name], quadrant)
			buffer.Write([]byte(line + "\n"))
		}
	} else {
		set[name] = append(set[name], line[len(name)+1:strings.IndexByte(line[len(name)+1:], ',')+len(name)+1])
		buffer.Write([]byte(line + "\n"))
	}

	return nil
}
