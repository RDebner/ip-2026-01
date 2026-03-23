programa {

  inclua biblioteca Matematica --> mat

  real a 
  real b 
  real c 
  real d 

  funcao inicio() {
    escreva("Valor dos elementos a,b,c e d da matriz quadrada: ")
    leia(a,b,c,d)

    real determinante = a*d - b*c 
    escreva("O VALOR DO DETERMINANTE E = ", mat.arredondar(determinante, 2), "\n")
  }
}
