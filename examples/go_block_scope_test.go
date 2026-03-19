package examples

import "testing"

func TestSplitEmail_Hard_Prod(t *testing.T) {
	tests := []struct {
		name       string
		input      string
		wantUser   string
		wantDomain string
		shouldFail bool
	}{
		// valid
		{"valid_simple", "drogon@dragonstone.com", "drogon", "dragonstone.com", false},
		{"trailing_space", "jon@winterfell.com ", "jon", "winterfell.com", false},
		{"leading_space", " jon@winterfell.com", "jon", "winterfell.com", false},
		{"special_chars", "user+tag@domain.co.uk", "user+tag", "domain.co.uk", false},
		{"subdomain", "user@mail.server.com", "user", "mail.server.com", false},

		// invalid
		{"multiple_at", "a@b@c.com", "", "", true},
		{"empty_string", "", "", "", true},
		{"no_at", "invalidemail", "", "", true},
		{"missing_username", "@domain.com", "", "", true},
		{"missing_domain", "user@", "", "", true},
		{"only_at", "@", "", "", true},
		{"spaces_only", "   ", "", "", true},
		{"username_with_space", "user name@domain.com", "", "", true}, // username не може містити пробіл
		{"domain_with_space", "user@domain .com", "", "", true},       // domain не може містити пробіл
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

			if u != tt.wantUser || d != tt.wantDomain {
				t.Errorf("got (%q, %q), want (%q, %q)", u, d, tt.wantUser, tt.wantDomain)
			}
		})
	}
}