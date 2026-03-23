programa {

  inteiro numero

  funcao inicio() {
    escreva("Número de 6 até 1999: ")
    leia(numero)
    se(numero < 6 ou numero > 1999) {
      retorne
    }

    para (inteiro i = 2 ; i <= numero; i = i + 2) {
      escreva(i,"^2 = ", i*i,"\n")
    }
  }
}
