package auth

import (
	"bytes"
	"net/http"
	"testing"
)

func Test_GetAPIKey(t *testing.T) {

	type test struct {
		inputAPIKey  string
		outputString string
		outputErr    error
	}

	dummydata := []byte("dummy data")
	req, err := http.NewRequest("GET", "/users", bytes.NewBuffer(dummydata))
	if err != nil {
		t.Fatalf("Failure generating dummy request: %v", err)
	}

	tests := []test{
		{inputAPIKey: "", outputString: "", outputErr: ErrNoAuthHeaderIncluded},
		{inputAPIKey: "ApiKey", outputString: "", outputErr: ErrMalformedAuthorization},
		{inputAPIKey: "ApiKey 1", outputString: "1", outputErr: nil},
		{inputAPIKey: "Apikey 1", outputString: "", outputErr: ErrMalformedAuthorization},
		{inputAPIKey: "ApiKey aeujgnhsapodgnpowakrng", outputString: "aeujgnhsapodgnpowakrng", outputErr: nil},
		{inputAPIKey: "ApiKey aeujgnhs apodgnpowakrng", outputString: "aeujgnhs", outputErr: nil},
	}

	for i, test := range tests {
		req.Header.Set("Authorization", test.inputAPIKey)
		key, err := GetAPIKey(req.Header)
		if !(key == test.outputString && err == test.outputErr) {
			t.Fatalf("test %d: expected output \"%v\", %v, got: \"%v\", %v", i+1, test.outputString, test.outputErr, key, err)
		}
	}

}

// need test cases for "", valid key, keu with multiple delimiters (" ")
