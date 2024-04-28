package controllers

import (
	"html/template"
	"net/http"
	"strconv"
	"time"

	models "github.com/Vicrrs/GerenciadorDeEventos/eventos" // Importa o pacote de modelos
)

// Templates compilados para serem reutilizados
var temp = template.Must(template.ParseGlob("templates/*.html"))

// Index exibe a página inicial com todos os eventos
func Index(w http.ResponseWriter, r *http.Request) {
	todosOsEventos := models.BuscaTodosOsEventos()
	temp.ExecuteTemplate(w, "Index", todosOsEventos)
}

// New exibe o formulário para criar um novo evento
func New(w http.ResponseWriter, r *http.Request) {
	temp.ExecuteTemplate(w, "New", nil)
}

// Insert processa o formulário de novo evento e insere os dados no banco
func Insert(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		titulo := r.FormValue("titulo")
		descricao := r.FormValue("descricao")
		data := r.FormValue("data")
		local := r.FormValue("local")

		dataParsed, err := time.Parse("2006-01-02", data) // Formato YYYY-MM-DD
		if err != nil {
			http.Error(w, "Data inválida", http.StatusBadRequest)
			return
		}

		models.CriaNovoEvento(titulo, descricao, local, dataParsed)
		http.Redirect(w, r, "/", http.StatusFound)
	}
}

// Delete remove um evento pelo ID
func Delete(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}

	models.DeletaEvento(idInt)
	http.Redirect(w, r, "/", http.StatusFound)
}

// Edit exibe o formulário de edição para um evento existente
func Edit(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}

	evento := models.EditaEvento(idInt)
	temp.ExecuteTemplate(w, "Edit", evento)
}

// Update processa o formulário de edição e atualiza o evento no banco
func Update(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		id := r.FormValue("id")
		titulo := r.FormValue("titulo")
		descricao := r.FormValue("descricao")
		data := r.FormValue("data")
		local := r.FormValue("local")

		idInt, err := strconv.Atoi(id)
		if err != nil {
			http.Error(w, "ID inválido", http.StatusBadRequest)
			return
		}

		dataParsed, err := time.Parse("2006-01-02", data)
		if err != nil {
			http.Error(w, "Data inválida", http.StatusBadRequest)
			return
		}

		models.AtualizaEvento(idInt, titulo, descricao, local, dataParsed)
		http.Redirect(w, r, "/", http.StatusFound)
	}
}
