programa {

  inteiro horas
  inteiro valorPagar

  funcao inicio() {
    escreva("Quantidade de horas: ")
    leia(horas)
    se(horas < 3) {
      valorPagar = horas*5
    } senao
    se ( horas == 3) {
      valorPagar = 10
    } senao
    se (horas > 3) {
      valorPagar = ((horas - horas % 3) /3)*10 + (horas % 3) * 5
    }

    escreva("O VALOR A PAGAR E = ", valorPagar,"\n")
  }
}
