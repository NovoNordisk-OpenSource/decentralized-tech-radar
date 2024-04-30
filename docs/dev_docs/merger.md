# Preface

*WIP: This document will be updated throughout the project.*

# Merger

This feature was developed to aggregate multiple specification files[[1]](#1-user-docs--formatting-specification-files) into one file.

This can be used to create an overview of e.g. an entire department's tech stack, which can be used for reviews, analysis, or other types of breakdowns.

The merger ensures that no duplicates appear in the merged file. It checks the documents line by line and adds blips to a set of seen names and will remove any duplicate found in the future if it is in the same ring (this may be changed for LLM duplicate handling later). It also adds what blips come from which repos, if repos are specified in the file that are being merged. 

The verifier uses a alt_names (currently not used) map to ensure that alternative names are counted as the same thing. E.g. C#, CSharp, CS will all be mapped the same value ensuring they are counted as the same thing. Currently this value is hardcoded.

#### [1] [User Docs: Formatting specification files.](../user_docs/spec_file_format.md)

## Functions

The merger currently has three functions:

* `getHeader(filepath string) ([]byte, error)`
  * **What it is**: A private function taking a string argument, which returns an Array of data type bytes.
  * **What it does**: This function takes the provided filepath to an CSV file to open it, reads the specification file's header, and writes it to a byte array.

* `MergeCSV(filepaths []string, header string) error`
  * **What it is**: A public function taking two arguments: An Array of data type string, and a string. It returns nothing.
  * **What it does**: If Merged_file.csv already exists, this is removed. Then it uses `ReadCsvData()` to read each provided csv-file into a buffer. The buffer is then written to a Merged_file.csv file. These are read in the order of the file-paths provided.
  
* `MergeFromFolder(folderPath string) error`
  * **What it is**: A public function that takes one argument: A path to a folder, which in the default case, when adding the cache flag, is the cache folder itself.
  * **What it does**: If the cache folder exists, it reads all files from said folder, and appends them to cachePaths. It then checks whether or not cachePaths contain anything, and if so, merges the file with `MergeCSV(cachePaths)`.

* `ReadCsvData(buffer *bytes.Buffer, filepaths ...string) error`
  * **What it is**: A public function that takes a pointer to a byte buffer that contains the merged specfiles' data, without the duplicates, and multiple filepaths to specfiles of type string.
  * **What it does**: Goes through the files given, line by line, and calls `scanFile()` with each line.

* `scanFile(file *os.File, set map[string][]string, blipRepos *map[string]map[string]byte)`
 * **What it is**: A private function that takes a file, a map for all the seen blips' names, and a pseudo buffer (a map of a set, but Go doesn't have a set, so it's a map).
 * **What it does**: Scans and calls duplicateRemoval() for each line in the file.

* `bufferWriter(buffer *bytes.Buffer, blips map[string]map[string]byte) error`
  * **What it is**: A private function that takes a buffer and a pseudo buffer.
  * **What it does**: Correctly writes the pseudo buffer to the real buffer, so it can be written to a file.

* `duplicateRemoval(name, line string, set map[string][]string, blipRepos map[string]map[string]byte) error`
  * **What it is**: A private function that takes a blip name from the given line, the line, the map for all the seen blips' names, and a pseudo buffer (a map of a set, but Go doesn't have a set, so it's a map).
  * **What it does**: Checks if the line has references to the repos that the blip came from, and runs the duplication removal functions accordingly.

* `duplicateRemovalWithoutUrl(name, line string, set map[string][]string, blipRepos map[string]map[string]byte) error`
  * **What it is**: A private function that takes a blip name from the given line, the line, the map for all the seen blips' names, and a pseudo buffer (a map of a set, but Go doesn't have a set, so it's a map).
  * **What it does**: Writes the line to the pseudo buffer if the blip's name is not already in that quadrant. If the blip's name is already in the quadrant it is not added to the pseudo buffer.

* `duplicateRemovalWithUrl(name, line string, set map[string][]string, blipRepos map[string]map[string]byte) error`
  * **What it is**: A private function that takes a blip name from the given line, the line, the map for all the seen blips' names, and a pseudo buffer (a map of a set, but Go doesn't have a set, so it's a map).
  * **What it does**: Writes the line to the pseudo buffer if the blip's name is not already in that quadrant, and adds the repos that blip came from to the pseudo buffer. If the blip's name is already in the quadrant it is not added to the pseudo buffer.