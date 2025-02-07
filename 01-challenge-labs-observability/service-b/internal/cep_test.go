package internal

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/jarcoal/httpmock"
)

func TestGetAddressByCep(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("GET", "https://viacep.com.br/ws/83512360/json",
		httpmock.NewStringResponder(200, `{
			"localidade": "Curitiba"
		}`))
	httpmock.RegisterResponder("GET", "https://viacep.com.br/ws/11111111/json",
		httpmock.NewStringResponder(404, `{"error": "CEP not found"}`))

	tests := []struct {
		name           string
		cep            string
		expectedStatus int
		expectedBody   map[string]string
	}{
		{
			name:           "Invalid CEP length",
			cep:            "123",
			expectedStatus: http.StatusUnprocessableEntity,
			expectedBody:   map[string]string{"error": "CEP inválido"},
		},
		{
			name:           "CEP not found",
			cep:            "11111111",
			expectedStatus: http.StatusNotFound,
			expectedBody:   map[string]string{"error": "Endereço não encontrado"},
		},
		{
			name:           "Valid CEP",
			cep:            "83512360",
			expectedStatus: http.StatusOK,
			expectedBody:   map[string]string{"localidade": "Curitiba"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := httptest.NewRecorder()
			ctx := context.Background()
			GetAddressByCep(ctx, w, tt.cep)

			resp := w.Result()
			defer resp.Body.Close()

			if resp.StatusCode != tt.expectedStatus {
				t.Errorf("expected status %d, got %d", tt.expectedStatus, resp.StatusCode)
			}
		})
	}
}
