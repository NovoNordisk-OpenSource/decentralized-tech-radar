# Merger usage

*Subject to change.*

## Description

The **Decentralized Tech Radar** has a Merger feature to combine the contents of multiple CSV spec-files.

**Note:** The Merger does not currently handle duplicated content.

## How to use

When wanting to merge `file0.csv` and `file1.csv`, run the following command:

```
go run main.go -merge "<path-to-file0.csv> <path-to-file1>"
```

The file-paths are separated by a space.