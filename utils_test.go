package main

import "testing"

func TestSanitizeWord(t *testing.T) {
	tests := []struct {
		description string
		input       string
		want        string
	}{
		{"Uma pontuação", "Palavra!", "Palavra"},
		{"Pontuação no começo e no fim", "!palavra!", "palavra"},
		{"Em sequência", "!!palavra", "palavra"},
		{"Diferentes pontuações", ";.palavra!", "palavra"},
		{"Diferentes pontuações", ";.palavra\"!", "palavra"},
		{"Sem nenhuma pontuação", "palavra", "palavra"},
	}

	for _, test := range tests {
		got := sanitizeWord(test.input)
		if got != test.want {
			t.Errorf("%s: got %q but wanted %q", test.description, got, test.want)
		}
	}
}

func TestSurroundWord(t *testing.T) {
	tests := []struct {
		description string
		input       string
		want        string
	}{
		{"Virgula", "errou,", "[errou],"},
		{"Ponto final", "errou.", "[errou]."},
		{"Aspas", "errou\"", "[errou]\""},
		{"No início e fim", ",errou.", ",[errou]."},
		{"Repetidos", ",,!errou..", ",,![errou].."},
	}

	for _, test := range tests {
		got := surroundWord(test.input)
		if got != test.want {
			t.Errorf("%s: got %q but wanted %q", test.description, got, test.want)
		}
	}
}
