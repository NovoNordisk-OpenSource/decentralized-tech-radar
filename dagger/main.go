package main

import (
	"context"
	"os"

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

	runner := client.Container().From("alpine:3.20.1").
		WithExec([]string{"apk", "add", "git"}).
		WithExec([]string{"git", "version"}).
		WithExec([]string{"sh", "-c", "ls -l > test.txt"}).
		WithExec([]string{"cat", "test.txt"})

	runner = runner.
		WithExec([]string{"git", "clone", "https://github.com/NovoNordisk-OpenSource/decentralized-tech-radar.git"})

	_, err = runner.
		WithExec([]string{"ls", "decentralized-tech-radar"}).
		Export(ctx, "test.txt")
	if err != nil {
		return err
	}

	return nil
}
