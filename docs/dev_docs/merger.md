# Preface

*WIP: This document will be updated throughout the project.*

# Merger

This feature was developed to aggregate multiple specification files[[1]](#1-user-docs--formatting-specification-files) into one file.

This can be used to create an overview of e.g. an entire department's tech stack, which can be used for reviews, analysis, or other types of breakdowns.

#### [1] [User Docs: Formatting specification files.](../user_docs/spec_file_format.md)

## Functions

The merger currently has three functions:

* `getHeader(filepath string) []byte`
  * **What it is:** A private function taking a string argument, which returns an Array of data type bytes.
  * **What it does:** This function takes the provided filepath to an CSV file to open it, reads the specification file's header, and writes it to a byte array.

* `readCsvContent(filepath string) []byte`
  * **What it is:** A private function taking a string argument, which returns an Array of data type bytes.
  * **What it does:** This function takes the provided filepath to an CSV file, opens the file, ignores the file's header but proceeds to read its content line by line, where each line is added to a bytes array.


* `MergeCSV(filepaths []string, header string)`
  * **What it is:** A public function taking two arguments: An Array of data type string, and a string. It returns nothing.
  * **What it does:** If Merged_file.csv already exists, this is removed. Then it uses `readCsvContent()` to read each provided csv-file, writing each file's contents line by line into one new Merged_file.csv. These are read and written in the order of the file-paths provided.