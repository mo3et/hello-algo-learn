package complexity

import (
	"fmt"
	"testing"
)

func TestTimeComplexity(t *testing.T) {
	var tests = []struct {
		name string
		n    int
	}{
		{"testcase1", 8},
		{"testcase2", 10},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fmt.Println("Input data is n =", tt.n)
			count := expRecur(tt.n)
			fmt.Println("calculate num is ", count)
		})
	}
}
