package coverage

import (
	"path/filepath"
	"runtime"
	"testing"

	"github.com/Liu-Chunhui/line-coverage/pkg/fileparser"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMapProfileToBranch(t *testing.T) {
	_, filename, _, _ := runtime.Caller(0)
	execPath := filepath.Dir(filename)
	testfile := filepath.Join(execPath, "../../test/data/testcodefile")

	lines, _ := fileparser.ReadLines(testfile)
	base := filepath.Join(execPath, "../../")

	tests := []struct {
		name        string
		profileLine string
		expected    []*branch
	}{
		{
			name:        "uncoveredLines",
			profileLine: "github.com/Liu-Chunhui/line-coverage/test/data/testcodefile:16.16,18.3 1 0",
			expected: []*branch{
				{
					Start:   17,
					Finish:  17,
					Covered: false,
				},
			},
		},
		{
			name:        "finishingLineIs, nil",
			profileLine: "github.com/Liu-Chunhui/line-coverage/test/data/testcodefile:59.2,59.19 1 60",
			expected: []*branch{
				{
					Start:   59,
					Finish:  59,
					Covered: true,
				},
			},
		},
		{
			name:        "WhenStatementIs0ButExecuteTimesAreNot0_ThenShouldMarkedAsCovered",
			profileLine: "github.com/Liu-Chunhui/line-coverage/test/data/testcodefile:15.14,15.14 0 2",
			expected: []*branch{
				{
					Start:   15,
					Finish:  15,
					Covered: true,
				},
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			profile, err := mapLineToCoverageProfile(tt.profileLine, "github.com/Liu-Chunhui/line-coverage", base)
			require.NoError(t, err)

			target, branches := convertProfileToBranch(profile, lines)
			require.NotNil(t, target)
			assert.NotEmpty(t, target)
			assert.Equal(t, tt.expected, branches)
		})
	}

}

func TestStartPositionAdjustment(t *testing.T) {
	tests := []struct {
		name               string
		lines              []string
		startLine          int
		startPos           int
		expectedChar       byte
		expectedAdjustment int
	}{
		{
			name:               "WhenStartPositionIsNotNewLineChar_ThenReturnNoAdjustment",
			lines:              []string{"\t\tif values, ok := maps[profile.TargetFilename]; ok {\n"},
			startLine:          1,
			startPos:           32,
			expectedChar:       '.',
			expectedAdjustment: 0,
		},
		{
			name: "WhenStartPositionIsNewLineChar_ThenReturnExpectedAdjustment",
			lines: []string{
				"\t\tif values, ok := maps[profile.TargetFilename]; ok {\n",
				"some code",
			},
			startLine:          1,
			startPos:           54,
			expectedChar:       '\n',
			expectedAdjustment: 1,
		},
		{
			name: "WhenStartPositionIsOpenBracketLeadingWithNewLineChar_ThenReturnExpectedAdjustment",
			lines: []string{
				"func Get(ctx context.Context) (*Identity, error) {\n",
				"some code",
			},
			startLine:          1,
			startPos:           50,
			expectedChar:       '{',
			expectedAdjustment: 1,
		},
		{
			name: "WhenStartPositionIsNewLineCharAndNextLineIsNewLine_ThenReturnExpectedAdjustment",
			lines: []string{
				"func Get(ctx context.Context) (*Identity, error) {\n",
				"\n",
				"some code",
			},
			startLine:          1,
			startPos:           51,
			expectedChar:       '\n',
			expectedAdjustment: 2,
		},
		{
			name: "WhenStartPositionIsNewLineCharAndNextLineIsBracketLeadingWithNewLine_ThenReturnExpectedAdjustment",
			lines: []string{
				"func Get(ctx context.Context) (*Identity, error) {\n",
				"	{\n",
				"some code",
			},
			startLine:          1,
			startPos:           51,
			expectedChar:       '\n',
			expectedAdjustment: 2,
		},
		{
			name: "WhenStartPositionIsNewLineCharAndThisLineIsCase_ThenReturnExpectedAdjustment",
			lines: []string{
				"\tcase http.StatusInternalServerError, http.StatusServiceUnavailable, http.StatusBadGateway, http.StatusGatewayTimeout:\n",
			},
			startLine:          1,
			startPos:           119,
			expectedChar:       '\n',
			expectedAdjustment: 0,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got := startLineAdjustment(tt.lines, tt.startLine, tt.startPos)
			require.Equal(t, tt.expectedChar, tt.lines[tt.startLine-1][tt.startPos-1])
			assert.Equal(t, tt.expectedAdjustment, got)
		})
	}
}

