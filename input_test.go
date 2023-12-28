package shelly

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestInputCheckExpressionResponse(t *testing.T) {
	tcs := []struct {
		name   string
		in     string
		expect InputCheckExpressionResponse
	}{
		{
			name: "website example",
			in: `{
				"results": [
				  [
					1,
					4.333333
				  ],
				  [
					2,
					4.666667
				  ],
				  [
					3,
					5
				  ]
				]
			  }`,
			expect: InputCheckExpressionResponse{
				Results: []InputCheckExpressionResult{
					{
						Input:  Float64Ptr(1),
						Output: Float64Ptr(4.333333),
					},
					{
						Input:  Float64Ptr(2),
						Output: Float64Ptr(4.666667),
					},
					{
						Input:  Float64Ptr(3),
						Output: Float64Ptr(5),
					},
				},
			},
		},
		{
			name: "result includes an error",
			in: `{
				"results": [
				  [
					1,
					4.333333,
					"this is an error"
				  ]
				]
			  }`,
			expect: InputCheckExpressionResponse{
				Results: []InputCheckExpressionResult{
					{
						Input:  Float64Ptr(1),
						Output: Float64Ptr(4.333333),
						Error:  StrPtr("this is an error"),
					},
				},
			},
		},
		{
			name: "result includes an null output",
			in: `{
				"results": [
				  [
					1,
					null
				  ]
				]
			  }`,
			expect: InputCheckExpressionResponse{
				Results: []InputCheckExpressionResult{
					{
						Input:  Float64Ptr(1),
						Output: nil,
					},
				},
			},
		},
	}
	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			var got InputCheckExpressionResponse
			require.NoError(t, json.Unmarshal([]byte(tc.in), &got))
			assert.Equal(t, tc.expect, got)
		})
	}
}
