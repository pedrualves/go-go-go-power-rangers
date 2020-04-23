package funcoes

import "log"

// essa funcao publica eh responsavel por invocar as funcoes privadas
func FuncaoPublica() { //para declarar uma funcao publica a primeira letra do nome deve ser em Maisculo
	retornoFuncaoPrivada := funcaoPrivada() //recebendo o retorno de um funcao privada
	log.Println(retornoFuncaoPrivada)       //imprimindo o valor recebido

	retorno1, retorno2, retorno3 := funcaoMultiploRetorno()          //recebendo multiplos retornos de uma funcao privada
	log.Println("multiplos retornos ", retorno1, retorno2, retorno3) //imprimindo os retornos de uma funcao privada

	//funcao com spread parameter de entrada o famoso ...
	funcaoSpreadParameter("eu", "sou", "um", "conjuto", "de", "parametros", "de", "entrada")

	_, eai := recursao(10)   //aqui podemos chamar uma funcao com multiplos retornos e ignorar em especificio utilizando _
	log.Println("eai ", eai) //e aqui imprimos somente o necessario

}

func funcaoPrivada() string { //assinatura de uma funcao privada com apenas um elemento de retorno
	quemSouEu := "sou uma funcao privada, pois meu nome comeca com letra minuscula "
	return quemSouEu + "e retorno apenas um elemento, uma string"
}

func funcaoMultiploRetorno() (string, int, bool) { //assinatura de uma funcao privada com multiplos retornos
	log.Println("ja eu sou porreta, tenho retorno multiplo, excelente para callbacks")
	return "abc", 1, true //e os retornos podem ser de tipos diferentes
}

func funcaoSpreadParameter(params ...string) { //spread parameter como input
	for i := 0; i < len(params); i++ {
		log.Print(params[i])
	}
}

func recursao(quantasVezes int) (int, string) { //exemplo tosco de recursao e multiplos retornos
	if quantasVezes == 0 {
		return 0, "chega!"
	}
	quantasVezes = quantasVezes - 1
	return recursao(quantasVezes)
}

func retornarFunction() func() { //funcao que retorna outra funcao
	return retornarFunction()
}
