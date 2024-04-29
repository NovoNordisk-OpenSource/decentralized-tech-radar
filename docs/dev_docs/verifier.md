# CSV Verifier
The merger and fetcher use a CSV verifier to ensure that the CSV complies with what is expected from a specification file. This is done by checking that the header of the file matches the defined specification file header. 

The verifier is located in the Verifier package and can take any amount of filepaths. If the verifier finds a problem it will return an error. 
```go
func Verifier (filepaths ... string) error 
```
