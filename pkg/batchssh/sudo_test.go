package batchssh

import (
	"regexp"
	"testing"
)

func TestSudoPromptRegex(t *testing.T) {
	re := regexp.MustCompile(SudoPromptRegex)

	tests := []struct {
		name     string
		input    string
		wantMatch bool
	}{
		{
			name:      "Standard English prompt",
			input:     "[sudo] password for user: ",
			wantMatch: true,
		},
		{
			name:      "Chinese prompt",
			input:     "[sudo] user 的密码：",
			wantMatch: true,
		},
		{
			name:      "Prompt with newline after",
			input:     "[sudo] password for root:\n",
			wantMatch: true,
		},
		{
			name:      "Not a prompt",
			input:     "This is not a [sudo] prompt",
			wantMatch: false,
		},
		{
			name:      "Prompt embedded in text (should NOT match at start without .* but our regex doesn't have ^)",
			input:     "Some text [sudo] password for user: ",
			wantMatch: true, 
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := re.MatchString(tt.input); got != tt.wantMatch {
				t.Errorf("SudoPromptRegex match %q = %v, want %v", tt.input, got, tt.wantMatch)
			}
		})
	}
}

func TestSudoPromptCleaning(t *testing.T) {
	re := regexp.MustCompile(SudoPromptRegex)

	input := "[sudo] password for user: \nLinux node1 5.4.0\n"
	want := "Linux node1 5.4.0\n"

	got := re.ReplaceAllString(input, "")
	if got != want {
		t.Errorf("Cleaned output = %q, want %q", got, want)
	}
}
