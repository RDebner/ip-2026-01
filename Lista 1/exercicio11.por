programa {

  real numero

  funcao inicio() {
    escreva("Número: ")
    leia(numero)

    se(numero % 3 == 0 e numero % 5 == 0) {
      escreva("O NUMERO E DIVISIVEL")
    } senao {
      escreva("O NUMERO NAO E DIVISIVEL")
    }
  }
}
