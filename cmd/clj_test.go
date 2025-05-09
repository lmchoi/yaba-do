package cmd

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/spf13/cobra"
)

func TestCreateDirectories(t *testing.T) {
	tests := []struct {
		name        string
		preCreate   bool
		wantOutput  string
	}{
		{
			name:        "successful creation",
			preCreate:   false,
			wantOutput:  "",
		},
		{
			name:        "fails when directory exists",
			preCreate:   true,
			wantOutput:  "aborting: directory",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Setup test environment
			testDir := filepath.Join(os.TempDir(), "testdir")
			defer os.RemoveAll(testDir)

			if tt.preCreate {
				if err := os.Mkdir(testDir, 0755); err != nil {
					t.Fatalf("Test setup failed: %v", err)
				}
			}

			// Redirect stdout
			old := os.Stdout
			defer func() { os.Stdout = old }()
			r, w, _ := os.Pipe()
			os.Stdout = w

			// Execute function
			cmd := &cobra.Command{}
			createDirectories(cmd, []string{testDir})

			// Read output
			w.Close()
			var buf [1024]byte
			n, _ := r.Read(buf[:])
			output := string(buf[:n])

			// Verify output
			if tt.wantOutput != "" && !contains(output, tt.wantOutput) {
				t.Errorf("Expected output containing %q, got %q", tt.wantOutput, output)
			}

			// Verify directories only for successful case
			if tt.wantOutput == "" {
				for _, dir := range []string{"src", "test"} {
					path := filepath.Join(testDir, dir)
					if _, err := os.Stat(path); os.IsNotExist(err) {
						t.Errorf("Directory %s was not created", path)
					}
				}
			}
		})
	}
}

// Helper function to check if string contains substring
func contains(s, substr string) bool {
	for i := 0; i < len(s)-len(substr)+1; i++ {
		if s[i:i+len(substr)] == substr {
			return true
		}
	}
	return false
}