# Generate Usage
*WIP this document will be updated throughout the project*

## Description

The **Decentralized Tech Radar** has a Generate feature to create a HTML file from a CSV specfile.

## How to use generate

### General usage
When wanting to generate html from a CSV file, simply give the file path to the  CSV file in the generate command. This will generate an `index.html` file that can be opened to view the tech radar representation of the CSV file.

Command:
```bash
./Tech_radar-<OS> generate <csvFilePath>
```

### Example 
If we want to generate a tech radar from the `./data/technology_data.csv` file we can use use the `generate` command to achieve this:
```bash
./Tech_radar-<OS> generate ./data/technology_data.csv
```

*Note: on Windows the example ./Tech_Radar-\<OS\> should be written as .\Tech_Radar-\<OS\>. Note the slash "/" becoming different like so: "\\"*