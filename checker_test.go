package main

import "testing"

func TestGetNoValidationReason(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		comments string
		want     string
	}{
		{
			name:     "empty input",
			comments: "",
			want:     "",
		},
		{
			name:     "single line with marker",
			comments: " No Validation Reason: business decision",
			want:     "business decision",
		},
		{
			name:     "marker on second line",
			comments: " unrelated comment\n No Validation Reason: legacy field",
			want:     "legacy field",
		},
		{
			name:     "marker with empty reason",
			comments: " No Validation Reason: ",
			want:     "",
		},
		{
			name:     "marker not present",
			comments: " a comment\n another comment",
			want:     "",
		},
		{
			name:     "first matching marker wins",
			comments: " No Validation Reason: first\n No Validation Reason: second",
			want:     "first",
		},
		{
			name:     "marker without leading space ignored",
			comments: "No Validation Reason: missing leading space",
			want:     "",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			got := getNoValidationReason(tc.comments)
			if got != tc.want {
				t.Errorf("getNoValidationReason(%q) = %q, want %q", tc.comments, got, tc.want)
			}
		})
	}
}
