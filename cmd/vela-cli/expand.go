// Copyright (c) 2022 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package main

import (
	"github.com/go-vela/cli/command/pipeline"
	"github.com/go-vela/cli/internal"

	"github.com/urfave/cli/v2"
)

// expandCmds defines the commands for expanding resources.
var expandCmds = &cli.Command{
	Name:                   internal.ActionExpand,
	Category:               "Pipeline Management",
	Description:            "Use this command to expand a resource for Vela.",
	Usage:                  "Expand a resource for Vela via subcommands",
	UseShortOptionHandling: true,
	Subcommands: []*cli.Command{
		// add the sub command for expanding a pipeline
		//
		// https://pkg.go.dev/github.com/go-vela/cli/command/pipeline?tab=doc#CommandExpand
		pipeline.CommandExpand,
	},
}
