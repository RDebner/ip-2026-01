programa {

  inclua biblioteca Matematica --> mat

  real coefA
  real coefB
  real coefC

  funcao inicio() {
    escreva("Valor dos coeficientes A, B e C: ")
    leia(coefA, coefB, coefC)

    real valorDeterminante = coefB*coefB - 4*coefA*coefC
    escreva("O VALOR DE DELTA E = ", mat.arredondar(valorDeterminante, 2),"\n")
  }
}
