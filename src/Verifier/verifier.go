package Verifier

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"regexp"
	"strings"
)

// This is for alternative names for blips e.g. CSharp, CS, C# all should be counted as C#)
// var alt_names_str = `Go;golang,go-lang
// 					C++;cpp
// 					C#;csharp,cs
// 					Python;python3,py
// 					`

// Checks that the given string matches the defined header of the specfile
func checkHeader(header string) bool {
	correctHeader := "name,ring,quadrant,isNew,moved,description"
	return header == correctHeader
}

// The pattern is created as a singleton so that we don't need
// to compile a new on every line that is checked
var regexPattern *regexp.Regexp = nil

// Creates the regex pattern string with the names on the rings
func createRegexPattern(ring1, ring2, ring3, ring4, qudrant1, quadrant2, quadrant3, quadrant4 string) {
	var err error
	// This is the regex pattern that matches the correct format for a data line in the csv specfile
	// Example of a correct line:
	// 		Python,hold,language,false,0,Lorem ipsum dolor sit amet consectetur adipiscing elit.
	// Example of a incorrect line:
	// 		Python,wait,infrastructure,False,5,Lorem ipsum dolor sit amet consectetur adipiscing elit.
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

