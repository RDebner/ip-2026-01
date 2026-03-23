programa {

  inclua biblioteca Matematica --> mat

  real pi = 3.14159
  real raio
  real altura

  funcao inicio() {
    escreva("Raio da lata em metros: ")
    leia(raio)
    escreva("Altura da lata em metros: ")
    leia(altura)
    real areaTotal = 2*(pi*raio*raio) + (2*pi*raio*altura)
    real custoLata = areaTotal*100
    escreva("O VALOR DO CUSTO E = ", mat.arredondar(custoLata, 2),"\n")
  }
}
