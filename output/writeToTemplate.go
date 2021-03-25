// Copyright 2021 Seiichi Ariga <seiichi.ariga@gmail.com>
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
package output

import (
	"cert-template/input"
	"fmt"
)

func WriteToTemplate(filename string, dt input.KVArray) error {
	data, err := input.ReadJSON(filename)
	if err != nil {
		panic("Error")
	}
	fmt.Println(data)

	return nil
}
