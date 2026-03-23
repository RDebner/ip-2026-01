programa {

  inclua biblioteca Matematica --> mat

  real salario
  real salarioReajustado

  funcao inicio() {
    escreva("Salário do funcionário: ")
    leia(salario)

    se(salario <= 300) {
      salarioReajustado = salario * 1.5
    } senao {
      salarioReajustado = salario * 1.3
    }

    escreva("SALARIO COM REAJUSTE = ", mat.arredondar(salarioReajustado, 2),"\n")
  }
}
