package Merger

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"slices"
	"strings"

	"github.com/NovoNordisk-OpenSource/decentralized-tech-radar/Verifier"
)

// Map of alternative names for the same blip
var alt_names = make(map[string]string) //{"golang":"Go","go-lang:Go","cpp":"C++","csharp":"C#","cs":"C#","python3":"Python","py":"Python"}

func getHeader(filepath string) ([]byte, error) {
	file, err := os.Open(filepath)
	if err != nil {
		return []byte{}, err // Propagate error
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	headerBytes := scanner.Bytes()
	headerBytes = append(headerBytes, []byte("\n")...)

	return headerBytes, nil
}

func MergeFromFolder(folderPath string) error {
	_, err := os.Stat(folderPath)
	if os.IsNotExist(err) {
		return errors.New("Folder does not exist or could not be found. \nError: " + err.Error())
	} else if err != nil {
		return err
	}

	cachedRepos, err := os.ReadDir(folderPath)
	if err != nil {
		return err
	}

	var cachePaths []string
	for _, repo := range cachedRepos {
		if filepath.Ext(repo.Name()) == ".csv" {
			cachePaths = append(cachePaths, filepath.Join(folderPath, repo.Name()))
		}
	}

	if len(cachePaths) == 0 {
		fmt.Println("There are currently no files in the cache.")
	}

	MergeCSV(cachePaths)

	return nil
}

func MergeCSV(filepaths []string) error {
	os.Remove("Merged_file.csv") // Remove file in case it already exists

	// Run data verifier on files
	err := Verifier.Verifier(filepaths...)
	if err != nil {
		panic(err)
	}

	var buf bytes.Buffer
	// Add header to buffer
	header, err := getHeader(filepaths[0])
	if err != nil {
		return err // Propagate error
	}
	buf.Write(header)

	// Read csv data which removes duplicates
	// This only adds non-duplicates to the buffer
	err = ReadCsvData(&buf, filepaths...)
	if err != nil {
		panic(err)
	}

	// Write combined files to one file
	err = os.WriteFile("Merged_file.csv", buf.Bytes(), 0644)
	if err != nil {
		return err
	}

	return nil
}

func ReadCsvData(buffer *bytes.Buffer, filepaths ...string) error {
	// Map functions as a set (name -> quadrant)
	var set = make(map[string][]string)

	// This is hacky, but don't worry about it :)
	// Map [lineWithFirstUrl] -> map[uniqueUrls[1:]]nil
	var blipRepos = make(map[string]map[string]byte)
	for _, filepath := range filepaths {
		file, err := os.Open(filepath)
		if err != nil {
			panic(err)
		}

		defer file.Close()
		scanFile(file, buffer, set, &blipRepos)
	}
	
	bufferWriter(buffer, blipRepos)
	return nil
}

func scanFile(file *os.File, buffer *bytes.Buffer, set map[string][]string, blipRepos *map[string]map[string]byte) {
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
		
		err := duplicateRemoval(name, line, set, *blipRepos)
		if err != nil {
			panic(err)
		}
	}
}



// Regex:
// <br>Repos:<br> (?<atag><a href=(?<url>(https:\/\/www.|http:\/\/www.|https:\/\/|http:\/\/)([-a-zA-Z0-9]{2,})(.[a-zA-Z0-9]{2,})(.[a-zA-Z0-9]{2,})?(\/[-a-zA-Z0-9_\/.]{2,}))>(?<repoName>[-a-zA-Z\d_.]+)<\/a>)/gm

// Map [BlipName] -> map[uniqueUrls]nil
func bufferWriter(buffer *bytes.Buffer, blips map[string]map[string]byte) error {
	var sb strings.Builder
	for line, intMap := range blips {
		sb.WriteString(line)
		for atag, _ := range intMap {
			sb.WriteString("<br>")
			sb.WriteString(atag)
		}
		sb.WriteRune('\n')
	}

	buffer.WriteString(sb.String())

	return nil
}

var regexPattern *regexp.Regexp = nil
var pattern = "<br>Repos:(<br><a href=((https://www.|http://www.|https://|http://)([-a-zA-Z0-9]{2,})(.[a-zA-Z0-9]{2,})(.[a-zA-Z0-9]{2,})?(/[-a-zA-Z0-9_/.]{2,}))>([-a-zA-Z\\d_.]+)</a>)+"


func duplicateRemoval(name, line string, set map[string][]string, blipRepos map[string]map[string]byte) error {
	//TODO: Unmarshal the json file (or some other file based solution) to get the alternative names
	// Or just use a baked in str read line by line or combination
	//os.Stat("./Dictionary/alt_names.txt")

	
	var err error
	if regexPattern == nil {
		regexPattern, err = regexp.Compile(pattern)
		if err != nil {
			panic(err)
		}
	}

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
			splitLine := strings.Split(line, "<br>Repos:<br>")
			blipRepos[splitLine[0] + "<br>Repos:"] = make(map[string]byte)
			for _, repo := range strings.Split(splitLine[1], "<br>") {
				blipRepos[splitLine[0] + "<br>Repos:"][repo] = 0
			}

		// This else is for items we have already seen
		} else {
			splitLine := strings.Split(line, "<br>Repos:<br>")
			atags := strings.Split(splitLine[1], "<br>")
			
			for _, atag := range atags {
				blipRepos[splitLine[0] + "<br>Repos:"][atag] = 0
			}
		}
	} else {
		set[name] = append(set[name], line[len(name)+1:strings.IndexByte(line[len(name)+1:], ',')+len(name)+1])
		// buffer.Write([]byte(line + "\n"))'
		splitLine := strings.Split(line, "<br>Repos:<br>")
		blipRepos[splitLine[0] + "<br>Repos:"] = make(map[string]byte)
		for _, repo := range strings.Split(splitLine[1], "<br>") {
			blipRepos[splitLine[0] + "<br>Repos:"][repo] = 0
		}
	}

	return nil
}
