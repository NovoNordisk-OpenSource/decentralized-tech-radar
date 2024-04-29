#!/usr/bin/env bash

# This script will fully run the command chain to fetch a csv file from a remote repository, cache it, merge it and then render it to a html file.

# The script will take the following arguments:
# 1. The URL of the remote repository(s)
# 2. The name of the branch to fetch the csv file from (defaults to main)
# 3. The name of the specfile to use as a whitelist (defaults to root whitelisted specfile)

URL=$1
Branch=$2
Specfile=$3

# Checks if the URL is empty
if [ -z "$URL" ]; then
    echo "Please provide the URL of the remote repository(s)"
    exit 1
fi

if [ -z "$Branch" ]; then
    Branch="main"
fi


# Converts the URL to an array
strarr=($URL)

cd ../../src

if [ -z "$Specfile" ]; then
    touch specfile.txt
    cat "/" > specfile.txt
    Specfile="specfile.txt"
fi

# Incrementer 
i=0

# Loop through the array and add the branch and specfile to the array
for url in ${strarr[@]}; do
    strarr[$i]=$url
    strarr[$i+1]=$Branch
    strarr[$i+2]=$Specfile
    i=$i+3
done

# Fetch the csv file from the remote repository(s)
go run main.go fetch "${strarr[@]}"

# ls the cache directory and get the csv files
# it gets converted to an array and then to a string (don't question it)
# csvfiles=$(ls ./cache/*.csv)
# csvfiles=($csvfiles)


# Merge the csv files
go run main.go merge --cache

# Render the csv file to a html file
go run main.go generate "./Merged_file.csv"

# Check OS and open the html file in the browser
if [[ "$OSTYPE" == "darwin"* ]]; then 
    open index.html
else 
    xdg-open index.html
fi