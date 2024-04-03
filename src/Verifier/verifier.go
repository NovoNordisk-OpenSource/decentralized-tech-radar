package Verifier

import (
	"bufio"
	"errors"
	"os"
	"slices"
	"strings"
)

// This is for alternative names for blips e.g. CSharp, CS, C# all should be counted as C#)
var alt_names = make(map[string]string)

//TODO: Add ... param argument for filepath instead to allow for multiple CSV files to be checked in one go
func Verifier (filepaths ... string) error {
	// Map functions as a set (name -> ring)
	set := make(map[string][]string)

	//TODO: This is where a for loop for multi csv should go
	for _, filepath := range filepaths {
	
	file, err := os.Open(filepath)
	if err != nil {
		panic(err)
	}

	// Create temp file to overwrite primary file
	tempfile, err := os.Create("tempfile.csv")
	if err != nil {
		panic(err)
	}

	defer file.Close()
	defer tempfile.Close()
	defer os.RemoveAll(tempfile.Name())

	scanner := bufio.NewScanner(file)
	
	//TODO: Check if header matches (this will be another branch) for now it will just eat the first line
	scanner.Scan()
	tempfile.WriteString(scanner.Text()+"\n")
	
	// https://stackoverflow.com/questions/44073754/how-to-slice-string-till-particular-character-in-go
	for scanner.Scan() {
		line := scanner.Text()

		if line == "" {
			break
		}

		// Faster than splitting
		name := line[:strings.IndexByte(line, ',')]
		if name == "" {
			return errors.New("No comma was found format of csv file is wrong: triggered by line -> "+line)
		}

		if alt_names[name] != "" {
			name = alt_names[name]
		}
		
		if set[name] != nil {
			// Skips the name + first comma and does the same forward search for next comma
			ring := line[len(name):strings.IndexByte(line, ',')]
			if !slices.Contains(set[name], ring) {
				set[name] = append(set[name],line[len(name):strings.IndexByte(line, ',')])
				tempfile.WriteString(line+"\n")	
			}
		} else {
			set[name] = append(set[name],line[len(name):strings.IndexByte(line, ',')])
			tempfile.WriteString(line+"\n")
		}
	}
	// Overwrite filepath with tempfile (has the removed changes)
	
	os.Rename("tempfile.csv", filepath)
	}
return nil
}