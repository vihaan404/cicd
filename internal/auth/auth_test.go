package auth

import (
	"errors"
	"net/http"
	"reflect"
	"testing"
)

var headerNormal = http.Header{
	"Authorization": []string{"ApiKey 123"},
}
var headerMore = http.Header{
	"Authorization": []string{"piKey 123 something"},
}
var headerEmpty = http.Header{
	"Authorization": []string{""},
}
var headerEmpty2 = http.Header{}

func TestGetAPIKey(t *testing.T) {
	type test struct {
		inputHeader http.Header
		want        string
		errWant     error
	}
	tests := []test{
		{inputHeader: headerNormal, want: "123", errWant: nil},
		{inputHeader: headerMore, want: "", errWant: ErrMalFormed},
		{inputHeader: headerEmpty, want: "", errWant: ErrNoAuthHeaderIncluded},
		{inputHeader: headerEmpty2, want: "", errWant: ErrNoAuthHeaderIncluded},
	}
	for _, tc := range tests {

		got, errGot := GetAPIKey(tc.inputHeader)
		if !reflect.DeepEqual(got, tc.want) {
			t.Fatalf("expected: %v, got: %v", tc.want, got)
		}
		if !errors.Is(errGot, tc.errWant) {
			t.Fatalf("expected Error: %v, got: %v", tc.errWant, errGot)
		}
	}
}
