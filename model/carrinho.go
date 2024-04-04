package model

type Carrinho struct {
	ID           string
	InfosProduto []string
	Quantidade   int
	IdUsuario    string
	ValorTotal   float64
}
type InfosProduto struct {
	ProdutoId  string
	Quantidade int
}
