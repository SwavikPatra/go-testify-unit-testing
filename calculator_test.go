package calculator

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestDivide(t *testing.T) {
	tests := []struct {
		name    string
		a, b    float64
		want    float64
		wantErr bool
	}{
		{
			name:    "valid division",
			a:       10,
			b:       2,
			want:    5,
			wantErr: false,
		},
		{
			name:    "division by zero",
			a:       10,
			b:       0,
			want:    0,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Divide(tt.a, tt.b)
			// if tt.wantErr && err == nil {
			// 	t.Fatalf("unexpected error: %v | Expected: %v, got: %v", err, tt.want, got)
			// }
			// Clean assertions replacing verbose if-statements:
			if tt.wantErr {
				require.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.want, got)
			}
		})
	}
}
