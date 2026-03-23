programa {

  real nota
  cadeia conceito

  funcao inicio() {
    escreva("Nota: ")
    leia(nota)

    se (nota <=10 e nota >= 9) {
      conceito = "A"
    } senao
    se (nota < 9 e nota >= 7.5) {
      conceito = "B"
    } senao
    se (nota < 7.5 e nota >= 6) {
      conceito = "C"
    } senao {
      conceito = "D"
    }

    escreva("NOTA = ", nota, " CONCEITO = ", conceito,"\n")
  }
}
