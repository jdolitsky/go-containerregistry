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
	"github.com/google/go-containerregistry/pkg/v1/types"
)

// Artifact defines the interface for interacting with an OCI v1.1+ artifact.
type Artifact interface {
	// Blobs returns the unordered collection of blobs that comprise this artifact.
	Blobs() ([]Blob, error)

	// MediaType of this image's manifest.
	MediaType() (types.MediaType, error)

	// Size returns the size of the manifest.
	Size() (int64, error)

	// ConfigName returns the hash of the image's config file, also known as
	// the Image ID.
	ConfigName() (Hash, error)

	// ConfigFile returns this image's config file.
	ConfigFile() (*ConfigFile, error)

	// RawConfigFile returns the serialized bytes of ConfigFile().
	RawConfigFile() ([]byte, error)

	// Digest returns the sha256 of this image's manifest.
	Digest() (Hash, error)

	// Manifest returns this image's Manifest object.
	Manifest() (*ArtifactManifest, error)

	// RawManifest returns the serialized bytes of Manifest()
	RawManifest() ([]byte, error)

	// BlobByDigest returns a Blob for interacting with a particular blob of
	// the image, looking it up by "digest" (the compressed hash).
	BlobByDigest(Hash) (Blob, error)
}
