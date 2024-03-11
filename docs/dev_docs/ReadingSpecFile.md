# Preface
*This file is WIP and will be updated continually*

# General
Technologies for the radar (or blips) are stored in the program using structs. These structs can then be passed between different parts of the program.

# Reading Spec files from CSV format
For reading spec files in csv format, the `ReadCsvSpec` function is used. Since we know the structure of our data we can predefine our structs to fit that structure.  
By using the [gocarina/gocsv](https://github.com/gocarina/gocsv) module we can  unmarshal the csv file. Taking the csv file and turning it into our data structure.