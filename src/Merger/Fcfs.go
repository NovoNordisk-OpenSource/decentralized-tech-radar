package Merger

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"regexp"
	"slices"
	"strings"

	"go.uber.org/zap"
)

// Map of alternative names for the same blip
// var alt_names = make(map[string]string) //{"golang":"Go","go-lang:Go","cpp":"C++","csharp":"C#","cs":"C#","python3":"Python","py":"Python"}

var dup_count = 0 // Counter for duplicates

var filepaths_set = make(map[string]string) // Map of names to filepaths (name -> filepath) to keep track of which file the name was picked from

type Fcfs struct{}
var blipRepos = make(map[string]map[string]byte)

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
		f.scanFile(file, set, sugar) // Calls duplicateRemoval on each line with the logger
	}
	bufferWriter(buffer, blipRepos)
	return nil
}

func (f Fcfs) scanFile(file *os.File, set map[string][]string, sugar *zap.SugaredLogger) {
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

		err := f.duplicateRemoval(name, line, file.Name() , set, blipRepos)
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

var regexPattern *regexp.Regexp = nil

// The regex pattern to check for repo URLs
// Example of string with URL:
//
//	Python,hold,language,false,0,Lorem ipsum dolor sit amet consectetur adipiscing elit.<br>Repos:<br> <a href=https://github.com/Agile-Arch-Angels/decentralized-tech-radar_dev>decentralized-tech-radar_dev</a>
var pattern = "<br>Repos:(<br><a href=((https://www.|http://www.|https://|http://)([-a-zA-Z0-9]{2,})(.[a-zA-Z0-9]{2,})(.[a-zA-Z0-9]{2,})?(/[-a-zA-Z0-9_/.]{2,}))>([-a-zA-Z\\d_.]+)</a>)+"

func (f Fcfs) duplicateRemoval(name, line, filename string, set map[string][]string, blipRepos map[string]map[string]byte) error {
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

	// Check if the line has repo urls
	if regexPattern.MatchString(line) {
		err := f.duplicateRemovalWithUrl(name, line, filename , set, blipRepos)
		if err != nil {
			return err
		}
	} else {
		err := f.duplicateRemovalWithoutUrl(name, line, filename , set, blipRepos)
		if err != nil {
			return err
		}
	}

	return nil
}

func (f Fcfs) duplicateRemovalWithoutUrl(name, line, filename string, set map[string][]string, blipRepos map[string]map[string]byte) error {
	// This code does noting right now but can be added back if alt_names get properly populated
	// real_name := name
	// if alt_names[name] != "" {
	// 	//TODO: Figure out how to handle numbers in names
	// 	name = alt_names[strings.ToLower(name)]
	// }
	if set[name] != nil {
		// Skips the name + first comma and does the same forward search for next comma
		quadrant := strings.Split(line, ",")[2]
		if !(slices.Contains(set[name], quadrant)) { // If blip name not in quadrant
			set[name] = append(set[name], quadrant) // Add quadrant to blip set
			blipRepos[line] = make(map[string]byte) // Add line to the pseudo buffer (dumb map of map thing)
		} else {
			return fmt.Errorf(line)
		}
	} else { // Blip with current name is not in blip set
		set[name] = append(set[name], strings.Split(line, ",")[2]) // Add quadrant to blip set
		filepaths_set[name] = filename                               // Add the name to the filepaths set
		blipRepos[line] = make(map[string]byte)                    // Add line to the pseudo buffer (dumb map of map thing)
	}

	return nil
}

func (f Fcfs) duplicateRemovalWithUrl(name, line, filename string, set map[string][]string, blipRepos map[string]map[string]byte) error {
	// This code does noting right now but can be added back if alt_names get properly populated
	// real_name := name
	// if alt_names[name] != "" {
	// 	//TODO: Figure out how to handle numbers in names
	// 	name = alt_names[strings.ToLower(name)]
	// }
	splitLine := strings.Split(line, "<br>Repos:<br>")
	blipInfo := splitLine[0] + "<br>Repos:"

	if set[name] != nil {
		// Skips the name + first comma and does the same forward search for next comma
		quadrant := strings.Split(line, ",")[2]
		if !(slices.Contains(set[name], quadrant)) { // If blip name not in quadrant
			set[name] = append(set[name], quadrant) // Add quadrant to blip set
			filepaths_set[name] = filename       // Add the name to the filepaths set

			// Split up the blip info and the repos
			// Example of line with repos:
			// 		Python,hold,language,false,0,Lorem ipsum dolor sit amet consectetur adipiscing elit.<br>Repos:<br> <a href=https://github.com/Agile-Arch-Angels/decentralized-tech-radar_dev>decentralized-tech-radar_dev</a>
			blipRepos[blipInfo] = make(map[string]byte) // Add the string that is the blip info without every repo

			// Get all repos by splitting on the <br> between them
			// Example of multiple repos:
			// 		[...]scing elit.<br>Repos:<br> <a href=https://github.com/Agile-Arch-Angels/decentralized-tech-radar_dev>decentralized-tech-radar_dev</a><br> <a href=https://github.com/August-Brandt/DTR-specfile-generator>DTR-specfile-generator</a>
			for _, repo := range strings.Split(splitLine[1], "<br>") {
				blipRepos[blipInfo][repo] = 0 // Add repos to pseudo buffer (0 is just pointless value. It's because Go got no sets)
			}

			// This else is for items we have already seen
		} else {
			atags := strings.Split(splitLine[1], "<br>") // Get all the repos

			for _, atag := range atags {
				blipRepos[blipInfo][atag] = 0 // the repos to the pseudo buffer
			}
			return fmt.Errorf(line)
		}
	} else { // Blip with current name is not in blip set
		// See comments above
		set[name] = append(set[name], strings.Split(line, ",")[2])
		filepaths_set[name] = filename
		blipRepos[blipInfo] = make(map[string]byte)
		for _, repo := range strings.Split(splitLine[1], "<br>") {
			blipRepos[blipInfo][repo] = 0
		}
	}

	return nil
}
