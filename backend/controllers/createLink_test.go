package controllers_test

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/kazion500/secure-share/app"
	"github.com/kazion500/secure-share/types"
	"github.com/stretchr/testify/assert"
)

func TestCreateLinkController(t *testing.T) {
	testCases := []struct {
		name        string
		response    types.ResponseType[any]
		requestData []byte
		statusCode  int
	}{
		{
			name:       "Should return 400 Bad Request",
			statusCode: 400,
			response: types.ResponseType[any]{
				Success: false,
				Data:    types.ErrorType{Error: "data is required"},
			},
			requestData: []byte(`{"data": ""}`),
		},
		{
			name:       "Should return 422 Unprocessable Entity",
			statusCode: 422,
			response: types.ResponseType[any]{
				Success: false,
				Data:    types.DataType{Link: "http://localhost:3000/abc"},
			},
			requestData: []byte(`gjh`),
		},
		{
			name:       "Should return 200 OK",
			statusCode: 200,
			response: types.ResponseType[any]{
				Success: true,
				Data:    types.DataType{Link: "http://localhost:3000/abc"},
			},
			requestData: []byte(`{"data": "hello world"}`),
		},
	}

	app := app.CreateApp()

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			req := httptest.NewRequest(http.MethodPost, "/links", bytes.NewReader(tc.requestData))
			req.Header.Set("Content-Type", "application/json")
			res, _ := app.Test(req)
			defer res.Body.Close()

			var response types.ResponseType[any]
			body, _ := io.ReadAll(res.Body)

			if err := json.Unmarshal(body, &response); err != nil {
				t.Fatal(err)
			}

			assert.Equal(t, tc.statusCode, res.StatusCode)
			assert.Equal(t, tc.response.Success, response.Success)

		})
	}

}
