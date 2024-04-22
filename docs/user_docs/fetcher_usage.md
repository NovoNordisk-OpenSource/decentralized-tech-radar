# Fetcher Usage

To fetch a file from a repository you can run the following command
For single repo fetch:
```bash
./Tech_Radar_<OS> fetch <Url> <Branch> <path/to/whitelist>
```
For multiple repos:
```bash
./Tech_Radar_<OS> fetch <Url> <Branch> <path/to/whitelist> <Url_1> <Branch_1> <path/to/whitelist1>
```
*The Url must be a valid public git repository otherwise it will not work*

The whitelist file should be a file containing filepaths to either files or folders you would like the fetcher to grab e.g.
```
/path
/root/test.csv
/usr/bin
```
## Fetcher Flags
`--branch=` Flag for setting the branch for all repos   
`--whitelist=` Flag for setting the whitelist file for all repos  
`--repo-file=` Flag for setting a file containing repositories to pull from *<u>& the branch + whitelist file if those flags have no been set<u>*   

Example:
```
./Tech_Radar_<OS> fetch <Url> --branch=<branch> --whitelist=<path/to/whitelist> --file=<path/to/repo-file>
```
