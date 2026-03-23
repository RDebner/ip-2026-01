programa {

  inteiro n1, n2, n3

  funcao inicio() {
    escreva()
    leia(n1, n2, n3)
    se (n1 == 0 ou n1 > 9 ou n2 > 9 ou n3 > 9) {
      escreva("DÍGITO INVÁLIDO")
      retorne
    }
    inteiro numero = n1*100 + n2*10 + n3
    inteiro quadrado = numero*numero
    escreva(numero , ", ", quadrado)
  }
}
