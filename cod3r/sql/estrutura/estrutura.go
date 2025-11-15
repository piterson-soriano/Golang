package main

import "database/sql" // Importa a biblioteca padrão para trabalhar com banco de dados em Go

/*
go get -u github.com/go-sql-driver/mysql
Esse comando deve ser rodado no terminal para instalar o driver MySQL,
que permite que o Go se conecte a bancos de dados MySQL.
*/

// Função que executa um comando SQL no banco de dados
func exec(db *sql.DB, sql string) sql.Result {
	result, err := db.Exec(sql) // Executa o comando SQL passado como parâmetro
	if err != nil {             // Se ocorrer algum erro...
		panic(err) // ...o programa para e mostra o erro
	}
	return result // Retorna o resultado da execução
}

func main() {
	// Abre uma conexão com o banco de dados MySQL
	// "root:12345@/" significa usuário root, senha 12345 e sem banco específico
	db, err := sql.Open("mysql", "root:teste")
	if err != nil { // Se não conseguir conectar...
		panic(err) // ...o programa para e mostra o erro
	}

	defer db.Close() // Garante que a conexão será fechada ao final do programa

	// Executa um comando SQL para criar um banco de dados se ele não existir
	// OBS: o correto seria "CREATE DATABASE IF NOT EXISTS nome_do_banco"
	exec(db, "create database if not exist cursogo")

	exec(db, "use cursogo")
	exec(db, "drop table if exists usuarios")
	exec(db, `create table usuarios(
	id integer auto_increment, 
	nome varchar(80), 
	PRIMARY KEY (id)
	)`)
}
