programa {

  inclua biblioteca Matematica --> mat

  real salarioMinimo
  real qntKw


  funcao inicio() {
    escreva("Valor do salário mínimo: ")
    leia(salarioMinimo)
    escreva("Quantidade de kW gasta: ")
    leia(qntKw)
    escreva("Custo por kW: R$ ", mat.arredondar(valorKw()/qntKw, 2), "\n")
    escreva("Custo do consumo: R$ ", mat.arredondar(valorKw(), 2), "\n")
    escreva("Custo com desconto: R$ ", mat.arredondar(valorKw()*0.9, 2))
  }

  funcao real valorKw() {
    real valorKw = (0.7*salarioMinimo*qntKw) / 100
    retorne valorKw
  }
}
