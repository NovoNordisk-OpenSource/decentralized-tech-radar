# Fetcher Usage
The fetcher fetches CSV files from the repos and puts them in a `./cache` folder next to the program. When fetching a file the repo it came from is also added to the description of each blip in that file.


## Usage
To fetch a file from a repository you can run the following command
For single repo fetch:
```bash
./Tech_Radar-<OS> fetch <Url> <Branch> <path/to/whitelist>
```
For multiple repos:
```bash
./Tech_Radar-<OS> fetch <Url> <Branch> <path/to/whitelist> <Url_1> <Branch_1> <path/to/whitelist1>
```

*Note: on Windows the example ./Tech_Radar-\<OS\> should be written as .\Tech_Radar-\<OS\>. Note the slash "/" becoming different like so: "\\"*

*The Url must be a valid public git repository otherwise it will not work*

The whitelist file should be a file containing filepaths to either files or folders you would like the fetcher to grab from e.g:
```
/path
/root/test.csv
/usr/bin
```
The fetcher will get every `.csv` file in the file tree that is specified in the whitelist file. If you want it to only get a specific file. That file can be specified in the whitelist like on line 2 in the example.

## Fetcher Flags
`--branch=` Flag for setting the branch for all repos   
`--whitelist=` Flag for setting the whitelist file for all repos  
`--repo-file=` Flag for setting a file containing repositories to pull from *<u>& the branch + whitelist file if those flags have no been set</u>* 

## Examples
Example fetching one repo:
```bash
./Tech_Radar-<OS> fetch https://github.com/NovoNordisk-OpenSource/example-repo-1 main ./whitelist.txt
```

Example fetching one repo using flags:
```bash
./Tech_Radar-<OS> fetch https://github.com/NovoNordisk-OpenSource/example-repo-1 --branch=main --whitelist=./whitelist.txt
```

Example fetching multiple repos:
```bash
./Tech_Radar-<OS> fetch https://github.com/NovoNordisk-OpenSource/example-repo-1 main ./whitelist.txt https://github.com/NovoNordisk-OpenSource/example-repo-2 main ./whitelist.txt
```

Example fetching multiple repos using flags:
```bash
./Tech_Radar-<OS> fetch https://github.com/NovoNordisk-OpenSource/example-repo-1 https://github.com/NovoNordisk-OpenSource/example-repo-2 --branch=main --whitelist=./whitelist.txt
```