package api_test

import (
	"avoxi-interview/api"
	"avoxi-interview/models"
	"bytes"

	"github.com/nsf/jsondiff"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestIpCountryCheck(t *testing.T) {

	type args struct {
		Payload  models.Payload
		IPStatus models.IpStatus
	}
	tests := []struct {
		name             string
		Path             string
		Payload          string
		ExpectedJSON     string
		ErrorExpected    bool
		HttpCodeExpected int
	}{
		{
			name:             "first test",
			Path:             "/ipcountrycheck",
			Payload:          `{"Ip": "14.137.32.0", "Countries": ["Italy", "Papua New Guinea"]}`,
			ExpectedJSON:     `{"Ip": "14.137.32.0", "Approved": true}`,
			ErrorExpected:    false,
			HttpCodeExpected: http.StatusOK,
		},
		{
			name:             "case misMAtcH",
			Path:             "/ipcountrycheck",
			Payload:          `{"Ip": "14.137.32.0", "Countries": ["ItAly", "Papua New GuInea"]}`,
			ExpectedJSON:     `{"Ip": "14.137.32.0", "Approved": true}`,
			ErrorExpected:    false,
			HttpCodeExpected: http.StatusOK,
		},
		{
			name:             "bad ip",
			Path:             "/ipcountrycheck",
			Payload:          `{"Ip": "14.137.32.0.4.3.4.5", "Countries": ["ItAly", "Papua New GuInea"]}`,
			ExpectedJSON:     "",
			ErrorExpected:    false,
			HttpCodeExpected: http.StatusBadRequest,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req, err := http.NewRequest("GET", tt.Path, bytes.NewReader([]byte(tt.Payload)))
			if err != nil {
				t.Fatal(err)
			}
			rr := httptest.NewRecorder()
			handler := http.HandlerFunc(api.IpCountryCheck)
			handler.ServeHTTP(rr, req)
			if status := rr.Code; status != tt.HttpCodeExpected {
				t.Errorf("handler returned wrong status code: got %v want %v",
					status, http.StatusOK)
			}

			body := rr.Body.String()
			if tt.ExpectedJSON != "" {
				options := jsondiff.DefaultJSONOptions()
				val, diff := jsondiff.Compare([]byte(body), []byte(tt.ExpectedJSON), &options)
				if val != 0 {
					t.Fatalf("JSON request does not match:\n%s\n", diff)
				}
			}
		})
	}
}
