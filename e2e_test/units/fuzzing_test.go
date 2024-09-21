package units

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func FuzzCalculateHighestHandler(f *testing.F) {
	server := httptest.NewServer(http.HandlerFunc(CalculateHighestHandler))
	defer server.Close()

	testCases := []ValuesRequest{
		{[]int{1, 2, 3, 4, 5, 6, 6, 7, 8, 9, 10}},
		{[]int{-1, -2, -3, -4, 5, -6, -6, -7, -8, -9, -10}},
		{[]int{1, 2, 30, 4, 50, 6, 6, 707, 8, 19, 10}},
		{[]int{10, 20, 30, 40, 50, 60, 60, 70, 80, 90, 100}},
	}

	for _, tc := range testCases {
		data, _ := json.Marshal(tc)
		f.Add(data)
	}

	f.Fuzz(func(t *testing.T, data []byte) {
		resp, err := http.DefaultClient.Post(server.URL, "application/json", bytes.NewBuffer(data))
		if err != nil {
			t.Errorf("error reaching http API: %v", err)
		}
		if resp.StatusCode != http.StatusOK {
			t.Errorf("Expected status code %v , got %v", http.StatusOK, resp.StatusCode)
		}
		var response int
		if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
			t.Fatal(err)
		}
	})
}
