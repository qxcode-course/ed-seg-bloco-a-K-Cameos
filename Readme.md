[![Open in Codespaces](https://classroom.github.com/assets/launch-codespace-2972f46106e565e64193e422d61a12cf1da4916b45550586e14ef0a7c637dd04.svg)](https://classroom.github.com/open-in-codespaces?assignment_repo_id=22956086)
---
nomeAluno: "Seu nome"
matricula: "Digite sua matrícula aqui por obséquio"
---

# Instalando o TKO e configurando o repositório

## Vou programar no meu setup local

- Se for a primeira configuração
  - Instale git, python, ide(vscode), compiladores
  - Configure sua chave git ssh
  - Configure o pipx
    - `pipx ensurepath`
    - Reinicie o terminal
  - Instale o tko
    - `pipx install tko`

## Vou programar no Codespace

```bash
# instalando o tko e extensões úteis
# Escolha se quer usar python ou typescript
# As outras linguagens c, c++, java, etc, já vem automaticamente na máquina virtual
./setup.sh

```


## Configurando o repositório do TKO

```bash
# Se está com uma instalação antiga do tko,
# por favor atualizar!
pipx upgrade tko

# Criar o repositório de atividades
tko init -f tasks -l py -r poo

# Interagir com seu repositório
tko open tasks
```

## Comandos do bash

- `cd <pasta>` para mudar de pasta
- `ls` para mostrar o conteúdo da pasta
- `cd ..` para subir um nível
- `code <arquivo>` para abrir um arquivo ou pasta pelo terminal
- `control c` para matar um programa do terminal
- `control d` para matar o terminal

## Comandos do Git

- `git add <arquivo>`: Adiciona um arquivo ao stage
- `git add .`: Adiciona todos os arquivos ao stage
- `git commit -m "mensagem"`: Cria um commit com os arquivos no stage
- `git push`: Envia os commits para o repositório remoto
- `git pull`: Atualiza o repositório local com as mudanças do repositório remoto
- `git status`: Mostra o estado atual do repositório
- `git log`: Mostra o histórico de commits
- `git clone <url>`: Clona um repositório remoto para uma pasta local

## Git pela interface do vscode

- Escolha a aba do `Source Control` ou `Control Shift G`
- Clique no + para `stage all changes`, ou seja, marcar tudo pra ser salvo.
- Escreva a mensagem no campo do commit e clique em `Commit` para salvar a versão na máquina virtual.
- Clique em `Sync Changes` para enviar pro site.

## Resolvendo Conflitos

- **Antes de começar a trabalhar**, atualize o repositório:
  - `git pull --ff-only`
  - (isso evita merges automáticos inesperados)

- **Se o push der erro**, o problema normalmente é que há commits novos no repositório remoto.
  - Vamos puxar as atualizações e ver os arquivos que estão conflitando.
    - `git pull --no-rebase`
  - Veja como está seu repositório com `git status`. 
  - Edite os arquivos marcados com ! ou “both modified”, escolhendo o que deve permanecer.
  - Volte para o fluxo original
    - `git add .`
    - `git commit -m mensagem`
    - `git push`
