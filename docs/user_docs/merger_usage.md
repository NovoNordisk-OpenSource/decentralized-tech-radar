# Merger usage

*Subject to change.*

## Description

The **Decentralized Tech Radar** has a Merger feature to combine the contents of multiple CSV spec-files.

The merged files will be put into one file called `Merged_file.csv` generated in the root directory.

## How to use merge

### General usage
When wanting to merge multiple files into one, the `merge` command can be used. This takes the paths to the files that you would like to merge. At least two file paths are required (if the `--cache` flag is not used). This will call the merger to merge the provided files into one file.
Command:
```bash
./Tech_radar-<OS> merge <path/to/csv-file1> [path/to/csv-file2] [path/to/csv-file3] ...
```
The file-paths are separated by a space.

When wanting to merge `./data/file0.csv` and `./data/file1.csv`, run the following command:
```bash
./Tech_radar-<OS> merge ./data/file0.csv ./data/file1.csv
```

*Note: on Windows the example ./Tech_Radar-\<OS\> should be written as .\Tech_Radar-\<OS\>. Note the slash "/" becoming different like so: "\\"*

### Merge flags
`--cache || -c` Flag for merging all files in the cache the `./cache` directory. This is used instead of the individual file paths.

### Example
```bash
./Tech_radar-<OS> merge --cache
```
