// Copyright 2020 ZUP IT SERVICOS EM TECNOLOGIA E INOVACAO SA
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

package call_test

import (
	"testing"

	"github.com/ZupIT/horusec-engine/internal/ir"
	"github.com/ZupIT/horusec-engine/internal/utils/testutil"
	"github.com/ZupIT/horusec-engine/semantic/analysis"
	"github.com/ZupIT/horusec-engine/semantic/analysis/call"
)

type fakeAnalyzerValue struct {
	returnValue bool
}

func (a fakeAnalyzerValue) Run(_ ir.Value) bool { return a.returnValue }

func TestAnalyzerCall(t *testing.T) {
	src := `function f() { eval("console.log('eval')") }`

	testcases := []testutil.TestCaseAnalyzer{
		{
			Name: "MatchArguments",
			Src:  src,
			Analyzer: &call.Analyzer{
				Name:      "eval",
				ArgsIndex: 1,
				ArgValue:  fakeAnalyzerValue{false},
			},
			ExpectedIssues: []analysis.Issue{
				{
					Filename:    "MatchArguments",
					StartOffset: 15,
					EndOffset:   42,
					Line:        1,
					Column:      15,
				},
			},
		},
		{
			Name: "NotMatchArguments",
			Src:  src,
			Analyzer: &call.Analyzer{
				Name:      "eval",
				ArgsIndex: 1,
				ArgValue:  fakeAnalyzerValue{true},
			},
			ExpectedIssues: []analysis.Issue{},
		},
		{
			Name: "NotMatchFunctionName",
			Src:  src,
			Analyzer: &call.Analyzer{
				Name:      "fs.readFile",
				ArgsIndex: 1,
				ArgValue:  fakeAnalyzerValue{true},
			},
			ExpectedIssues: []analysis.Issue{},
		},
	}
	testutil.TestAnalayzer(t, testcases)
}