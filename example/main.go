package main

import (
	"bytes"
	"context"
	"fmt"
	"log"
	"time"

	"github.com/micnncim/homebrew-go"
)

func main() {
	stdout := &bytes.Buffer{}
	stderr := &bytes.Buffer{}
	h := homebrew.NewHomebrew(
		homebrew.WithStdout(stdout),
		homebrew.WithStderr(stderr),
	)

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Minute)
	defer cancel()
	if err := h.Install(ctx, "jq", homebrew.InstallOptionVerbose, homebrew.InstallOptionForce); err != nil {
		log.Fatal(err)
	}
	fmt.Println(stdout.String())
	fmt.Println(stderr.String())
}
