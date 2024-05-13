# Fetcher

The fetcher is used to pull whitelisted files or folders from one or more git repositories. The progress of the fetcher can be tracked using the small spinner and progress bar displayed in the terminal, where each dot represents a repository.

The fetcher takes a string containing 3 values:  
1. A URL to a git-based repository
2. A branch name
3. A path to a whitelist file
These can be repeated in the string as shown below for multiple fetches
```go
func FetchFiles(url, branch, specFile string) error 
```

### Asynchronous
The fetcher runs using Go functions, meaning it runs asynchronously. This allows simultaneous fetching of files from multiple repositories, greatly decreasing the total running time on fetch calls with many repositories.

### Automatic Caching
The fetcher caches fetched CSV files in a folder named `cache`. The caching works for any folder depth so long as a CSV file is present

### Auto append Repository URLs
The fetcher will automatically append urls to the CSV files after verifying that they have no formatting errors. 