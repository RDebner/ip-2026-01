programa {

  inteiro conta
  inteiro consumoAgua
  caracter tipoConsumidor

  funcao inicio() {
    escreva("Número da conta: ")
    leia(conta)
    escreva("Consumo de água: ")
    leia(consumoAgua)
    escreva("Tipo de consumidor (C, I ou R): ")
    leia(tipoConsumidor)
    escreva("CONTA = ", conta, "\n")
    escreva("VALOR DA CONTA = ", calcularValorConta())
  }

  funcao real calcularValorConta() {

    real valorConta

    se (tipoConsumidor == "C") {
      valorConta = 500 + 0.25*(consumoAgua - 80)
    } senao
    se (tipoConsumidor == "I") {
      valorConta = 800 + 0.04*(consumoAgua - 100)
    } senao
    se (tipoConsumidor == "R") {
      valorConta = 5 + 0.05*consumoAgua
    }

    retorne valorConta
  }
}
