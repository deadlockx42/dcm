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
	"fmt"
	"os"
	"strings"
)

func main() {

	const cfg = `
		{
			"name": "mock",
			"title": "Testing, 1, 2, 3",

			"datacenter": {
				"name": "dc"
			}
		}
	`

	void, err := New(strings.NewReader(cfg))
	if err != nil {
		fmt.Printf(err.Error())
		os.Exit(1)
	}

	fmt.Printf("name:  %s\n", void.Name())
	fmt.Printf("title: %s\n", void.Title())
	fmt.Printf("dc:    %s\n", void.DataCenter().Name())
}
