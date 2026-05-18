package handlers

import (
	"atividadeProjeto/app/utils"
	"net/http"
)

func MarcarConsultaHandler(response http.ResponseWriter, request *http.Request) {
	if request.Method != http.MethodPost {
		http.Error(response, "Método não suportado", http.StatusMethodNotAllowed)
		return
	}

	nome := request.FormValue("nome")
	email := request.FormValue("email")
	data := request.FormValue("data")

	err := utils.MarcarConsulta(nome, email, data)
	if err != nil {
		http.Error(response, "Erro ao salvar consulta no banco de dados", http.StatusInternalServerError)
		return
	}

    http.Redirect(response, request, "/index.html", http.StatusSeeOther)
}
