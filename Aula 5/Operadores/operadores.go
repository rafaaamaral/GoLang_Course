package main

func main() {
	//ARITMÉTICOS
	soma := 1 + 2
	subtracao := 1 - 2
	divisao := 10 / 4
	multiplicacao := 10 * 5
	restoDaDivisao := 10 % 2
	println(soma, subtracao, divisao, multiplicacao, restoDaDivisao)

	var numero1 int16 = 10
	var numero2 int16 = 15
	soma2 := numero1 + numero2
	println(soma2)

	//FIM DOS ARITMÉTICOS

	//ATRIBUIÇÃO
	var variavel1 string = "String"
	variavel2 := "String2"
	println(variavel1, variavel2)
	//FIM DOS DE ATRIBUIÇÃO

	//RELACIONAIS
	println(1 > 2)
	println(1 >= 2)
	println(1 < 2)
	println(1 <= 2)
	println(1 != 2)

	//FIM DOS RELACIONAIS

	//LÓGICOS
	println(true && true)
	println(true || false)
	println(!true)
	//FIM DOS LÓGICOS

	//OPERADORES UNARIOS
	numero10 := 10
	numero10++
	println(numero10)
	numero10 += 15
	println(numero10)
	numero10--
	println(numero10)
	numero10 -= 20
	println(numero10)
	numero10 *= 3
	println(numero10)
	numero10 /= 10
	println(numero10)
	numero10 %= 3
	println(numero10)
	//FIM DOS OPERADORES UNARIOS
}
