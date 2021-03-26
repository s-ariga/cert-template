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

/* JSONからテンプレートkeyとコンテンツを読み込んで、Wordのテンプレートに書き出す */

package main

import (
	"cert-template/input"
	"cert-template/output"
	"flag"
)

func main() {

	var inputFile string
	var templateFile string
	var outputFile string
	// コマンドラインオプションを追加する
	flag.StringVar(&inputFile, "i", "test.json", "入力ファイル(JSON)")
	flag.StringVar(&templateFile, "t", "test.docx", "テンプレートファイル(docx)")
	flag.StringVar(&outputFile, "o", "output.docx", "出力ファイル名(docx)")

	// オプションの解析
	flag.Parse()

	inputData, err := input.ReadJSON(inputFile)
	if err != nil {
		panic("Input error")
	}

	err = output.WriteToTemplate(templateFile, outputFile, inputData)
	if err != nil {
		panic("output error")
	}
}
