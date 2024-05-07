package main

import (
	"context"
	"os"

	"dagger.io/dagger"
)

func main() {
	ctx := context.Background()

	// initialize Dagger client
	client, err := dagger.Connect(ctx, dagger.WithLogOutput(os.Stderr))
	if err != nil {
		panic(err)
	}
	defer client.Close()

	goCache := client.CacheVolume("golang")

	// use a node:16-slim container
	// mount the source code directory on the host
	// at /src in the container
	source := client.Container().
		From("golang:1.21").
		WithDirectory("/d_src", client.Host().Directory(".", dagger.HostDirectoryOpts{
			Exclude: []string{},
		})).WithMountedCache("/d_src/ci/cache", goCache)

	// set the working directory in the container

	runner := source.WithWorkdir("/d_src/")

	// Run the fetcher using the binary
	runner = runner.WithWorkdir("/d_src/").WithExec([]string{"./Tech_Radar-linux", "fetch","--repo-file=repos.txt", "--whitelist=whitelist.txt"})

	// Run the merger using the binary
	runner = runner.WithWorkdir("/d_src/").WithExec([]string{"./Tech_Radar-linux", "merge", "--cache"})

	// Run the HTML generator using the binary
	exporter := runner.WithWorkdir("/d_src/").WithExec([]string{"./Tech_Radar-linux", "generate", "Merged_file.csv"})

	exp_dir := exporter.Directory("/d_src/d_src")

	// Output the HTML file to the host
	_, err = exp_dir.Export(ctx, ".")
	if err != nil {
		panic(err)
	}

}
