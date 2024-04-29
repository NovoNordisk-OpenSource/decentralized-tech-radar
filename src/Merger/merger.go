package Merger

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"slices"
	"strings"

	"github.com/NovoNordisk-OpenSource/decentralized-tech-radar/Verifier"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
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
	var filepaths_set = make(map[string]string)
	// Create a new logger

	os.Remove("Merge_log.txt")
	file, _ := os.OpenFile("Merge_log.txt", os.O_RDWR|os.O_CREATE, 0644)
	zapLogger(file)
	sugar := zapLogger(file)
	defer sugar.Sync()

	dup_count := 0
	for _, filepath := range filepaths {
		file, err := os.Open(filepath)
		if err != nil {
			panic(err)
		}

		defer file.Close()
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

			err := duplicateRemoval(name, line, filepath, buffer, set, filepaths_set)
			if err != nil {
				if dup_count == 0 {
					sugar.Info("Duplicates found in the following files:\n")
				}
				dup_count++
				sugar.Info("Duplicate overwrite "+fmt.Sprint(dup_count)+":",
					"\n\tLine "+fmt.Sprint(line_num)+":", err.Error(),
					"\n\tPicked: ", filepaths_set[name],
					"\n\tNot-Picked: ", filepath+"\n")
			}
		}
	}
	if dup_count == 0 {
		os.Remove("Merge_log.txt")
	}
	return nil
}

func duplicateRemoval(name, line, filepath string, buffer *bytes.Buffer, set map[string][]string, filepaths_set map[string]string) error {
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
			filepaths_set[name] = filepath
			buffer.Write([]byte(line + "\n"))
		} else {
			return fmt.Errorf(line)
		}
	} else {
		set[name] = append(set[name], line[len(name)+1:strings.IndexByte(line[len(name)+1:], ',')+len(name)+1])
		filepaths_set[name] = filepath
		buffer.Write([]byte(line + "\n"))
	}

	return nil
}