func TestFinishLineAdjustment(t *testing.T) {
	tests := []struct {
		name               string
		lines              []string
		finishLine         int
		endPos             int
		expectedChar       byte
		expectedAdjustment int
	}{
		{
			name:               "WhenEndPositionCodeWith'}'_ThenReturnNoAdjustment",
			lines:              []string{"} else if a < b { a += 1 }\n"},
			finishLine:         1,
			endPos:             26,
			expectedChar:       '}',
			expectedAdjustment: 0,
		},
		{
			name:               "WhenEndPositionIsNewLineChar_ThenReturnNoAdjustment",
			lines:              []string{"} else if a < b { a += 1 }\n"},
			finishLine:         1,
			endPos:             27,
			expectedChar:       '\n',
			expectedAdjustment: 0,
		},
		{
			name: "WhenLineIsEndingWith'})'_ThenReturnExpectedAdjustment",
			lines: []string{
				"\t\t\t\t\t Some Code,\n",
				"\t\t\t\t})\n",
				"\t\t\t}\n"},
			finishLine:         3,
			endPos:             5,
			expectedChar:       '\n',
			expectedAdjustment: 1,
		},
		{
			name: "WhenEndPositionIsOnlyBracketLeadingWithNewLineChar_ThenReturnExpectedAdjustment",
			lines: []string{
				"code }\n",
				"}\n"},
			finishLine:         2,
			endPos:             2,
			expectedChar:       '\n',
			expectedAdjustment: 1,
		},
		{
			name: "WhenEndPositionIsMultipleBracketsLeadingWithNewLineChar_ThenReturnExpectedAdjustment",
			lines: []string{
				"code }\n",
				"}}}}\n"},
			finishLine:         2,
			endPos:             5,
			expectedChar:       '\n',
			expectedAdjustment: 1,
		},
		{
			name: "WhenLastTwoLinesAreBracketLeadingWithNewLineChar_ThenReturnExpectedAdjustment",
			lines: []string{
				"		some code here }\n",
				"	}\n",
				"}\n"},
			finishLine:         3,
			endPos:             2,
			expectedChar:       '\n',
			expectedAdjustment: 2,
		},
		{
			name: "WhenEndWithNothingBut'}()'_ThenReturnExpectedAdjustment",
			lines: []string{
				"	return some code here\n",
				"\t\t\t}()\n"},
			finishLine:         2,
			endPos:             5,
			expectedChar:       '(',
			expectedAdjustment: 1,
		},
		{
			name: "WhenEndWithNothingBut'} else {'_ThenReturnExpectedAdjustment",
			lines: []string{
				"		if err != nil {\n",
				"			serverErr = some code\n",
				"		} else {\n"},
			finishLine:         3,
			endPos:             4,
			expectedChar:       ' ',
			expectedAdjustment: 1,
		},
		{
			name: "WhenEndWithNothingBut'},'_ThenReturnExpectedAdjustment",
			lines: []string{
				"\t\t\treturn nil\n",
				"\t\t},\n"},
			finishLine:         2,
			endPos:             4,
			expectedChar:       ',',
			expectedAdjustment: 1,
		},
		{
			name: "WhenEndWith'}'AndPreviousLineIs'},'_ThenReturnExpectedAdjustment",
			lines: []string{
				"\t\t\t},\n",
				"\t\t}\n"},
			finishLine:         2,
			endPos:             4,
			expectedChar:       '\n',
			expectedAdjustment: 1,
		},
		{
			// '}' closes for 'struct {'
			name: "WhenEndWith'})AndPreviousLineEndWith','_ThenReturnNoAdjustment",
			lines: []string{
				"\t\t some code,\n",
				"\t})\n"},
			finishLine:         2,
			endPos:             3,
			expectedChar:       ')',
			expectedAdjustment: 0,
		},
		{
			// '}' closes for 'function {'
			name: "WhenEndWith'})AndPreviousLineNotEndWith','_ThenReturnExpectedAdjustment",
			lines: []string{
				"	return some code here\n",
				"\t})\n"},
			finishLine:         2,
			endPos:             3,
			expectedChar:       ')',
			expectedAdjustment: 1,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got := finishLineAdjustment(tt.lines, tt.finishLine, tt.endPos)
			require.Equal(t, tt.expectedChar, tt.lines[tt.finishLine-1][tt.endPos-1])
			assert.Equal(t, tt.expectedAdjustment, got)
		})
	}
}

func TestSplitByNewLine(t *testing.T) {
	tests := []struct {
		name        string
		codeInLines []string
		startLine   int
		endLine     int
		covered     bool
		expected    []*branch
	}{
		{
			name: "WhenThereIsNoJustNewLines_ThenShouldReturnOneBranch",
			codeInLines: []string{
				"1\n",
				"2\n",
				"3\n",
				"4\n",
				"5\n",
				"6\n",
				"7\n",
				"8\n",
			},
			startLine: 3,
			endLine:   7,
			covered:   true,
			expected: []*branch{
				{
					Start:   3,
					Finish:  7,
					Covered: true,
				},
			},
		},
		{
			name: "WhenThereIsOneEntireLineIsNewLines_ThenShouldReturnExpectedBranch",
			codeInLines: []string{
				"1\n",
				"2\n",
				"3\n",
				"4\n",
				"\n",
				"6\n",
				"7\n",
				"8\n",
			},
			startLine: 1,
			endLine:   8,
			covered:   true,
			expected: []*branch{
				{
					Start:   1,
					Finish:  4,
					Covered: true,
				},
				{
					Start:   6,
					Finish:  8,
					Covered: true,
				},
			},
		},
		{
			name: "WhenThereAreMultipleEntireLineIsNewLines_ThenShouldReturnExpectedBranch",
			codeInLines: []string{
				"1\n",
				"2\n",
				"3\n",
				"4\n",
				"\n",
				"6\n",
				"7\n",
				"8\n",
				"\n",
				"10\n",
				"11\n",
				"12\n",
			},
			startLine: 2,
			endLine:   11,
			covered:   true,
			expected: []*branch{
				{
					Start:   2,
					Finish:  4,
					Covered: true,
				},
				{
					Start:   6,
					Finish:  8,
					Covered: true,
				},
				{
					Start:   10,
					Finish:  11,
					Covered: true,
				},
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got := splitByNewLine(tt.codeInLines, tt.startLine, tt.endLine, tt.covered)
			assert.Equal(t, tt.expected, got)
		})
	}
}
