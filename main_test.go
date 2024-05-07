// package main

// import (
// 	"fmt"
// 	"io"
// 	"net/http"
// 	"net/http/httptest"
// 	"strings"
// 	"testing"
// )

// func TestCreateAccounts(t *testing.T) {
// 	tests := []struct {
// 		name string
// 		body string
// 	}{
// 		{
// 			name: "TestCreateAccounts",
// 			body: `{"account_id": "123", "initial_balance": 100.00}`,
// 		},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			body := strings.NewReader(tt.body)

// 			req := httptest.NewRequest(http.MethodPost, "/accounts", body)
// 			w := httptest.NewRecorder()
// 			CreateAccount(w, req)
// 			res := w.Result()
// 			defer res.Body.Close()

// 			data, err := io.ReadAll(res.Body)
// 			if err != nil {
// 				t.Errorf("expected error to be nil got %v", err)
// 			}

// 			fmt.Println(string(data))
// 			// if string(data) != "ABC" {
// 			// 	t.Errorf("expected ABC got %v", string(data))
// 			// }
// 		})
// 	}
// }
