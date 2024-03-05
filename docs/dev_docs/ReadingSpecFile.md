# Preface
*This file is WIP and will be updated continually*

# General
Technologies for the radar (or blips) are stored in the program using structs. These structs can then be passed between different parts of the program.

# Reading Spec files in JSON format
For reading spec files in json format, the `ReadJsonSpec` function is used. Since we know the structure of our data we can predefine our structs to fit that structure.  
The json reader can then unmarshal the json file. Taking the json string and turning it into our data structure.