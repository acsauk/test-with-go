package underscore

import (
	"testing"
)

// func TestCamel(t *testing.T) {
// 	testCases := []struct{
// 		arg string
// 		want string
// 	}{
// 		{"thisIsACamelCaseString", "this_is_a_camel_case_string"},
// 		{"with a space", "with a space"},
// 		{"endsWithA", "ends_with_a"},
// 	}

// 	for _, tc := range testCases {
// 		got := Camel(tc.arg)
// 		if got != tc.want {
// 			t.Errorf("Camel(%q) = %q; want %q", tc.arg, got, tc.want)
// 		}
// 	}
// }

func TestCamel(t *testing.T) {
	type args struct {
		str string
	}
	tests := map[string]struct {
		args args
		want string
	}{
		"Simple": {args{"thisIsACamelCaseString"}, "this_is_a_camel_case_string"},
		"Spaces": {args{"with a space"}, "with a space"},
		"Ending with A": {args{"endsWithA"}, "ends_with_a"},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			if got := Camel(tt.args.str); got != tt.want {
				t.Fatalf("Camel() = %v, want %v", got, tt.want)
			}
		})
	}
}
