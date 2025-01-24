package app

import (
	"testing"
)

func TestNewApp(t *testing.T) {
	tests := []struct {
		name    string
		wantErr bool
	}{
		{
			name:    "successful initialization",
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a, err := New()
			if (err != nil) != tt.wantErr {
				t.Errorf("New() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if a == nil && !tt.wantErr {
				t.Error("Expected non-nil App")
				return
			}
			if a != nil {
				if err := a.Close(); err != nil {
					t.Errorf("Failed to close App: %v", err)
				}
			}
		})
	}
}

func TestApp_Close(t *testing.T) {
	a, err := New()
	if err != nil {
		t.Fatalf("Failed to create App: %v", err)
	}

	// Test multiple close calls
	for i := 0; i < 2; i++ {
		if err := a.Close(); err != nil {
			t.Errorf("Close() error = %v", err)
		}
	}
}
