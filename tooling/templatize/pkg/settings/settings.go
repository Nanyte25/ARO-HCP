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

package settings

import (
	"os"

	"sigs.k8s.io/yaml"
)

type Settings struct {
	Environments []Environment `json:"environments"`
}

type Environment struct {
	Name        string     `json:"name"`
	Description string     `json:"description"`
	Defaults    Parameters `json:"defaults"`
}

type Parameters struct {
	Cloud               string `json:"cloud"`
	Region              string `json:"region"`
	CxStamp             int    `json:"cxStamp"`
	RegionStampTemplate string `json:"regionStampTemplate"`
}

func Load(path string) (*Settings, error) {
	raw, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var out Settings
	return &out, yaml.Unmarshal(raw, &out)
}
