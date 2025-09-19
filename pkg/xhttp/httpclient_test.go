package xhttp

import (
	"testing"
)

func TestNewClient(t *testing.T) {
	tests := []struct {
		name string
	}{
		{
			name: "test default",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if NewClient() == nil {
				t.Error("NewClient() got nil")
			}
		})
	}
}

func TestGet(t *testing.T) {
	client := NewClient()
	resp, err := client.Get("https://www.baidu.com")
	if err != nil {
		t.Error("get baidu fail", err)
	}
	t.Log(resp)
}
