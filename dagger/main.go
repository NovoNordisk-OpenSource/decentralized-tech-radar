package main

import (
	"context"
	"os"
	"time"

	"dagger.io/dagger"
	"github.com/rs/zerolog/log"

	"dagger/logger"
)

func init() {
	log.Logger = logger.InitLogger()
}

func main() {
	log.Info().Msg("Running Dagger workflow")

	currentDir, err := os.Getwd()
	if err != nil {
		log.Error().Err(err).Msg("Failed to get current directory")
	} else {
		log.Info().Str("currentDir", currentDir).Msg("Variable:")
	}

	ctx := context.Background()

	if err := GenerateTechRadar(ctx); err != nil {
		log.Fatal().Err(err).Msg("Failed to generate tech radar")
	}

	log.Info().Msg("Dagger workflow completed")
}

func GenerateTechRadar(ctx context.Context) error {
	log.Info().Msg("Generating Tech Radar")

	// Initialize Dagger client
	client, err := dagger.Connect(ctx, dagger.WithLogOutput(log.Logger), dagger.WithLogOutput(os.Stdout))
	if err != nil {
		return err
	}
	defer client.Close()

	runner := client.Container().From("golang:1.22.5-bullseye").
		WithEnvVariable("CACHEBUSTER", time.Now().String())

	printOut(runner.WithExec([]string{"git", "version"}).Stdout(ctx))

	runner = runner.
		WithExec([]string{"git", "clone", "https://github.com/NovoNordisk-OpenSource/decentralized-tech-radar.git"}).
		WithExec([]string{"ls", "decentralized-tech-radar"})

	runner = runner.
		WithWorkdir("decentralized-tech-radar/src").
		WithDirectory("input", client.Host().Directory("input")).
		WithExec([]string{"ls", "-la"}).
		WithExec([]string{"cat", "input/whitelist.txt"}).
		WithExec([]string{"go", "mod", "tidy"}).
		WithExec([]string{"go", "build"}).
		WithExec([]string{"./decentralized-tech-radar"}).
		WithExec([]string{"./decentralized-tech-radar", "fetch", "--repo-file=input/repos.txt", "--whitelist=input/whitelist.txt"}).
		WithExec([]string{"ls", "cache"}).
		WithExec([]string{"bash", "-c", "for f in cache/*.csv; do echo $f; cat $f; done"}).
		WithExec([]string{"./decentralized-tech-radar", "merge", "--cache"}).
		WithExec([]string{"ls", "cache"}).
		WithExec([]string{"ls"}).
		WithExec([]string{"cat", "Merge_log.txt"}).
		WithExec([]string{"./decentralized-tech-radar", "generate", "Merged_file.csv"}).
		WithExec([]string{"ls"}).
		WithExec([]string{"mkdir", "output"}).
		WithExec([]string{"cp", "index.html", "output"}).
		WithExec([]string{"cp", "-r", "HTML", "output"}).
		WithExec([]string{"rm", "-rf", "output/HTML/html_test.go"}).
		WithExec([]string{"rm", "-rf", "output/HTML/makeHtml.go"}).
		WithExec([]string{"ls", "HTML"}).
		WithExec([]string{"ls", "output/HTML"})

	_, err = runner.
		Directory("output").
		Export(ctx, "output")
	if err != nil {
		return err
	}

	return nil
}

func printOut(o string, err error) {
	if err == nil {
		log.Info().Str("output", o).Msg("Output:")
	}
}
