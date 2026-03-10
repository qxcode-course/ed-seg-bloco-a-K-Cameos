---
key: user_003
---

# pedra

Escreva aqui as informações que você quer salvar, esse é o seu rascunho.
O texto abaixo é informativo e você pode apagar depois de aprender como usar os rascunhos.

## Como usar os rascunhos

- Para renomear seu rascunho, basta renomear a pasta do rascunho.
- Para reorganizar, você pode criar subpastas dentro do sandbox e mover as pastas de rascunhos pra lá.
- Você pode usar o atalho R(Reload) no TKO para recarregar os rascunhos depois de criar, renomear ou reorganizar eles.
- Tudo que você fizer nos rascunhos também será rastreado pela ferramenta.

## Como criar seus próprios testes

Um teste é composto de um `input` (o texto que será fornecido para o programa) e um `output` (o texto que o programa deve retornar para esse input) e pode ter opcionalmente um `label` para facilitar a identificação do teste.

Seus casos de teste personalizados podem ser escritos diretamente aqui na descrição do problema dentro de um fenced code block com a linguagem `toml` ou em um arquivo `tests.toml` na pasta do problema. O TKO irá carregar automaticamente os testes quando a tarefa for aberta ou executada novamente.

Exemplo de teste para ler dois números, um por linha, e imprimir a soma e a subtração deles.

Se quiser habilitar esses casos de teste e ver funcionando, altere o fenced abaixo de `bash` para `toml` e execute novamente a tarefa no TKO.

```toml
[[tests]]
input = '''
2
8 11
10 15
'''
output = '''
1
'''

[[tests]]
label = 'números negativos'
input = '''
-3
-4
'''
output = '''
-7
1
'''
```

