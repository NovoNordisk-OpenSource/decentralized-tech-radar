# Fetcher Usage

To fetch a file from a repository you can run the following command
For single repo fetch:
```bash
./Tech_Radar_<OS> --fetch "<Url> <Branch> <Whitelist_Filepath>"
```
For multiple repos:
```bash
./Tech_Radar_<OS> --fetch "<Url> <Branch> <Whitelist_Filepath> <Url_1> <Branch_1> <Whitelist_Filepath_1>"
```
*The Url must be a valid public git repository otherwise it will not work*

The whitelist file should be a file containing filepaths to either files or folders you would like the fetcher to grab e.g.
```
/path
/root/test.csv
/usr/bin
```

