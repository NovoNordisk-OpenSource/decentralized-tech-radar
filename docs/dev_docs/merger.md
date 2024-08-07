# Preface

*WIP: This document will be updated throughout the project.*

# Merger

This feature was developed to aggregate multiple specification files[[1]](#1-user-docs--formatting-specification-files) into one file.

This can be used to create an overview of e.g. an entire department's tech stack, which can be used for reviews, analysis, or other types of breakdowns.

The merger ensures that no duplicates appear in the merged file. It does this by calling the `MergeFiles()` function from the `MergeStrat` interface. A merge strategy is passed to the merger and defines how merging is handled. This allows for easily changing the merging strategy. See [here](./merger_strategies.md) for more information on merging strategies.

#### [1] [User Docs: Formatting specification files.](../user_docs/spec_file_format.md)

## Functions

The merger currently has three functions:

* `getHeader(filepath string) ([]byte, error)`
  * **What it is**: A private function taking a string argument, which returns an Array of data type bytes.
  * **What it does**: This function takes the provided filepath to an CSV file to open it, reads the specification file's header, and writes it to a byte array.

* `MergeCSV(filepaths []string, header string, strat MergeStrat) error`
  * **What it is:** A public function taking three arguments: An Array of data type string, a string defining the header, and an implementation of the MergeStrat interface.
  * **What it does:** If Merged_file.csv already exists, this is removed. Then it uses `ReadCsvData()` to read each provided csv-file into a buffer. The buffer is then written to a Merged_file.csv file. These are read in the order of the file-paths provided.
  
* `MergeFromFolder(folderPath string, strat MergeStrat) error`
  * **What it is**: A public function that takes two arguments: A path to a folder, which in the default case, when adding the cache flag, is the cache folder itself; and an implementation of the MergeStrat interface.
  * **What it does**: If the cache folder exists, it reads all files from said folder, and appends them to cachePaths. It then checks whether or not cachePaths contain anything, and if so, merges the file with `MergeCSV(cachePaths)`.

* `bufferWriter(buffer *bytes.Buffer, blips map[string]map[string]byte) error`
  * **What it is**: A private function that takes a buffer and a pseudo buffer.
  * **What it does**: Correctly writes the pseudo buffer to the real buffer, so it can be written to a file.

* `zapLogger(f *os.File) *zap.SugaredLogger`
  * **What it is**: A private function that takes a file and returns a logger
  * **What it does**: Creates a zap logger to the file that the is given as param.

The merger also contains the interface for merging strategies:

```Go
type MergeStrat {
  MergeFiles(buffer *bytes.Buffer, filepaths ...string) error
}
```