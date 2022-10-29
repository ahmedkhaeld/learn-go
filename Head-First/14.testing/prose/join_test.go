package prose

import (
	"fmt"
	"testing"
)

type testData struct {
	list []string
	exp  string
}

func TestJoinWithCommas(t *testing.T) {
	tests := []testData{
		{list: []string{}, exp: ""},
		{list: []string{"apple"}, exp: "apple"},
		{list: []string{"apple", "orange"}, exp: "apple and orange"},
		{list: []string{"apple", "orange", "pear"}, exp: "apple, orange, and pear"},
	}

	for _, test := range tests {
		got := JoinWithCommas(test.list)
		if got != test.exp {
			t.Errorf(errorString(test.list, got, test.exp))
		}
	}
}

func errorString(list []string, got, exp string) string {
	return fmt.Sprintf("JoinWithCommas(%#v) gives: \"%s\", Expected: \"%s \"", list, got, exp)
}
