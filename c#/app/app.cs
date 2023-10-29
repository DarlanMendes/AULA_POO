using System;

class Aluno
{
    public string Nome { get; }
    public int Idade { get; }
    public int Matricula { get; }

    public Aluno(string nome, int idade, int matricula)
    {
        Nome = nome;
        Idade = idade;
        Matricula = matricula;
    }

    public string GetInfo()
    {
        return $"O aluno {Nome} tem {Idade} anos e possui matrícula {Matricula}";
    }
}

class Program
{
    static void Main()
    {
        Aluno aluno1 = new Aluno("José", 15, 23098);
        Console.WriteLine(aluno1.GetInfo());
    }
}
