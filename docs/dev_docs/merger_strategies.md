# Preface

*WIP: This document will be updated throughout the project.*

# Merger strategies

Merger strategies are used by the merger for merging multiple specfiles into one. The merger strategy defines how merging is done. It defines how writing to the buffer, that is written to the final file, is done and what is written.

All merge strategies need to implement the `MergeStrat` interface from [Merger](./merger.md). This requires that the function `MergeFiles(buffer *bytes.Buffer, filepaths ...string) error` is implemented by the struct for the merging strategy. 

## Creating a Merge strategy
Creating a new merge strategy can be done by:
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
  * **What it does**: Goes through the files given and calls `scanFile()` with each file. Creates the logger for the merge report and and calls `bufferWriter()` with the buffer and the pseudo buffer.

* `(f Fcfs) scanFile(file *os.File, set map[string][]string, sugar *zap.SugaredLogger)`
  * **What it is**: A private function that takes a file, a map for all the seen blips' names, and a pseudo buffer (a map of a set, but Go doesn't have a set, so it's a map).
  * **What it does**: Scans and calls duplicateRemoval() for each line in the file, and adds information about the scan to the logger.

* `(f Fcfs) duplicateRemoval(name, line string, set map[string][]string, blipRepos map[string]map[string]byte) error`
  * **What it is**: A private function that takes a blip name from the given line, the line, the map for all the seen blips' names, and a pseudo buffer (a map of a set, but Go doesn't have a set, so it's a map).
  * **What it does**: Checks if the line has references to the repos that the blip came from, and runs the duplication removal functions accordingly.

* `(f Fcfs) duplicateRemovalWithoutUrl(name, line string, set map[string][]string, blipRepos map[string]map[string]byte) error`
  * **What it is**: A private function that takes a blip name from the given line, the line, the map for all the seen blips' names, and a pseudo buffer (a map of a set, but Go doesn't have a set, so it's a map).
  * **What it does**: Writes the line to the pseudo buffer if the blip's name is not already in that quadrant. If the blip's name is already in the quadrant it is not added to the pseudo buffer.

* `(f Fcfs) duplicateRemovalWithUrl(name, line string, set map[string][]string, blipRepos map[string]map[string]byte) error`
  * **What it is**: A private function that takes a blip name from the given line, the line, the map for all the seen blips' names, and a pseudo buffer (a map of a set, but Go doesn't have a set, so it's a map).
  * **What it does**: Writes the line to the pseudo buffer if the blip's name is not already in that quadrant, and adds the repos that blip came from to the pseudo buffer. If the blip's name is already in the quadrant it is not added to the pseudo buffer.