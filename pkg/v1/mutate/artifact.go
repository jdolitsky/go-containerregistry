// Copyright 2019 Google LLC All Rights Reserved.
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

package mutate

import (
	"encoding/json"
	"errors"

	v1 "github.com/google/go-containerregistry/pkg/v1"
	"github.com/google/go-containerregistry/pkg/v1/partial"
	"github.com/google/go-containerregistry/pkg/v1/stream"
	"github.com/google/go-containerregistry/pkg/v1/types"
)

type artifact struct {
	base v1.Artifact
	adds []Addendum

	computed     bool
	manifest     *v1.ArtifactManifest
	annotations  map[string]string
	artifactType *types.ArtifactType
	mediaType    *types.MediaType
	digestMap    map[v1.Hash]v1.Blob
	subject      *v1.Descriptor
}

var _ v1.Artifact = (*artifact)(nil)

func (a *artifact) ArtifactType() (types.ArtifactType, error) {
	if a.mediaType != nil {
		return *a.artifactType, nil
	}
	return a.base.ArtifactType()
}

func (a *artifact) MediaType() (types.MediaType, error) {
	if a.mediaType != nil {
		return *a.mediaType, nil
	}
	return a.base.MediaType()
}

func (a *artifact) compute() error {
	return nil
}

// Blobs returns the unordered collection of blobs that comprise this artifact.
func (a *artifact) Blobs() ([]v1.Blob, error) {
	if err := a.compute(); errors.Is(err, stream.ErrNotComputed) {
		blobs, err := a.base.Blobs()
		if err != nil {
			return nil, err
		}
		for _, add := range a.adds {
			blobs = append(blobs, add.Blob)
		}
		return blobs, nil
	} else if err != nil {
		return nil, err
	}
	return a.base.Blobs()
}

// Digest returns the sha256 of this artifact's manifest.
func (i *artifact) Digest() (v1.Hash, error) {
	if err := i.compute(); err != nil {
		return v1.Hash{}, err
	}
	return partial.Digest(i)
}

// Size implements v1.Image.
func (i *artifact) Size() (int64, error) {
	if err := i.compute(); err != nil {
		return -1, err
	}
	return partial.Size(i)
}

// Manifest returns this image's Manifest object.
func (i *artifact) Manifest() (*v1.ArtifactManifest, error) {
	if err := i.compute(); err != nil {
		return nil, err
	}
	return i.manifest.DeepCopy(), nil
}

// RawManifest returns the serialized bytes of Manifest()
func (i *artifact) RawManifest() ([]byte, error) {
	if err := i.compute(); err != nil {
		return nil, err
	}
	return json.Marshal(i.manifest)
}

// BlobByDigest returns a Blob for interacting with a particular blob of
// the artifact, looking it up by "digest" (the compressed hash).
func (i *artifact) BlobByDigest(h v1.Hash) (v1.Blob, error) {
	if layer, ok := i.digestMap[h]; ok {
		return layer, nil
	}
	return i.base.BlobByDigest(h)
}
