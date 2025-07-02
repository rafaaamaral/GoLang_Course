package enderecos

import "testing"

func TestTipoEndereco(t *testing.T) {
	testCases := []struct {
		endereco string
		tipo     string
	}{
		{"Rua das Flores", "rua"},
		{"Avenida Paulista", "avenida"},
		{"Estrada Real", "estrada"},
		{"Alameda Santos", "alameda"},
		{"Travessa da Liberdade", "travessa"},
		{"Largo do Machado", "largo"},
		{"Vila Madalena", "vila"},
		{"Praça da Sé", "praça"},
		{"Beco do Batman", "tipo inválido"},
	}

	for _, tc := range testCases {
		t.Run(tc.endereco, func(t *testing.T) {
			resultado := TipoEndereco(tc.endereco)
			if resultado != tc.tipo {
				t.Errorf("Esperado %s, mas obteve %s para o endereço %s", tc.tipo, resultado, tc.endereco)
			}
		})
	}
}
