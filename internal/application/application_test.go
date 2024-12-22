package application_test

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/arbuzick57/calc_go/internal/application"
)

func TestCalcHandler(t *testing.T) {
	tests := []struct {
		name           string
		method         string
		body           interface{}
		expectedStatus int
		expectedResult string
	}{
		{
			name:           "Valid Expression",
			method:         http.MethodPost,
			body:           map[string]string{"expression": "5+2"},
			expectedStatus: http.StatusOK,
			expectedResult: `{"result":"7"}`,
		},
		{
			name:           "Empty Expression",
			method:         http.MethodPost,
			body:           map[string]string{"expression": ""},
			expectedStatus: http.StatusUnprocessableEntity,
			expectedResult: `{"error":"empty expression"}`,
		},
		{
			name:           "Not Numbers",
			method:         http.MethodPost,
			body:           map[string]string{"expression": "a"},
			expectedStatus: http.StatusUnprocessableEntity,
			expectedResult: `{"error":"not numbers"}`,
		},
		{
			name:           "Empty Body",
			method:         http.MethodPost,
			body:           nil,
			expectedStatus: http.StatusBadRequest,
			expectedResult: `{"error":"incorrect request"}`,
		},
		{
			name:           "Division By Zero",
			method:         http.MethodPost,
			body:           map[string]string{"expression": "5/0"},
			expectedStatus: http.StatusUnprocessableEntity,
			expectedResult: `{"error":"division by zero is forbidden"}`,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			var requestBody []byte
			if test.body != nil {
				var err error
				requestBody, err = json.Marshal(test.body)
				if err != nil {
					t.Fatal(err)
				}
			}
			req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(requestBody))
			w := httptest.NewRecorder()
			application.CalcHandler(w, req)
			res := w.Result()
			defer res.Body.Close()
			data, err := io.ReadAll(res.Body)
			if err != nil {
				t.Errorf("Error: %v", err)
			}
			if res.StatusCode != test.expectedStatus {
				t.Errorf("Expected %d status code, but got %d", test.expectedStatus, res.StatusCode)
			}
			if res.StatusCode != http.StatusInternalServerError && string(data) != test.expectedResult {
				t.Errorf("Expected %s, but got %s", test.expectedResult, string(data))
			}
		})
	}
}
