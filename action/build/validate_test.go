// Copyright (c) 2020 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package build

import (
	"testing"
)

func TestBuild_Config_Validate(t *testing.T) {
	// setup tests
	tests := []struct {
		failure bool
		config  *Config
	}{
		{
			failure: false,
			config: &Config{
				Action:  "get",
				Org:     "MyOrg",
				Repo:    "HelloWorld",
				Page:    1,
				PerPage: 10,
				Output:  "json",
			},
		},
		{
			failure: false,
			config: &Config{
				Action: "restart",
				Org:    "MyOrg",
				Repo:   "HelloWorld",
				Number: 1,
				Output: "json",
			},
		},
		{
			failure: false,
			config: &Config{
				Action: "view",
				Org:    "MyOrg",
				Repo:   "HelloWorld",
				Number: 1,
				Output: "json",
			},
		},
		{
			failure: true,
			config: &Config{
				Action: "view",
				Org:    "",
				Repo:   "HelloWorld",
				Number: 1,
				Output: "json",
			},
		},
		{
			failure: true,
			config: &Config{
				Action: "view",
				Org:    "MyOrg",
				Repo:   "",
				Number: 1,
				Output: "json",
			},
		},
		{
			failure: true,
			config: &Config{
				Action: "view",
				Org:    "MyOrg",
				Repo:   "HelloWorld",
				Number: 0,
				Output: "json",
			},
		},
	}

	// run tests
	for _, test := range tests {
		err := test.config.Validate()

		if test.failure {
			if err == nil {
				t.Errorf("Validate should have returned err")
			}

			continue
		}

		if err != nil {
			t.Errorf("Validate returned err: %v", err)
		}
	}
}