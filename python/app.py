class Aluno:
    def __init__(self, nome, idade, matricula):
        self.nome=nome;
        self.idade=idade;
        self.matricula = matricula;
    def get_info(self):
        return f"O aluno {self.nome} têm {self.idade} anos e possui matrícula {self.matricula}"

aluno1 = Aluno('José', 15, 23098);
print(aluno1.get_info())
