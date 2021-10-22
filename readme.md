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

    | Nome | Required | Type  | Description                                                                          |
    |------|----------|-------|--------------------------------------------------------------------------------------|
    |informacoes  | true     |String |Nome do arquivo|
    |&emsp;id  | true     |String |Nome do arquivo|
    |&emsp;nome_do_arquivo  | true     |String |Nome do arquivo|
    |&emsp;quantidade_redespacho  | true     |String |Nome do arquivo|
    |&emsp;quantidade_ocorrencia  | true     |String |Nome do arquivo|
    |cabecalho  | true     |String |Nome do arquivo|
    |&emsp;identificador  | true     |String |Nome do arquivo|
    |&emsp;remetente  | true     |String |Nome do arquivo|
    |&emsp;destinatario  | true     |String |Nome do arquivo|
    |&emsp;data_criacao  | true     |String |Nome do arquivo|
    |&emsp;complemento  | true     |String |Nome do arquivo|
    |cabecalhoDois  | true     |String |Nome do arquivo|
    |&emsp;identificador  | true     |String |Nome do arquivo|
    |&emsp;identificador_arquivo  | true     |String |Nome do arquivo|
    |&emsp;complemento  | true     |String |Nome do arquivo|
    |transportadora  | true     |String |Nome do arquivo|
    |&emsp;identificador  | true     |String |Nome do arquivo|
    |&emsp;cnpj_transportadora  | true     |String |Nome do arquivo|
    |&emsp;nome_transportadora  | true     |String |Nome do arquivo|
    |&emsp;complemento  | true     |String |Nome do arquivo|
    |&emsp;ocorrencias  | true     |String |Nome do arquivo|
    |&emsp;&emsp;identificador  | true     |String |Nome do arquivo|
    |&emsp;&emsp;nf-e  | true     |String |Nome do arquivo|
    |&emsp;&emsp;&emsp;nfe_cnpj_emitente  | true     |String |Nome do arquivo|
    |&emsp;&emsp;&emsp;nfe_serie  | true     |String |Nome do arquivo|
    |&emsp;&emsp;&emsp;nfe_numero  | true     |String |Nome do arquivo|
    |&emsp;&emsp;codigo_ocorencia  | true     |String |Nome do arquivo|
    |&emsp;&emsp;&emsp;codigo_ocorencia  | true     |String |Nome do arquivo|
    |&emsp;&emsp;&emsp;nome_ocorrencia  | true     |String |Nome do arquivo|
    |&emsp;&emsp;data_ocorencia  | true     |String |Nome do arquivo|
    |&emsp;&emsp;observacao_entrega  | true     |String |Nome do arquivo|
    |&emsp;&emsp;texto  | true     |String |Nome do arquivo|
    |&emsp;&emsp;complemento  | true     |String |Nome do arquivo|
    |&emsp;&emsp;redespacho  | true     |String |Nome do arquivo|
    |&emsp;&emsp;&emsp;identificador  | true     |String |Nome do arquivo|
    |&emsp;&emsp;&emsp;cgc_contratante  | true     |String |Nome do arquivo|
    |&emsp;&emsp;&emsp;transportadora_contratante  | true     |String |Nome do arquivo|
    |&emsp;&emsp;&emsp;quantidade_ocorrencia  | true     |String |Nome do arquivo|
    |&emsp;&emsp;&emsp;cte_serie  | true     |String |Nome do arquivo|
    |&emsp;&emsp;&emsp;cte_numero  | true     |String |Nome do arquivo|
    
  
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
              "quantidade_ocorrencia": 0,
              "cte_serie": 0,
              "cte_numero": 0000
            }
          ]
        }
      ]
    }
  }
```
    

