programa {

  inclua biblioteca Matematica --> mat

  real altura
  real aresta

  funcao inicio() {
    escreva("Altura da pirâmide em metros: ")
    leia(altura)
    escreva("Aresta da base em metros: ")
    leia(aresta)

    real areaBase = 3*aresta*aresta*mat.raiz(3, 2)/2
    real volumePiramide = areaBase*altura/3

    escreva("O VOLUME DA PIRAMIDE E = ", mat.arredondar(volumePiramide, 2)," METROS CUBICOS \n")
  }
}
