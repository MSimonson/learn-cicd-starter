package auth

import (
	"net/http"
	"reflect"
	"testing"
)

func TestGetApiKey(t *testing.T) {
	tests := map[string]struct {
		input   http.Header
		want    string
		wantErr error
	}{
		"simple": {
			input: http.Header{
				"Authorization": {"ApiKey 12345"},
			},
			want:    "12345",
			wantErr: nil,
		},
		"malformed": {
			input: http.Header{
				"Authorization": {"ApKey 54321"},
			},
			want:    "",
			wantErr: ErrMalformedAuthHeader,
		},
		"no header": {
			input:   http.Header{"": {}},
			want:    "",
			wantErr: ErrNoAuthHeaderIncluded,
		},
		"mispelled": {
			input: http.Header{
				"Auhtrozation": {"ApiKey 54321"},
			},
			want:    "",
			wantErr: ErrNoAuthHeaderIncluded,
		},
	}
	for name, testCase := range tests {
		t.Run(name, func(t *testing.T) {
			got, gotErr := GetAPIKey(testCase.input)
			if !reflect.DeepEqual(testCase.want, got) {
				t.Fatalf("expected %v, got %v", testCase.want, got)
			}
			if testCase.wantErr != gotErr {
				t.Fatalf("expected %v, got %v", testCase.wantErr, gotErr)
			}
		})
	}
}
