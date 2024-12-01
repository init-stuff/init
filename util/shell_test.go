package util

import (
	"os/exec"
	"strings"
	"testing"
)

func init() {
	cmd := exec.Command("py", "--version")
	if err := cmd.Run(); err != nil {
		panic("Python3 is required to run these tests")
	}
}

func TestRunPythonCode(t *testing.T) {
	tests := []struct {
		name     string
		code     string
		want     string
		wantErr  bool
		errMatch string
	}{
		{
			name: "simple print",
			code: `print("Hello, World!")`,
			want: "Hello, World!",
		},
		{
			name: "multiple lines",
			code: `
x = 5
y = 3
print(f"{x} + {y} = {x+y}")`,
			want: "5 + 3 = 8",
		},
		{
			name: "syntax error",
			code: `print("Unclosed string`,
			wantErr: true,
			errMatch: "SyntaxError",
		},
		{
			name: "runtime error",
			code: `
x = 1/0
print(x)`,
			wantErr: true,
			errMatch: "ZeroDivisionError",
		},
		{
			name: "multiple print statements",
			code: `
print("Line 1")
print("Line 2")`,
			want: "Line 1\r\nLine 2",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := RunPythonCode(tt.code)

			// Check error cases
			if tt.wantErr {
				if err == nil {
					t.Errorf("RunPythonCode() expected error but got none")
					return
				}
				if tt.errMatch != "" && !strings.Contains(err.Error(), tt.errMatch) {
					t.Errorf("RunPythonCode() error = %v, should contain %v", err, tt.errMatch)
				}
				return
			}

			// Check non-error cases
			if err != nil {
				t.Errorf("RunPythonCode() unexpected error: %v", err)
				return
			}

			if got != tt.want {
				t.Errorf("RunPythonCode() = %q, want %q", got, tt.want)
			}
		})
	}
} 