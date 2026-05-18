package handlers

import (
	"atividadeProjeto/app/utils"
	"net/http"
	"text/template"
)

func VerConsultaHandler(response http.ResponseWriter, request *http.Request) {
	if request.Method != http.MethodPost {
		http.Error(response, "Método não suportado", http.StatusMethodNotAllowed)
		return
	}

	email := request.FormValue("email")

	user, err := utils.VerConsulta(email)
	if err != nil {
		// Retorna um erro caso ocorra falha ao buscar as informações do usuário
		http.Error(response, "Erro ao buscar informações do usuário", http.StatusInternalServerError)
		return
	}

	tmpl, err := template.ParseFiles("static/consultas.html")
	if err != nil {
		http.Error(response, "Erro ao carregar o template", http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(response, user)
	if err != nil {
        // Retorna um erro caso ocorra falha ao renderizar o template
        http.Error(response, "Erro ao renderizar o template", http.StatusInternalServerError)
        return
    }

    // Redireciona o usuário para a página de perfil após o sucesso
    http.Redirect(response, request, "/consultas.html", http.StatusOK)
}
