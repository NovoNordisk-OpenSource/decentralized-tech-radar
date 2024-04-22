# Preface

*WIP: This document will be updated throughout the project.*

# Merger

This feature was developed to aggregate multiple specification files[[1]](#1-user-docs--formatting-specification-files) into one file.

This can be used to create an overview of e.g. an entire department's tech stack, which can be used for reviews, analysis, or other types of breakdowns.

The merger ensures that no duplicates appear in the merged file. It checks the documents line by line and adds blips to a set of seen names and will remove any duplicate found in the future if it is in the same ring (this may be changed for LLM duplicate handling later). 

The verifier uses a alt_names map to ensure that alternative names are counted as the same thing. E.g. C#, CSharp, CS will all be mapped the same value ensuring they are counted as the same thing. Currently this value is hardcoded.

#### [1] [User Docs: Formatting specification files.](../user_docs/spec_file_format.md)

## Functions

The merger currently has three functions:

* `getHeader(filepath string) ([]byte, error)`
  * **What it is:** A private function taking a string argument, which returns an Array of data type bytes.
  * **What it does:** This function takes the provided filepath to an CSV file to open it, reads the specification file's header, and writes it to a byte array.

* `readCsvContent(filepath string) ([]byte, error)`
  * **What it is:** A private function taking a string argument, which returns an Array of data type bytes.
  * **What it does:** This function takes the provided filepath to an CSV file, opens the file, ignores the file's header but proceeds to read its content line by line, where each line is added to a bytes array.


* `MergeCSV(filepaths []string, header string) error`
  * **What it is:** A public function taking two arguments: An Array of data type string, and a string. It returns nothing.
  * **What it does:** If Merged_file.csv already exists, this is removed. Then it uses `readCsvContent()` to read each provided csv-file, writing each file's contents line by line into one new Merged_file.csv. These are read and written in the order of the file-paths provided.
  
* `MergeFromFolder(folderPath string) error`
  * **What it is**: A public function that takes one argument: A path to a folder, which in the default case, when adding the cache flag, is the cache folder itself.
  * **What it does**: If the cache folder exists, it reads all files from said folder, and appends them to cachePaths. It then checks whether or not cachePaths contain anything, and if so, merges the file with `MergeCSV(cachePaths)`.

* `DuplicateRemoval(filepaths ...string) error`
  * **What it is:**: A public function that takes multiple filepaths to specfiles of type string
  * **What it does**: Goes through the files given, removes duplicates, and overwrites the original file with a new file without the duplicates.