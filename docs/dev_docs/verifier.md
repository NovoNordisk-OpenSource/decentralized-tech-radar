# CSV Verifier
The merger and fetcher use a CSV verifier to ensure that exact duplicates don't occur as well as to ensure the integrity of the CSV file (e.g. does it have commas as separators). The verifier is located in the Verifier package and can take any amount of filepaths. If the verifier finds a problem it will return an error. 
```go
func Verifier (filepaths ... string) error 
```
It checks the documents line by line and adds blips to a set of seen names and will remove any duplicate found in the future if it is in the same ring (this may be changed for LLM duplicate handling later). 

The verifier uses a alt_names map to ensure that alternative names are counted as the same thing. E.g. C#, CSharp, CS will all be mapped the same value ensuring they are counted as the same thing. Currently this value is hardcoded.

*Future expansion of will allow for header verification as well*
