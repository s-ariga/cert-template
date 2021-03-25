<!--
 Copyright 2021 Seiichi Ariga <seiichi.ariga@gmail.com>

 Licensed under the Apache License, Version 2.0 (the "License");
 you may not use this file except in compliance with the License.
 You may obtain a copy of the License at

     http://www.apache.org/licenses/LICENSE-2.0

 Unless required by applicable law or agreed to in writing, software
 distributed under the License is distributed on an "AS IS" BASIS,
 WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 See the License for the specific language governing permissions and
 limitations under the License.
-->

# cert-template

WordのテンプレートファイルにJSONを使って、文字列を流し込む

## 仕様

Wordファイル内の文字列を、別の文字列で置き換えるので、テンプレートとして使える。

## JSON

```JSON
[{"Key": "KEY_1",
"Content": "CONTENT_1"},
{"Key": "KEY_2",
"Content": "CONTENT_2"}
]
```

KeyとContentの組み合わせの配列。(GoのArray)

KeyをContentで置き換えます。

## 依存

github.com/nguyenthenguyen/docx

