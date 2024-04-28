package routes

import (
	"net/http"

	"github.com/Vicrrs/GerenciadorDeEventos/controllers"
)

// CarregaRotas associa as URLs aos handlers específicos dos controladores
func CarregaRotas() {
	http.HandleFunc("/", controllers.Index)        // Rota para a página inicial que lista todos os eventos
	http.HandleFunc("/new", controllers.New)       // Rota para mostrar o formulário de criação de um novo evento
	http.HandleFunc("/insert", controllers.Insert) // Rota para receber os dados do formulário e inserir um novo evento
	http.HandleFunc("/delete", controllers.Delete) // Rota para deletar um evento
	http.HandleFunc("/edit", controllers.Edit)     // Rota para mostrar o formulário de edição de um evento
	http.HandleFunc("/update", controllers.Update) // Rota para atualizar os dados de um evento no banco de dados
}
