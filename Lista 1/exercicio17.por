programa {

  inteiro numPar
  inteiro qntPares

  funcao inicio() {
    escreva("Números: ")
    leia(numPar, qntPares)

    se(numPar % 2 != 0) {
      escreva("O PRIMEIRO NUMERO NAO E PAR")
    } senao {
      para (inteiro i = 1 ; i <= qntPares; i++) {
        escreva(numPar,"\t")
        numPar = numPar + 2
      }
    }
  }
}
