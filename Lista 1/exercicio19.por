programa {

  inclua biblioteca Matematica --> mat

  inteiro numero
  real soma = 1

  funcao inicio() {
    escreva("Número: ")
    leia(numero)

    se (numero <= 1) {
      escreva("Numero invalido!")
    }

    para (inteiro i = 1; i < numero; i++) {
      soma = soma + 1/(i + 1)
    }

    escreva(mat.arredondar(soma, 6))
  }
}
