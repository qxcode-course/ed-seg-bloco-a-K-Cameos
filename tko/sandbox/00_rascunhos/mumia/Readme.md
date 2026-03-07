---
# Não altere essa chave, ela deve ser única para cada rascunho
key=user_001
---

# mumia

Escreva aqui as informações que você quiser salvar, esse é o seu rascunho.

## Como criar seus próprios testes

Um teste é composto de um `input` (o texto que será fornecido para o programa) e um `output` (o texto que o programa deve retornar para esse input) e pode ter opcionalmente um `label` para facilitar a identificação do teste.

Seus casos de teste personalizados podem ser escritos diretamente aqui na descrição do problema dentro de um fenced code block com a linguagem `toml` ou em um arquivo `tests.toml` na pasta do problema. O TKO irá carregar automaticamente os testes quando a tarefa for aberta ou executada novamente.

Exemplo de teste para ler dois números, um por linha, e imprimir a soma e a subtração deles.

Se quiser habilitar esses casos de teste e ver funcionando, altere o fenced abaixo de `bash` para `toml` e execute novamente a tarefa no TKO.

```toml
[[tests]]
input = '''
leonardo
25
'''
output = '''
leonardo eh adulto
'''

[[tests]]
label = 'mumia'
input = '''
erikson
1056
'''
output = '''
erikson eh mumia
'''
```

