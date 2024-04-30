package Merger

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"slices"
	"strings"

	"go.uber.org/zap"
)

type Fcfs struct{}

func (f Fcfs) MergeFiles(buffer *bytes.Buffer, filepaths ...string) error {
	// Map functions as a set (name -> quadrant)
	var set = make(map[string][]string)

	// Clear the log file
	os.Remove("Merge_log.txt")
	file, _ := os.OpenFile("Merge_log.txt", os.O_RDWR|os.O_CREATE, 0644)

	// Create a new sugared logger (zap)
	// The logger requires 3 things to log correctly.
	// 1. The filepath_set which is a map of names to filepaths (name -> filepath) to keep track of which file the name was picked from
	// 2. The current file you are scanning (file)
	// 3. The Merge strat should return a custom error that will be the line that was a duplicate (line)
	sugar := zapLogger(file)
	defer sugar.Sync() // Flushes buffer, if any events are buffered

	for _, filepath := range filepaths {
		file, err := os.Open(filepath)
		if err != nil {
			panic(err)
		}

		defer file.Close()
		f.scanFile(file, buffer, set, sugar) // Calls duplicateRemoval on each line with the logger
	}
	return nil
}

func (f Fcfs) scanFile(file *os.File, buffer *bytes.Buffer, set map[string][]string, sugar *zap.SugaredLogger) {
	scanner := bufio.NewScanner(file)

	// Skip header
	scanner.Scan()
	line_num := 0
	for scanner.Scan() {
		line := scanner.Text()
		line_num++

		// Faster than splitting
		// Panic handler
		name := ""
		index := strings.IndexByte(line, ',')
		if index != -1 {
			name = line[:index]
		}

		err := f.duplicateRemoval(name, line, file.Name() , buffer, set)
		if err != nil { // Duplicate found
			if dup_count == 0 {
				// Log the header of the merge log file
				sugar.Info("Duplicates found in the following files:\n")
			}
			dup_count++
			// Log the duplicate
			sugar.Info("Duplicate overwrite "+fmt.Sprint(dup_count)+":",
				"\n\tLine "+fmt.Sprint(line_num)+":", err.Error(),
				"\n\tPicked: ", filepaths_set[name],
				"\n\tNot-Picked: ", file.Name()+"\n")
		}
	}
}

func (f Fcfs) duplicateRemoval(name, line, filepath string, buffer *bytes.Buffer, set map[string][]string) error {

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
			filepaths_set[name] = filepath // Add the filepath to the set (name -> filepath)
			buffer.Write([]byte(line + "\n"))
		} else {
			return fmt.Errorf(line)
		}
	} else {
		set[name] = append(set[name], line[len(name)+1:strings.IndexByte(line[len(name)+1:], ',')+len(name)+1])
		filepaths_set[name] = filepath // Add the filepath to the set (name -> filepath)
		buffer.Write([]byte(line + "\n"))
	}

	return nil
}
