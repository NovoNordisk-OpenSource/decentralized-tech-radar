# Preface

*WIP: This document will be updated throughout the project.*

# Merger strategies

Merger strategies are used by the merger for merging multiple specfiles into one. The merger strategy defines how merging is done. It defines how writing to the buffer, that is written to the final file, is done and what is written.

All merge strategies need to implement the `MergeStrat` interface from [Merger](./merger.md). This requires that the function `MergeFiles(buffer *bytes.Buffer, filepaths ...string) error` is implemented by the struct for the merging strategy. 

## Creating a Merge strategy
To create a new merge strategy can be done by:
- Creating a new file in the Merger package
- Create a struct for the merger strategy
- Implement the interface on the struct

The struct for the new merge strategy can now be used when calling the merger.

# Existing merger strategies
## Fcfs: First Come, First Served
Adds the first version of a technology in quadrant to the buffer and skips other occurrences.  
This strategy is defined as `type Fcfs struct {}`. This struct then implements the following functions:

### Functions
* `(f Fcfs) MergeFiles(buffer *bytes.Buffer, filepaths ...string) error`
  * **What it is:**: A public function that takes a pointer to a byte buffer that contains the merged specfiles' data, without the duplicates, and multiple filepaths to specfiles of type string.
  * **What it does**: Goes through the files given, line by line, and calls `scanFile()` with each line.

* `(f Fcfs) scanFile(file *os.File, buffer *bytes.Buffer, set map[string][]string)`
  * **What it is:**: A private function that takes a file, a byte buffer and a map for all the seen blips' names.
  * **What it does:** Scans and calls duplicateRemoval() for each line in the file.

* `(f Fcfs) duplicateRemoval(name, line string, buffer *bytes.Buffer, set map[string][]string) error`
  * **What it is**: A private function that takes a blip name from the given line, the line, the buffer for the merged specfiles, and the map for all the seen blips' names.
  * **What it does**: Writes the line to the buffer if the blip's name is not already in that quadrant. If the blip's name is already in the quadrant it is not added to the buffer.