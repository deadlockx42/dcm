//
//   Copyright 2017 Deadlock X42 <deadlock.x42@gmail.com>
//
//   Licensed under the Apache License, Version 2.0 (the "License");
//   you may not use this file except in compliance with the License.
//   You may obtain a copy of the License at
//
//       http://www.apache.org/licenses/LICENSE-2.0
//
//   Unless required by applicable law or agreed to in writing, software
//   distributed under the License is distributed on an "AS IS" BASIS,
//   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//   See the License for the specific language governing permissions and
//   limitations under the License.
//

package main

import (
	"encoding/json"
	"io"
)

type Dcm interface {
	Name() string
	Title() string
	DataCenter() DataCenter
}

func New(r io.Reader) (Dcm, error) {
	var d dcm
	for {
		err := json.NewDecoder(r).Decode(&d)
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}
	}
	return &d, nil
}

type dcm struct {
	Name_       string      `json:"name"`
	Title_      string      `json:"title"`
	DataCenter_ *datacenter `json:"datacenter"`
}

func (d *dcm) Name() string {
	return d.Name_
}

func (d *dcm) Title() string {
	return d.Title_
}

func (d *dcm) DataCenter() DataCenter {
	return d.DataCenter_
}
