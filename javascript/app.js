class Aluno{
    constructor(nome, idade, matricula){
        this.nome=nome;
        this.idade = idade;
        this.matricula = matricula;
    }
    getInfo(){
        return `O aluno ${this.nome} têm ${this.idade} anos e possui matrícula ${this.matricula}`
    }
    
}
aluno1 = new Aluno('José',15,23098)
console.log(aluno1.getInfo())