

![enter image description here](https://blog.guru.com.vc/wp-content/uploads/2020/05/logo-png.png)

# Teste para Dev Backend

## O que precisa ser feito?

**Você precisará criar uma API que mostra todos os ativos que compõem o Índice Bovespa, ordenando pela maior variação diária para a menor. Independente se a variação for de alta ou de baixa.**

### Requisitos:

 - A API precisa ser escrito em Go
 -  A API precisa ser RESTFUL
 - Todos os ativos que compõem o IBOV precisam ser exibidos.
 - A API deverá ser entregue com um Dockerfile que componha todo o ambiente que você precisa.

### Desejável:
  - Usar algum In-Memory cache
  - Usar melhores práticas de arquitetura e organização de código
  - Todo o processamento deverá ser feito no menor tempo possível

### Payload:
o payload de retorno deverá conter as seguintes informações:

- Ticker do ativo
- Nome da empresa
- Preço de Fechamento
- Preço de Abertura
- Preço atual
- Variação atual em R$
- Variação atual em %

### Entrega
A solução de código deverá ser entregue como um PR nesse repositório. O PR também deve conter suas informações, tais como:
- Nome
- Idade
- Cidade onde mora
- Estado
- Tempo de experiência como desenvolvedor(a)
