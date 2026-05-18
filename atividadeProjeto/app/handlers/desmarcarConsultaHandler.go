package handlers

import (
	"atividadeProjeto/app/utils"
	"net/http"
)

func DesmarcarConsulta(response http.ResponseWriter, request *http.Request) {
	if request.Method != http.MethodPost {
		http.Error(response, "Método não suportado", http.StatusMethodNotAllowed)
		return
	}

	nome := request.FormValue("nome")

	err := utils.DesmarcarConsulta(nome)
	if err != nil {
		http.Error(response, "Erro ao remover consulta do banco de dados", http.StatusInternalServerError)
		return
	}

    http.Redirect(response, request, "/index.html", http.StatusSeeOther)
}
