# Welcome to Tech_Radar quickstart guide
In this guide you will find the essential commands for the program. The program is split into 3 primary commands: Fetcher, Merger, and Generator.

## Legend for this guide
When running commands, you will see...
**<...>** : Arrows indicate required arguments.
**[...]** : Square-brackets indicate optional arguments.

## The Fetcher
The fetcher will fetch files from one or more repositories (depending on amount of arguments given). It will then run through those files and cache them if they are CSV files.  

To run the fetcher you can use the following command depending on your OS:

Mac:
```bash
./Tech_Radar-<OS> fetch <git-url> <branch> <path/to/whitelistfile> [git-url2] [branch2] [path/to/whitelistfile2] ...
```

Windows:
```bash
.\Tech_Radar-<OS> fetch <git-url> <branch> <path/to/whitelistfile> [git-url2] [branch2] [path/to/whitelistfile2] ...
```


`git-url` being any valid public git repository  
`branch` being any valid branch name for the given repository  
`path/to/whitelistfile` being a path to a locally stored whitelist file *__see fetcher_usage.md__*  

Example:  
![Gif of fetching using the CLI](https://github.com/NovoNordisk-OpenSource/decentralized-tech-radar/blob/main/docs/images/quickstart_gifs/Fetch.gif)

## The Merger
The merger takes in one or more CSV files and merges them into one big file called `Merged_file.csv`. The merger will use a verifier to check if any of the CSV files are formatted wrong or contain the wrong information.
To run the merger you can use the following command on your OS:

Mac:
```bash
./Tech_Radar-<OS> merge <path/to/csvfile> <path/to/csvfile2> [path/to/csvfile3] ...
```

Windows:
```bash
.\Tech_Radar-<OS> merge <path\to\csvfile> <path\to\csvfile2> [path\to\csvfile3] ...
```


`path/to/csvfile` being a path to a locally stored csvfile with correct format *__see spec_file_format.md__*  


Example:  
![Gif of merging using the CLI](https://github.com/NovoNordisk-OpenSource/decentralized-tech-radar/blob/main/docs/images/quickstart_gifs/Merging.gif)

The merger also includes a cache flag this will run the merger using the cache folder generated from the fetcher (If the cache is located next to the program):

Mac:
```bash
./Tech_Radar-<OS> merge --cache
```

Windows:
```bash
.\Tech_Radar-<OS> merge --cache
```

## The Generator
The generator takes a merged CSV file and generates a tech radar HTML file from it. This can then be opened in a web browser.

To run the generator you can use the following command in your OS:

Mac:
```bash
./Tech_Radar-<OS> generate <path/to/csvfile>
```

Windows:
```bash
.\Tech_Radar-<OS> generate <path\to\csvfile>
```


Example:  
![Gif of generating using the CLI](https://github.com/NovoNordisk-OpenSource/decentralized-tech-radar/blob/main/docs/images/quickstart_gifs/Generate.gif)