package models

import (
	"firstWebApp/db"
)

type Produto struct {
	Id         int
	Nome       string
	Descricao  string
	Preco      float64
	Quantidade int
}

func BuscaTodosOsProdutos() []Produto {
	db := db.ConnectaComBancoDeDados()
	selectTodosOsProdutos, err := db.Query("select *from produtos order by nome asc")
	if err != nil {
		panic(err.Error())
	}

	p := Produto{}
	produtos := []Produto{}

	for selectTodosOsProdutos.Next() {
		var id, quantidade int
		var nome, descricao string
		var preco float64

		err = selectTodosOsProdutos.Scan(&id, &nome, &descricao, &preco, &quantidade)
		if err != nil {
			panic(err.Error())
		}
		p.Id = id
		p.Nome = nome
		p.Descricao = descricao
		p.Preco = preco
		p.Quantidade = quantidade

		produtos = append(produtos, p)
	}
	defer db.Close()
	return produtos
}

func CriarNovoProduto(nome, descricao string, preco float64, quantidade int) {
	db := db.ConnectaComBancoDeDados()
	insereDadosNoBanco, err := db.Prepare("insert into produtos(nome, descricao, preco, quantidade) values($1, $2, $3, $4)")
	if err != nil {
		panic(err.Error())
	}

	insereDadosNoBanco.Exec(nome, descricao, preco, quantidade)
	defer db.Close()
}

func DeletaProduto(id string) {
	db := db.ConnectaComBancoDeDados()
	deleteDadosNoBanco, err := db.Prepare("delete from produtos where id=$1")
	if err != nil {
		panic(err.Error())
	}
	deleteDadosNoBanco.Exec(id)
	defer db.Close()
}

func EditarProduto(id string) Produto {
	db := db.ConnectaComBancoDeDados()
	selectProduto, err := db.Query("select * from produtos where id=$1", id)
	if err != nil {
		panic(err.Error())
	}
	produto := Produto{}

	for selectProduto.Next() {
		var id, quantidade int
		var nome, descricao string
		var preco float64

		err = selectProduto.Scan(&id, &nome, &descricao, &preco, &quantidade)
		if err != nil {
			panic(err.Error())
		}
		produto.Id = id
		produto.Nome = nome
		produto.Descricao = descricao
		produto.Preco = preco
		produto.Quantidade = quantidade
	}
	return produto
}

func Update(id int, nome, descricao string, preco float64, quantidade int) {
	db := db.ConnectaComBancoDeDados()
	atualizaDadosNoBanco, err := db.Prepare(
		"update produtos set nome=$2, descricao=$3, preco=$4, quantidade=$5 where id=$1")

	if err != nil {
		panic(err.Error())
	}

	atualizaDadosNoBanco.Exec(id, nome, descricao, preco, quantidade)
	defer db.Close()
}
