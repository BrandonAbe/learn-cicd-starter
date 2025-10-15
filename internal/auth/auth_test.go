package auth

import (
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	tests := []struct {
		name        string
		headers     http.Header
		wantKey     string
		expectError bool
		expectedErr error
	}{
		{
			name:        "valid API key header",
			headers:     http.Header{"Authorization": {"ApiKey abc123"}},
			wantKey:     "abc123",
			expectError: false,
		},
		{
			name:        "missing authorization header",
			headers:     http.Header{},
			expectError: true,
			expectedErr: ErrNoAuthHeaderIncluded,
		},
		{
			name:        "malformed header - missing ApiKey prefix",
			headers:     http.Header{"Authorization": {"Bearer abc123"}},
			expectError: true,
		},
		{
			name:        "malformed header - missing key value",
			headers:     http.Header{"Authorization": {"ApiKey"}},
			expectError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetAPIKey(tt.headers)

			if tt.expectError {
				if err == nil {
					t.Fatalf("expected an error but got nil")
				}
				if tt.expectedErr != nil && err != tt.expectedErr {
					t.Errorf("expected error %v, got %v", tt.expectedErr, err)
				}
				return
			}

			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}

		if got != tt.wantKey {
				t.Errorf("expected key %q, got %q", tt.wantKey, got)
			}
		})
	}
}
