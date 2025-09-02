package app

import "testing"

func TestNormalizeCEP(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  string
		ok    bool
	}{
		{"CEP válido com hífen", "01310-000", "01310000", true},
		{"CEP válido puro", "01310000", "01310000", true},
		{"CEP com espaços inválido", "13 100 00", "", false},
		{"CEP alfanumérico", "abc", "", false},
		{"CEP vazio", "", "", false},
		{"CEP com 7 dígitos", "1234567", "", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NormalizeCEP(tt.input)
			if (err == nil) != tt.ok {
				t.Fatalf("expected ok=%v but got err=%v", tt.ok, err)
			}

			if got != tt.want {
				t.Errorf("NormalizeCEP(%q) = %q, want %q", tt.input, got, tt.want)
			}
		})
	}
}
