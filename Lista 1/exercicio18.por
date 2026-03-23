programa {

  inteiro valorInicial
  inteiro razao
  inteiro qntTermos

  funcao inicio() {
    escreva("Valor inicial, razão e quantidade de termos: ")
    leia(valorInicial, razao, qntTermos)
    inteiro termosPa[qntTermos]
    inteiro soma = valorInicial
    termosPa[0] = valorInicial


    para (inteiro i = 1 ; i < qntTermos; i++) {
      termosPa[i] = termosPa[i - 1] + razao
      soma = soma + termosPa[i]
    }

      escreva(soma)
    
  }
}
