package Verifier

import (
	"bufio"
	"errors"
	"os"
	"slices"
	"strings"
)

// This is for alternative names for blips e.g. CSharp, CS, C# all should be counted as C#)
// var alt_names_str = `Go;golang,go-lang
// 					C++;cpp
// 					C#;csharp,cs
// 					Python;python3,py
// 					`


func Verifier (filepaths ... string) error {

	//TODO: Unmarshal the json file (or some other file based solution) to get the alternative names
	// Or just use a baked in str read line by line or combination
	//os.Stat("./Dictionary/alt_names.txt")

	alt_names := make(map[string]string)//{"golang":"Go","go-lang:Go","cpp":"C++","csharp":"C#","cs":"C#","python3":"Python","py":"Python"}
	alt_names["python3"] = "Python"



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
		// Panic handler 
		name := ""
		index := strings.IndexByte(line, ',')
		if index != -1 {
		name = line[:index]
		} 
		
		//TODO: Change these to be variable names later if we extend the program with custom ring names
		if strings.Contains(name,"Hold") || strings.Contains(name,"Adopt") || strings.Contains(name,"Assess") || strings.Contains(name, "Trial") || name == "" {
			os.Rename("tempfile.csv", filepath)
			return errors.New("No comma was found format of csv file is wrong: triggered by line -> "+line)
		}

		real_name := name
		if alt_names[name] != "" {
			//TODO: Figure out how to handle numbers in names
			name = alt_names[strings.ToLower(name)]
		}

		if set[name] != nil {
			// Skips the name + first comma and does the same forward search for next comma
			ring := line[len(real_name)+1:strings.IndexByte(line[len(real_name)+1:], ',')+len(real_name)+1]
			if !(slices.Contains(set[name], ring)) {
				//print(set[name][0],ring)
				set[name] = append(set[name],ring)
				tempfile.WriteString(line+"\n")	
			}
		} else {
			set[name] = append(set[name],line[len(name)+1:strings.IndexByte(line[len(name)+1:], ',')+len(name)+1])
			tempfile.WriteString(line+"\n")
		}
	}
	// Overwrite filepath with tempfile (has the removed changes)
	os.Rename("tempfile.csv", filepath)
	}
return nil
}