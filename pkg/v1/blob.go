// Copyright 2018 Google LLC All Rights Reserved.
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

package v1

import (
	"io"

	"github.com/google/go-containerregistry/pkg/v1/types"
)

// Blob is an interface for accessing the properties of a particular blob of a v1.Artifact
type Blob interface {
	// Digest returns the Hash of the compressed blob.
	Digest() (Hash, error)

	// Compressed returns an io.ReadCloser for the compressed blob contents.
	Compressed() (io.ReadCloser, error)

	// Uncompressed returns an io.ReadCloser for the uncompressed blob contents.
	Uncompressed() (io.ReadCloser, error)

	// Size returns the compressed size of the Blob.
	Size() (int64, error)

	// MediaType returns the media type of the Blob.
	MediaType() (types.MediaType, error)
}
