package examples

import "testing"

func TestSplitEmail_Hard(t *testing.T) {
	tests := []struct {
		name       string
		input      string
		username   string
		domain     string
		shouldFail bool
	}{
		{"valid_simple", "drogon@dragonstone.com", "drogon", "dragonstone.com", false},
		{"trailing_space", "jon@winterfell.com ", "jon", "winterfell.com", false},
		{"leading_space", " jon@winterfell.com", "jon", "winterfell.com", false},
		{"multiple_at", "a@b@c.com", "a", "b@c.com", false},
		{"empty_string", "", "", "", true},
		{"no_at", "invalidemail", "", "", true},
		{"missing_username", "@domain.com", "", "", true},
		{"missing_domain", "user@", "", "", true},
		{"special_chars", "user+tag@domain.co.uk", "user+tag", "domain.co.uk", false},
		{"subdomain", "user@mail.server.com", "user", "mail.server.com", false},
		{"only_at", "@", "", "", true},
		{"spaces_only", "   ", "", "", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
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

			// trim spaces for comparison
			uTrim := u
			dTrim := d
			if uTrim != tt.username || dTrim != tt.domain {
				t.Errorf("got (%q, %q), want (%q, %q)", uTrim, dTrim, tt.username, tt.domain)
			}
		})
	}
}