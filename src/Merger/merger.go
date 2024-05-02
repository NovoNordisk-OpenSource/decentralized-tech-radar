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
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type MergeStrat interface {
	// A function that updates the buffer with the correct information
	// depending on that merge strategy
	MergeFiles(buffer *bytes.Buffer, filepaths ...string) error
}

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

func MergeFromFolder(folderPath string, start MergeStrat) error {
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

	MergeCSV(cachePaths, start)

	return nil
}

func MergeCSV(filepaths []string, strat MergeStrat) error {
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
	err = strat.MergeFiles(&buf, filepaths...)
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

func zapLogger(f *os.File) *zap.SugaredLogger {
	// https://stackoverflow.com/questions/50933936/zap-logger-print-both-to-console-and-to-log-file
	pe := zap.NewProductionEncoderConfig()

	//fileEncoder := zapcore.NewJSONEncoder(pe)

	cfg := zap.NewProductionConfig()

	cfg.EncoderConfig.LevelKey = zapcore.OmitKey
	cfg.EncoderConfig.TimeKey = zapcore.OmitKey

	pe.EncodeTime = zapcore.ISO8601TimeEncoder
	consoleEncoder := zapcore.NewConsoleEncoder(cfg.EncoderConfig)

	core := zapcore.NewTee(
		zapcore.NewCore(consoleEncoder, zapcore.AddSync(f), zap.InfoLevel),
		zapcore.NewCore(consoleEncoder, zapcore.AddSync(os.Stdout), zap.InfoLevel),
	)

	l := zap.New(core)

	return l.Sugar()

}

func ReadCsvData(buffer *bytes.Buffer, filepaths ...string) error {
	// Map functions as a set (name -> quadrant)
	var set = make(map[string][]string)

	// This is hacky, but don't worry about it :)
	// Map [lineWithoutURLs] -> map[uniqueUrls]byte
	// This map tracks all the repos a specific blip comes from
	// so that they can be added to the merged file
	var blipRepos = make(map[string]map[string]byte)
	for _, filepath := range filepaths {
		file, err := os.Open(filepath)
		if err != nil {
			panic(err)
		}

		defer file.Close()
		scanFile(file, set, &blipRepos)
	}
	
	bufferWriter(buffer, blipRepos)
	return nil
}

func scanFile(file *os.File, set map[string][]string, blipRepos *map[string]map[string]byte) {
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

// Create the string and write it to the buffer
func bufferWriter(buffer *bytes.Buffer, blips map[string]map[string]byte) error {
	var sb strings.Builder
	for line, intMap := range blips {
		sb.WriteString(line) // Write the main line 
		for atag := range intMap { // Write the repos (doesn't run if no URLs in the internal map)
			sb.WriteString("<br>")
			sb.WriteString(atag)
		}
		sb.WriteRune('\n')
	}

	buffer.WriteString(sb.String())

	return nil
}

var regexPattern *regexp.Regexp = nil
// The regex pattern to check for repo URLs
// Example of string with URL:
//		Python,hold,language,false,0,Lorem ipsum dolor sit amet consectetur adipiscing elit.<br>Repos:<br> <a href=https://github.com/Agile-Arch-Angels/decentralized-tech-radar_dev>decentralized-tech-radar_dev</a>
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
	
	// Check if the line has repo urls
	if regexPattern.MatchString(line) {
		duplicateRemovalWithUrl(name, line, set, blipRepos)
	} else {
		duplicateRemovalWithoutUrl(name, line, set, blipRepos)
	}

	return nil
}

func duplicateRemovalWithoutUrl(name, line string, set map[string][]string, blipRepos map[string]map[string]byte) error {
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
		}
	} else { // Blip with current name is not in blip set
		set[name] = append(set[name], strings.Split(line, ",")[2]) // Add quadrant to blip set
		blipRepos[line] = make(map[string]byte) // Add line to the pseudo buffer (dumb map of map thing)
	}

	return nil
}

func duplicateRemovalWithUrl(name, line string, set map[string][]string, blipRepos map[string]map[string]byte) error {
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
		}
	} else { // Blip with current name is not in blip set
		// See comments above
		set[name] = append(set[name], strings.Split(line, ",")[2])
		blipRepos[blipInfo] = make(map[string]byte)
		for _, repo := range strings.Split(splitLine[1], "<br>") {
			blipRepos[blipInfo][repo] = 0
		}
	}

	return nil
}