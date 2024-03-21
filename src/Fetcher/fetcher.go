package Fetcher

import  (
	"os/exec"
)

func FetchFiles (url, specFile string) {
	cmd := exec.Command("python","./Fetcher/fetchfile.py", url, "main", specFile)
	_, err := cmd.CombinedOutput()
	if err != nil {
		panic(string(err.Error()) + " Failed at fetcher")
	}

}