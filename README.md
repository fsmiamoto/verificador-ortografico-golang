# verificador-ortografico-golang

Solução para o projeto [Verificador Ortográfico](http://wiki.inf.ufpr.br/maziero/doku.php?id=prog2:verificador_ortografico) do Prof. Carlos Maziero (UFPR) em Golang.

Em suma, as palavras não contidas em um dicionário fornecido devem ser marcadas com `[]` na saída.

## Exemplo

Entrada:

```
Para que o pocessador possa interromper a execução de uma tarefa
e retornar a ela mais tarde, sem corromper seu estado interno,
é necessário definir operações para salvar e restaurrar o
contexto da tarefa.

O ato de salvar os valores do contexto atual em seu TCB e
possivelmente restaurar o contexto de outra tarefa, previamente
salvo em outro TCB, é denominado "troca de contexto".
```

Saída esperada:

```
Para que o [pocessador] possa interromper a execução de uma tarefa
e retornar a ela mais tarde, sem corromper seu estado interno,
é necessário definir operações para salvar e [restaurrar] o
contexto da tarefa.

O ato de salvar os valores do contexto atual em seu [TCB] e
possivelmente restaurar o contexto de outra tarefa, previamente
salvo em outro [TCB], é denominado "troca de contexto".
```

## Usando

```bash
# Na raíz do projeto...

# Roda os testes
$ go test

# Mostra os benchmarks
$ go test --bench .

$ go build -o ortografia

$ ./ortografia < testes/plutao.txt > saida.txt

# Se você tem uma arquivo com a saída esperada:
$ diff saida.txt esperado.txt
```
