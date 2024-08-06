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

	// runner := client.Container().From("alpine:3.20.1").
	// runner := client.Container().From("python:3.12.2-bookworm")
	runner := client.Container().From("golang:1.22.5-bullseye").WithEnvVariable("CACHEBUSTER", time.Now().String())
	// WithExec([]string{"apk", "add", "git"}).

	printOut(runner.WithExec([]string{"git", "version"}).Stdout(ctx))

	runner = runner.
		WithExec([]string{"git", "clone", "https://github.com/NovoNordisk-OpenSource/decentralized-tech-radar.git"}).
		WithExec([]string{"ls", "decentralized-tech-radar"})

	runner = runner.
		WithWorkdir("decentralized-tech-radar/src").
		WithExec([]string{"go", "mod tidy"}).
		WithExec([]string{"go", "build"}).
		WithExec([]string{"./decentralized-tech-radar"})

	_, err = runner.
		Export(ctx, "decentralized-tech-radar/src/decentralized-tech-radar")
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
