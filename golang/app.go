package main

import "fmt"

type Aluno struct {
	nome      string
	idade     int
	matricula int
}

// getInfo é um método que retorna uma string formatada
func (c *Aluno) getInfo() string {
	return fmt.Sprintf("O aluno %s tem %d anos e possui matrícula %d", c.nome, c.idade, c.matricula)
}

func main() {
	// Crie uma instância da classe
	aluno1 := Aluno{
		nome:      "José",
		idade:     15,
		matricula: 23098,
	}

	fmt.Println(aluno1.getInfo()) // Use "fmt.Println" para exibir a saída formatada
}
