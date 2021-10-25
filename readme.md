## Break file 
### Descrição
Api destinada a armazenar dados de arquivos OCOREN tipo PROCEDA (versão 3.1). Todas a ocorrências são salvas no banco de dados.

### Contato
savio.olhai.me

### Rota

- **POST** /proceda - Armazena e retorna as informações presentes no arquivo paginada.  
  - Query params
    
    | Nome        | Required | Type              | Description                                    |
    |-------------|----------|-------------------|------------------------------------------------|
    |Content-Type |true      |[application/json] |Tipo do body                                    |
    |perPage      |true      |Numero             |Quantidade de ocorrência que será retornada.    |
    |page         |true      |Numero             |Página atual que será retornada nas ocorrências |

  - Payload
    - Envio:
    
    | Nome | Required | Type  | Description                                                                          |
    |------|----------|-------|--------------------------------------------------------------------------------------|
    |Nome  | true     |String |Nome do arquivo que será interpretado. O mesmo deverá estar na pasta raíz do projeto. |

    ``` json
    {
    	"nome":"OCORENPROCEDA.txt"
    }
    ```

    - Retorno:

    | Nome                                | Required | Type  | Description                                                                          |
    |-------------------------------------|----------|-------|--------------------------------------------------------------------------------------|
    |informacoes                                   | true     |String |Informações gerais do arquivo|
    |&emsp;id                                      | true     |String |identificador unico|
    |&emsp;nome_do_arquivo                         | true     |String |Nome do arquivo|
    |&emsp;quantidade_redespacho                   | true     |String |Quantidade de redespanho presente no arquivo|
    |&emsp;quantidade_ocorrencia                   | true     |String |Quantidade de ocorrências presente no arquivo|
    |cabecalho                                     | true     |String |Primeiro cabelho 3.1|
    |&emsp;identificador                           | true     |String |Identificador proceda|
    |&emsp;remetente                               | true     |String |Transportadora responsável pelo arquivo|
    |&emsp;destinatario                            | true     |String |Embarcador responsável pelo arquivo|
    |&emsp;data_criacao                            | true     |String |Quando foi criado|
    |&emsp;complemento                             | false    |String |Complemento|
    |cabecalhoDois                                 | true     |String |Segundo cabelho|
    |&emsp;identificador                           | true     |String |Identificador proceda 3.1|
    |&emsp;identificador_arquivo                   | true     |String |Identificador do tipo de arquivo|
    |&emsp;complemento                             | false    |String |Complemento|
    |transportadora                                | true     |String |Informação da transportadora|
    |&emsp;identificador                           | true     |String |Identificador proceda 3.1|
    |&emsp;cnpj_transportadora                     | true     |String |CNPJ da transportadora|
    |&emsp;nome_transportadora                     | true     |String |Nome da transportadora|
    |&emsp;complemento                             | true     |String |Complemento|
    |&emsp;ocorrencias                             | true     |String |Lista de ocorrencia|
    |&emsp;&emsp;identificador                     | true     |String |Identificador proceda 3.1|
    |&emsp;&emsp;nf-e                              | true     |String |Informações sobre a nota fiscal|
    |&emsp;&emsp;&emsp;nfe_cnpj_emitente           | true     |String |CNPJ do emitende da nota fiscal|
    |&emsp;&emsp;&emsp;nfe_serie                   | true     |String |Série da nota fiscal|
    |&emsp;&emsp;&emsp;nfe_numero                  | true     |String |Número da nota fiscal|
    |&emsp;&emsp;codigo_ocorencia                  | true     |String |Informações da Ocorrência|
    |&emsp;&emsp;&emsp;codigo_ocorencia            | true     |String |Código da Ocorrência|
    |&emsp;&emsp;&emsp;nome_ocorrencia             | true     |String |Descrição da Ocorrência|
    |&emsp;&emsp;data_ocorencia                    | true     |String |Data que a ocorrência aconteceu|
    |&emsp;&emsp;observacao_entrega                | true     |String |Observação da entrega|
    |&emsp;&emsp;texto                             | true     |String |texto|
    |&emsp;&emsp;complemento                       | true     |String |Complemento de entrega|
    |&emsp;&emsp;redespacho                        | true     |String |Informações do redespacho|
    |&emsp;&emsp;&emsp;identificador               | true     |String |Identificador proceda 3.1|
    |&emsp;&emsp;&emsp;cgc_contratante             | true     |String |CNPJ do contratante|
    |&emsp;&emsp;&emsp;transportadora_contratante  | true     |String |Transportadora contratante|
    |&emsp;&emsp;&emsp;cte_serie                   | true     |String |Série do CT-e|
    |&emsp;&emsp;&emsp;cte_numero                  | true     |String |Número do CT-e|
    
  
  ``` json
  {
    "informacoes": {
      "id": 0,
      "nome_do_arquivo": "OCORENPROCEDA.txt",
      "quantidade_redespacho": 0,
      "quantidade_ocorrencia": 0
    },
    "cabecalho": {
      "identificador": 0,
      "remetente": "                   ",
      "destinatario": " ",
      "data_criacao": "0",
      "complemento": ""
    },
    "cabecalhoDois": {
      "identificador": 340,
      "identificador_arquivo": "OCORR14091",
      "complemento": ""
    },
    "transportadora": {
      "identificador": 341,
      "cnpj_transportadora": "   ",
      "nome_transportadora": "                        ",
      "complemento": "",
      "ocorrencias": [
        {
          "identificador": 342,
          "nf-e": {
            "nfe_cnpj_emitente": "0000000000",
            "nfe_serie": 0,
            "nfe_numero": 0
          },
          "codigo_ocorencia": {
            "codigo_ocorrencia": 11,
            "nome_ocorrencia": "                          "
          },
          "data_ocorencia": "000000",
          "observacao_entrega": 0,
          "texto": "                                                                       ",
          "complemento": "    ",
          "redespacho": [
            {
              "identificador": 343,
              "cgc_contratante": "00000000000",
              "transportadora_contratante": "000",
              "cte_serie": 0,
              "cte_numero": 0000
            }
          ]
        }
      ]
    }
  }
```
    

