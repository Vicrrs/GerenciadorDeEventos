package models

import (
	"time"

	db "github.com/Vicrrs/GerenciadorDeEventos/db" // Importa o pacote db que criamos anteriormente
)

type Evento struct {
	Id        int       // Identificador do evento
	Titulo    string    // Título do evento
	Descricao string    // Descrição do evento
	Data      time.Time // Data do evento
	Local     string    // Local do evento
}

// BuscaTodosOsEventos retorna todos os eventos armazenados no banco de dados
func BuscaTodosOsEventos() []Evento {
	db := db.ConectaComBancoDeDados() // Abre conexão com o banco de dados
	defer db.Close()                  // Fecha a conexão ao final da função

	query := "SELECT id, titulo, descricao, data, local FROM eventos"
	selecionaTodosOsEventos, err := db.Query(query)
	if err != nil {
		panic(err.Error())
	}
	defer selecionaTodosOsEventos.Close()

	var eventos []Evento
	for selecionaTodosOsEventos.Next() {
		var e Evento
		err := selecionaTodosOsEventos.Scan(&e.Id, &e.Titulo, &e.Descricao, &e.Data, &e.Local)
		if err != nil {
			panic(err.Error())
		}
		eventos = append(eventos, e)
	}
	return eventos
}

// CriaNovoEvento insere um novo evento no banco de dados
func CriaNovoEvento(titulo, descricao, local string, data time.Time) {
	db := db.ConectaComBancoDeDados() // Conecta com o banco
	defer db.Close()

	// Prepara o SQL para inserção de dados
	insereDadosNoBanco, err := db.Prepare("INSERT INTO eventos (titulo, descricao, data, local) VALUES ($1, $2, $3, $4)")
	if err != nil {
		panic(err.Error())
	}
	defer insereDadosNoBanco.Close()
	// Executa a inserção
	insereDadosNoBanco.Exec(titulo, descricao, data, local)
}

// DeletaEvento remove um evento pelo ID
func DeletaEvento(id int) {
	db := db.ConectaComBancoDeDados()
	defer db.Close()

	deletarEvento, err := db.Prepare("DELETE FROM eventos WHERE id=$1")
	if err != nil {
		panic(err.Error())
	}
	deletarEvento.Exec(id)
}

// EditaEvento recupera um evento pelo ID para edição
func EditaEvento(id int) Evento {
	db := db.ConectaComBancoDeDados()
	defer db.Close()

	query := "SELECT id, titulo, descricao, data, local FROM eventos WHERE id=$1"
	eventoDoBanco, err := db.Query(query, id)
	if err != nil {
		panic(err.Error())
	}
	defer eventoDoBanco.Close()

	var eventoParaAtualizar Evento
	if eventoDoBanco.Next() {
		err := eventoDoBanco.Scan(&eventoParaAtualizar.Id, &eventoParaAtualizar.Titulo, &eventoParaAtualizar.Descricao, &eventoParaAtualizar.Data, &eventoParaAtualizar.Local)
		if err != nil {
			panic(err.Error())
		}
	}
	return eventoParaAtualizar
}

// AtualizaEvento atualiza um evento existente no banco de dados
func AtualizaEvento(id int, titulo, descricao, local string, data time.Time) {
	db := db.ConectaComBancoDeDados()
	defer db.Close()

	// Prepara o SQL para atualização de dados
	atualizaEvento, err := db.Prepare("UPDATE eventos SET titulo=$1, descricao=$2, data=$3, local=$4 WHERE id=$5")
	if err != nil {
		panic(err.Error())
	}
	atualizaEvento.Exec(titulo, descricao, data, local, id)
}
