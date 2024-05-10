package auth

import (
	"net/http"
	"reflect"
	"testing"
)

func createHttpHeader() http.Header {
	return http.Header{
		"Content-Type":  []string{"application/json"},
		"Authorization": []string{"ApiKey 123"},
	}
}

var header = createHttpHeader()

func TestGetAPIKey(t *testing.T) {
	got, errGot := GetAPIKey(header)
	want := "123"
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("expected: %v, got: %v", want, got)
	}
	if errGot != nil {

		t.Fatalf("expected Error: %v, got: %v", want, got)
	}
}
