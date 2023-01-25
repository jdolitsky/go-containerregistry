// Copyright 2022 Google LLC All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"os"

	"github.com/google/go-containerregistry/pkg/crane"
	"github.com/google/go-containerregistry/pkg/v1/types"
	"github.com/spf13/cobra"
)

// NewCmdAttach creates a new cobra.Command for the attach subcommand.
func NewCmdAttach(options *[]crane.Option) *cobra.Command {
	var artifactType, mediaType string

	attachmentsCmd := &cobra.Command{
		Use:   "attach",
		Short: "Add an attachment to an image",
		// TODO: Long
		Args: cobra.ExactArgs(2),
		RunE: func(_ *cobra.Command, args []string) error {
			refstr := args[0]
			fn := args[1]
			b, err := os.ReadFile(fn)
			if err != nil {
				return err
			}
			return crane.Attach(refstr, b, artifactType, mediaType, *options...)
		},
	}
	attachmentsCmd.Flags().StringVar(&artifactType, "artifact-type", string(types.OCIConfigJSON), "Value for config.mediaType field")
	attachmentsCmd.Flags().StringVar(&mediaType, "media-type", string(types.OCILayer), "Value for layers[0].mediaType field")
	return attachmentsCmd
}
