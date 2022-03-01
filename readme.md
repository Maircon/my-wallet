* Requisitos - Wallet - bills

1. Models

- Despesas - ok
- Wallet (saldo total, salario, emprestimo, etc) - ok
- Usuario - ok

2. O que fazer

- Criar despesa(expenses) e entrada(earnings) - ok
- Quando ocorreu despesa - ok
- Valor da despesa - ok
- Tipo(category) de despesa (Cartao, a vista, emprestimo etc) - ok
- Tag da despesa (home, saude, alimentacao etc) - ok

3. Issues

### Infra
- Create a GO docker image with liveload
- Create SH? methods to use Migration and Seed

### Programming
- Create a method to sum all expenses, earns and total amount
- Create a filter by date, categories, payment types and transaction types
- Create an Authentication with keys on database. user -> key
- Create a middleware and pass who is requesting the action
- Update transactions and total balance
- Validate requests