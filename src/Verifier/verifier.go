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

// Checks that the given string matches the defined header of the specfile
func checkHeader(header string) bool {
	correctHeader := "name,ring,quadrant,isNew,moved,description"
	return header == correctHeader
}

// Creates the regex pattern string with the names on the rings
func createRegexPattern(ring1, ring2, ring3, ring4 string) string {
	regexPattern := fmt.Sprintf("^(([^,\n])([^,\n])*),([%s]%s|[%s]%s|[%s]%s|[%s]%s),([Dd]ata management|[Dd]atastore|[Ii]nfrastructure|[Ll]anguage),(false|true),[0123],(([^,\n])([^,\n])*)",
								strings.ToUpper(ring1[:1]), strings.ToLower(ring1[1:]), strings.ToUpper(ring2[:1]), strings.ToLower(ring2[1:]), 
								strings.ToUpper(ring3[:1]), strings.ToLower(ring3[1:]), strings.ToUpper(ring4[:1]), strings.ToLower(ring4[1:]))
	return regexPattern
}

// Checks that the given string matches the correct 
// format for data in a specfile
func checkDataLine(data string) (bool, error) {
	pattern := createRegexPattern("hold", "assess", "trial", "adopt")
	
	match, err := regexp.MatchString(pattern, data)
	if err != nil {
		return false, err
	}

	return match, nil
}

func Verifier (filepaths ... string) error {
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
	if !checkHeader(scanner.Text()) {
		return errors.New("The header of " + filepath + " is not correct.\n\tCorrect header: name,ring,quadrant,isNew,moved,description\n\tHeader of "+ filepath + ": " + scanner.Text())
	}
	tempfile.WriteString(scanner.Text()+"\n")
	
	// https://stackoverflow.com/questions/44073754/how-to-slice-string-till-particular-character-in-go
	for scanner.Scan() {
		line := scanner.Text()

		if line == "" {
			break
		}
		
		check, err := checkDataLine(line)
		if err != nil {
			return err
		}
		if !check {
			return errors.New(filepath + " contains invalid data: " + scanner.Text())
		}

		// Faster than splitting
		// Panic handler 
		name := ""
		index := strings.IndexByte(line, ',')
		if index != -1 {
		name = line[:index]
		} 

		duplicateRemoval(filepath, name, line, tempfile)
		
	}
}
return nil
}

func duplicateRemoval(filepath, name, line string, tempfile *os.File) error {
	//TODO: Unmarshal the json file (or some other file based solution) to get the alternative names
	// Or just use a baked in str read line by line or combination
	//os.Stat("./Dictionary/alt_names.txt")

	alt_names := make(map[string]string)//{"golang":"Go","go-lang:Go","cpp":"C++","csharp":"C#","cs":"C#","python3":"Python","py":"Python"}
	alt_names["python3"] = "Python"

	// Map functions as a set (name -> ring)
	set := make(map[string][]string)


	real_name := name
	if alt_names[name] != "" {
		//TODO: Figure out how to handle numbers in names
		name = alt_names[strings.ToLower(name)]
	}

	if set[name] != nil {
		// Skips the name + first comma and does the same forward search for next comma
		ring := line[len(real_name)+1:strings.IndexByte(line[len(real_name)+1:], ',')+len(real_name)+1]
		if !(slices.Contains(set[name], ring)) {
			set[name] = append(set[name],ring)
			tempfile.WriteString(line+"\n")	
		}
	} else {
		set[name] = append(set[name],line[len(name)+1:strings.IndexByte(line[len(name)+1:], ',')+len(name)+1])
		tempfile.WriteString(line+"\n")
	}
	// Overwrite filepath with tempfile (has the removed changes)
	os.Rename("tempfile.csv", filepath)
	return nil
}
