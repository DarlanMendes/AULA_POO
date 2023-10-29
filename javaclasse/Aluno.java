package javaclasse;

public class Aluno{
    public String nome;
    public int idade;
    public int matricula;

    public Aluno(String nome, int idade, int matricula){
        this.nome= nome;
        this.idade=idade;
        this.matricula=matricula;
    }
    public String getInfo(){
        return String.format("O Aluno %s têm %d anos e matrícula %d",this.nome, this.idade, this.matricula);
    }
}
