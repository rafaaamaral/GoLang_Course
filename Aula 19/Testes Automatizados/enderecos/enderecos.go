package enderecos

import "strings"

// TipoEndereco recebe um endereço e retorna o tipo do endereço (rua, avenida, etc.)
// Se o tipo não for válido, retorna "tipo inválido".
func TipoEndereco(endereco string) string {
	tipoValidos := []string{"rua", "avenida", "estrada", "alameda", "travessa", "largo", "vila", "praça"}

	enderecoComLetraMinuscula := strings.ToLower(endereco)
	primeiraPalavraDoEndereco := strings.Split(enderecoComLetraMinuscula, " ")[0]

	enderecoTemUmTipoValido := false
	for _, tipo := range tipoValidos {
		if tipo == primeiraPalavraDoEndereco {
			enderecoTemUmTipoValido = true
			break
		}
	}

	if enderecoTemUmTipoValido {
		return primeiraPalavraDoEndereco
	} else {
		return "tipo inválido"
	}
}
