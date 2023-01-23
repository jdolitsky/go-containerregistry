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

	computed        bool
	configFile      *v1.ConfigFile
	manifest        *v1.ArtifactManifest
	annotations     map[string]string
	mediaType       *types.MediaType
	configMediaType *types.MediaType
	digestMap       map[v1.Hash]v1.Layer
	subject         *v1.Descriptor
}

var _ v1.Artifact = (*artifact)(nil)

func (i *artifact) MediaType() (types.MediaType, error) {
	if i.mediaType != nil {
		return *i.mediaType, nil
	}
	return i.base.MediaType()
}

func (i *artifact) compute() error {
	return nil
}

// Blobs returns the unordered collection of blobs that comprise this artifact.
func (i *artifact) Blobs() ([]v1.Blob, error) {
	if err := i.compute(); errors.Is(err, stream.ErrNotComputed) {
		blobs, err := i.base.Blobs()
		if err != nil {
			return nil, err
		}
		for _, add := range i.adds {
			blobs = append(blobs, add.Layer)
		}
		return blobs, nil
	} else if err != nil {
		return nil, err
	}
	return i.base.Blobs()
}

// ConfigName returns the hash of the artifact's config file.
func (i *artifact) ConfigName() (v1.Hash, error) {
	if err := i.compute(); err != nil {
		return v1.Hash{}, err
	}
	return partial.ConfigName(i)
}

// ConfigFile returns this artifact's config file.
func (i *artifact) ConfigFile() (*v1.ConfigFile, error) {
	if err := i.compute(); err != nil {
		return nil, err
	}
	return i.configFile.DeepCopy(), nil
}

// ConfigFile returns this image's config file.
func (i *artifact) RawConfigFile() ([]byte, error) {
	if err := i.compute(); err != nil {
		return nil, err
	}
	//return i.configFile.DeepCopy(), nil
	return nil, nil
}

// Digest returns the sha256 of this image's manifest.
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
	if cn, err := i.ConfigName(); err != nil {
		return nil, err
	} else if h == cn {
		return partial.ConfigLayer(i)
	}
	if layer, ok := i.digestMap[h]; ok {
		return layer, nil
	}
	return i.base.BlobByDigest(h)
}
