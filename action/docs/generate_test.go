// Copyright (c) 2022 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package docs

import (
	"testing"

	"github.com/go-vela/cli/internal"
	"github.com/urfave/cli/v2"
)

func TestDocs_Config_Generate(t *testing.T) {
	// setup tests
	fakeApp := cli.NewApp()

	tests := []struct {
		failure bool
		config  *Config
	}{
		{
			failure: false,
			config: &Config{
				Action:   internal.ActionGenerate,
				Markdown: true,
			},
		},
		{
			failure: false,
			config: &Config{
				Action: internal.ActionGenerate,
				Man:    true,
			},
		},
		{
			failure: true,
			config: &Config{
				Action: internal.ActionGenerate,
			},
		},
	}

	// run tests
	for _, test := range tests {
		err := test.config.Generate(fakeApp)

		if test.failure {
			if err == nil {
				t.Errorf("Generate should have returned err")
			}

			continue
		}

		if err != nil {
			t.Errorf("Generate returned err: %v", err)
		}
	}
}
