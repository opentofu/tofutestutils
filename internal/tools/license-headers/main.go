// Copyright (c) The OpenTofu Authors
// SPDX-License-Identifier: MPL-2.0

package main

import (
	"errors"
	"flag"
	"log"
	"os"
	"os/exec"
)

func main() {
	checkOnly := false
	flag.BoolVar(&checkOnly, "check-only", checkOnly, "Only check if the license headers are correct.")
	flag.Parse()

	params := []string{"go", "run", "github.com/hashicorp/copywrite@v0.19.0", "headers"}
	if checkOnly {
		params = append(params, "--plan")
	}
	cmd := exec.Command(params[0], params[1:]...) //nolint:gosec // The parameters here are carefully assembled.
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		var exitError *exec.ExitError
		if errors.As(err, &exitError) {
			os.Exit(exitError.ExitCode())
		}

		log.Fatal(err)
	}
}
