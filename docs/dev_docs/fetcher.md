# Fetcher

The fetcher is used to pull whitelisted files/folders from one or more git repositories. It takes a string containing 3 values:  
1. A URL to a git based repository
2. A branch name
3. A path to a whitelist file  
These can be repeated in the string as shown below for multiple fetches
```go
func FetchFiles(url, branch, specFile string) error 
```

The fetcher caches fetched CSV files in a folder named `cache`. The caching works for any folder depth so long as a CSV file is present