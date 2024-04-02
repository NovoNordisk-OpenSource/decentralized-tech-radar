#!/usr/bin/env bash

# This script will fully run the command chain to fetch a csv file from a remote repository, cache it, merge it and then render it to a html file.

# The script will take the following arguments:
# 1. The URL of the remote repository(s)
# 2. The name of the branch to fetch the csv file from
# 3. The name of the specfile to use as a whitelist

URL=$1
Branch=$2
Specfile=$3

# Checks if the URL is empty
if [ -z "$URL" ]; then
    echo "Please provide the URL of the remote repository(s)"
    exit 1
fi

# Converts the URL to an array
strarr=($URL)

cd ../src

# Incrementer 
i=0

# Loop through the array and add the branch and specfile to the array
for url in ${strarr[@]}; do
    strarr[$i]=$url
    strarr[$i+1]=$Branch
    strarr[$i+2]=$Specfile
    i=$i+3
done

# Convert the array to a string
urlstr="${strarr[*]}"

# Fetch the csv file from the remote repository(s)
go run main.go --fetch "${urlstr}"

# ls the cache directory and get the csv files
# it gets converted to an array and then to a string (don't question it)
csvfiles=$(ls ./cache/*.csv)
csvfiles=($csvfiles)
csvfiles=${csvfiles[*]}

# Merge the csv files
go run main.go --merge "${csvfiles}"

# Render the csv file to a html file
go run main.go --file "./Merged_file.csv"
