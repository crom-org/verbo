# 📐 Gramática Formal da Linguagem Verbo

> **Nota de Navegação**: Esta documentação agora faz parte da estrutura modular do GitBook. Você pode consultar a versão atualizada e navegável em [`referencia/gramatica.md`](referencia/gramatica.md).

---

## Produções Principais (EBNF)

```ebnf
programa            = { declaracao_topo } ;

declaracao_topo     = decl_incluir
                    | decl_entidade
                    | decl_funcao
                    | decl_instrucao ;

declaracao_instrucao = decl_variavel
                    | decl_atribuicao
                    | decl_servidor
                    | decl_rota
                    | decl_iniciar_servidor
                    | decl_se
                    | decl_repita
                    | decl_enquanto
                    | decl_simultaneamente
                    | decl_tente
                    | decl_sinalize
                    | decl_enviar_canal
                    | decl_exibir
                    | decl_retorne
                    | decl_expressao ;

bloco               = { declaracao_instrucao } ;
```

Para a especificação completa de todas as produções (variáveis, entidades, funções, concorrência, canais, tratamento de erros e servidor web), acesse [Gramática Formal Completa](referencia/gramatica.md).
