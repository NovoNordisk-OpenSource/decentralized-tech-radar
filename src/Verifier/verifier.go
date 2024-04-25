package Verifier

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"regexp"
	"slices"
	"strings"
)

// This is for alternative names for blips e.g. CSharp, CS, C# all should be counted as C#)
// var alt_names_str = `Go;golang,go-lang
// 					C++;cpp
// 					C#;csharp,cs
// 					Python;python3,py
// 					`

// Map of alternative names for the same blip
var alt_names = make(map[string]string) //{"golang":"Go","go-lang:Go","cpp":"C++","csharp":"C#","cs":"C#","python3":"Python","py":"Python"}

// Checks that the given string matches the defined header of the specfile
func checkHeader(header string) bool {
	correctHeader := "name,ring,quadrant,isNew,moved,description"
	return header == correctHeader
}

// The pattern is created as a singleton so that we don't need
// to compile a new on every line that is checked
var regexPattern *regexp.Regexp = nil

// Creates the regex pattern string with the names on the rings
func createRegexPattern(ring1, ring2, ring3, ring4 string) {
	var err error
	regexPattern, err = regexp.Compile(fmt.Sprintf("^(([^,\n])([^,\n])*),([%s]%s|[%s]%s|[%s]%s|[%s]%s),([Dd]ata management|[Dd]atastore|[Ii]nfrastructure|[Ll]anguage),(false|true),-?[0123],(([^,\n])([^,\n])*)",
		strings.ToUpper(ring1[:1])+strings.ToLower(ring1[:1]), ring1[1:], strings.ToUpper(ring2[:1])+strings.ToLower(ring2[:1]), ring2[1:],
		strings.ToUpper(ring3[:1])+strings.ToLower(ring3[:1]), ring3[1:], strings.ToUpper(ring4[:1])+strings.ToLower(ring4[:1]), ring4[1:]))
	if err != nil {
		panic(err)
	}
}

// Checks that the given string matches the correct
// format for data in a specfile
func checkDataLine(data string) bool {
	if regexPattern == nil { // Check if we need to compile the pattern
		createRegexPattern("hold", "assess", "trial", "adopt")
	}

	match := regexPattern.MatchString(data)

	return match
}

func Verifier(filepaths ...string) error {
	for _, filepath := range filepaths {
		file, err := os.Open(filepath)
		if err != nil {
			panic(err)
		}

		defer file.Close()

		scanner := bufio.NewScanner(file)

		//TODO: Check if header matches (this will be another branch) for now it will just eat the first line
		scanner.Scan()
		if !checkHeader(scanner.Text()) {
			return errors.New("The header of " + filepath + " is not correct.\n\tCorrect header: name,ring,quadrant,isNew,moved,description\n\tHeader of " + filepath + ": " + scanner.Text())
		}

		// https://stackoverflow.com/questions/44073754/how-to-slice-string-till-particular-character-in-go
		for scanner.Scan() {
			line := scanner.Text()

			if line == "" {
				break
			}

			check := checkDataLine(line)

			if !check {
				return errors.New(filepath + " contains invalid data: " + scanner.Text())
			}
		}
	}
	return nil
}

func DuplicateRemoval(filepaths ...string) error {
	// Map functions as a set (name -> ring)
	var set = make(map[string][]string)
	for _, filepath := range filepaths {

		// Create temp file to overwrite primary file
		tempfile, err := os.Create("tempfile.csv")
		if err != nil {
			panic(err)
		}

		defer os.RemoveAll(tempfile.Name())
		defer tempfile.Close()

		file, err := os.Open(filepath)
		if err != nil {
			panic(err)
		}

		defer file.Close()
		scanner := bufio.NewScanner(file)

		// Skip header
		scanner.Scan()
		tempfile.WriteString(scanner.Text() + "\n")

		for scanner.Scan() {
			line := scanner.Text()
			// Faster than splitting
			// Panic handler
			name := ""
			index := strings.IndexByte(line, ',')
			if index != -1 {
				name = line[:index]
			}

			duplicateRemoval(name, line, tempfile, set)

		}
		file.Close()
		tempfile.Close()
		err = os.Rename("tempfile.csv", filepath)
		if err != nil {
			panic(err)
		}
	}
	return nil
}

func duplicateRemoval(name, line string, tempfile *os.File, set map[string][]string) error {
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
		ring := line[len(real_name)+1 : strings.IndexByte(line[len(real_name)+1:], ',')+len(real_name)+1]
		if !(slices.Contains(set[name], ring)) {
			set[name] = append(set[name], ring)
			tempfile.WriteString(line + "\n")
		}
	} else {
		set[name] = append(set[name], line[len(name)+1:strings.IndexByte(line[len(name)+1:], ',')+len(name)+1])
		tempfile.WriteString(line + "\n")
	}
	// Overwrite filepath with tempfile (has the removed changes)
	return nil
}
