# CSV Verifier
The verifier ensures that specfiles have the correct format. This means that the header is the one specified in [the documentation for the specfile](../user_docs/spec_file_format.md), and the each line of data also complies with format specified in the documentation for the specfile.

The verifier can be given multiple files and will return an error if a file does not comply with the expected format for a specfile.

## CSV header verification
The header of a csv file is compared directly against the correct header.

## CSV data verification
The verification for the data in the csv file is done line by line. A regex pattern is used to check that the line matches what can be expected from a line in a correct specfile.

## Functions
The verifier currently has the following functions:

* `getHeader(filepath string) ([]byte, error)`
  * **What it is:** A private function taking a string argument, which returns an Array of data type bytes.
  * **What it does:** This function takes the provided filepath to an CSV file to open it, reads the specification file's header, and writes it to a byte array.

* `checkHeader(header string) bool`
  * **What it is**: A private function for verifying the header, that takes a string being the header of the csv file, and return a boolean.
  * **What it does**: It takes the header and compares it to what the header should be, and returns a boolean based on whether or not the header is what is should be.

* `createRegexPattern(ring1, ring2, ring3, ring4 string)`
  * **What it is**: A private function for creating a regex pattern that takes the 4 string being the names of the rings in the tech radar.
  * **What it does**: It creates the regex pattern string using the ring names given, and compiles the pattern.

* `checkDataLine(data string) bool`
  * **What it is**: A private function for verifying the format of a csv data string.
  * **What it does**: It checks if a csv data string matches the specified regex pattern. If no pattern is compiled is calls `createRegexPattern()`.

* `Verifier(filepaths ...string) error`
  * **What it is**: A public function for verifying specfiles that takes 1 or more string being the path to csv files. Returns an error
  * **What it does**: Checks the header and each line of each file to see if they fit the format specified for specfiles. Returns an error if they don't.