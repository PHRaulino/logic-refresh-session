# Fila de refresh do token

## Variaveis de ambiente

```shell
TEMPO_DE_SESSAO=15 // 15 minutos
INTERVALO_DE_REFRESH=4 // 4 minutos
```

## Regras Gerais

1. O token do usuário dura 5 minutos
1. A sessão do usuário deve durar {TEMPO_DE_SESSAO}
1. A cada {INTERVALO_DE_REFRESH} o token deve ser renovado
1. A data limite da sessão do usuário deve ser a ultima interação do usuário + {TEMPO_DE_SESSAO}
1. O token deve ser renovado até que a data limite da sessão acabe
1. O processo de renovação acontece através de uma fila
1. Cada vez que o usuário interagir com o sistema os steps de renovação do token deve ir para a fila considerando os ultimos enviados

## GetRefreshSteps

### Objetivo principal da função

Retornar as datas dos steps de refresh de um token

exemplo
1. Usuario logou as 00:00
1. O token foi criado as 00:00
1. O usuario deve ter a sessão até 00:15
1. A fila deve ter 3 step de renovação ["00:04", "00:08", "00:12"]

### Passos da function

1. Recebe o horario de criação do token
1. Recebe o horario de expiração da sessão
1. Verifica se o horario de criação é menor do que o de expiração
1. Pega a data de expiração da sessão
1. Remove 5 minutos
1. Subtrai a data de criação do token
1. Pega a diferença em minutos e arredonda pra cima
1. Divide o valor pelo {INTERVALO_DE_REFRESH} e arredonda pra cima
1. Retorna quantos steps de refresh serão necessários
1. Cria um range com o numero de steps necessários
1. Cria um array com as datas (data de criação do token + (item do range * {INTERVALO_DE_REFRESH}))
