package examples

import "testing"

func TestSplitEmail_WithError(t *testing.T) {
	tests := []struct {
		input      string
		username   string
		domain     string
		shouldFail bool
	}{
		{"drogon@dragonstone.com", "drogon", "dragonstone.com", false},
		{"rhaenyra@targaryen.com", "rhaenyra", "targaryen.com", false},
		{"@domain.com", "", "", true},
		{"user@", "", "", true},
		{"invalid", "", "", true},
		{"", "", "", true},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			u, d, err := splitEmail(tt.input)

			if tt.shouldFail {
				if err == nil {
					t.Errorf("expected error for input=%q", tt.input)
				}
				return
			}

			if err != nil {
				t.Errorf("unexpected error for input=%q: %v", tt.input, err)
				return
			}

			if u != tt.username || d != tt.domain {
				t.Errorf("got (%q, %q), want (%q, %q)",
					u, d, tt.username, tt.domain)
			}
		})
	}
}