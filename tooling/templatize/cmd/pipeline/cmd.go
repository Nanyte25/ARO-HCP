// Copyright 2025 Microsoft Corporation
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package pipeline

import (
	"github.com/spf13/cobra"

	"github.com/Azure/ARO-HCP/tooling/templatize/cmd/pipeline/validate"

	"github.com/Azure/ARO-HCP/tooling/templatize/cmd/pipeline/inspect"
	"github.com/Azure/ARO-HCP/tooling/templatize/cmd/pipeline/run"
)

func NewCommand() (*cobra.Command, error) {
	cmd := &cobra.Command{
		Use:              "pipeline",
		Short:            "Operate over pipelines in the local environment.",
		SilenceUsage:     true,
		SilenceErrors:    true,
		TraverseChildren: true,
		CompletionOptions: cobra.CompletionOptions{
			HiddenDefaultCmd: true,
		},
	}

	commands := []func() (*cobra.Command, error){
		run.NewCommand,
		inspect.NewCommand,
		validate.NewCommand,
	}
	for _, newCmd := range commands {
		c, err := newCmd()
		if err != nil {
			return nil, err
		}
		cmd.AddCommand(c)
	}

	return cmd, nil
}
