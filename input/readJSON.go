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

package input

import (
	"encoding/json"
	"io/ioutil"
)

// JSON ファイルの読み込み
func ReadJSON(filename string) (KVArray, error) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	var kv KVArray
	err = json.Unmarshal([]byte(data), &kv)
	if err != nil {
		return kv, err
	}
	return kv, nil
}
