// Copyright (c) 2020 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package build

import (
	"github.com/go-vela/cli/internal/output"

	"github.com/go-vela/sdk-go/vela"
)

// View inspects a build based off the provided configuration.
func (c *Config) View(client *vela.Client) error {
	// send API call to capture a build
	build, _, err := client.Build.Get(c.Org, c.Repo, c.Number)
	if err != nil {
		return err
	}

	// handle the output based off the provided configuration
	switch c.Output {
	case "json":
		// output the build in JSON format
		err := output.JSON(build)
		if err != nil {
			return err
		}
	case "yaml":
		// output the build in YAML format
		err := output.YAML(build)
		if err != nil {
			return err
		}
	default:
		// output the build in default format
		err := output.Default(build)
		if err != nil {
			return err
		}
	}

	return nil
}