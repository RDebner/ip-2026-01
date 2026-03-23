programa {

  inteiro hrs
  inteiro min
  inteiro seg
  inteiro totalSeg

  funcao inicio() {
    escreva("Horas, minutos e segundos: ")
    leia(hrs, min, seg)

    totalSeg = hrs*3600 + min*60 + seg

    escreva("O TEMPO EM SEGUNDOS E = ", totalSeg,"\n")

  }
}
