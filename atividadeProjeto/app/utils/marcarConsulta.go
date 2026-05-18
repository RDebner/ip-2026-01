package utils

import (
	"log"
)

func MarcarConsulta(nome, email, dataConsulta string) error {
	query := `INSERT INTO consultas (nome, email, data_consulta) VALUES ($1, $2, $3)`
	_, err := DB.Exec(query, nome, email, dataConsulta)

	if err != nil {
		log.Printf("Erro ao inserir consulta no banco de dados: %v", err)
		return err
	}
	log.Println("Consulta inserida com sucesso!")
	return nil
}