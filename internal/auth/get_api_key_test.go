package auth

import (
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	tests := map[string]struct {
		headerKey   string
		headerValue string
		wantKey     string
		wantErr     bool
	}{
		"valid api key": {
			headerKey:   "Authorization",
			headerValue: "ApiKey my-secret-api-key",
			wantKey:     "my-secret-api-key",
			wantErr:     false,
		},
		"missing authorization header": {
			headerKey:   "",
			headerValue: "",
			wantKey:     "",
			wantErr:     true,
		},
		"malformed header prefix": {
			headerKey:   "Authorization",
			headerValue: "Bearer my-secret-api-key",
			wantKey:     "",
			wantErr:     true,
		},
		"missing key value": {
			headerKey:   "Authorization",
			headerValue: "ApiKey",
			wantKey:     "",
			wantErr:     true,
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			headers := http.Header{}
			if tc.headerKey != "" {
				headers.Set(tc.headerKey, tc.headerValue)
			}

			got, err := GetAPIKey(headers)
			if (err != nil) != tc.wantErr {
				t.Fatalf("expected error: %v, got: %v", tc.wantErr, err)
			}
			if got != tc.wantKey {
				t.Errorf("expected key %q, got %q", tc.wantKey, got)
			}
		})
	}
}
