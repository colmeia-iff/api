# api
# só de rodar a api já sobe o banco de dados
- sudo docker-compose up --build


## Envio de dados para a colmeia
#### Os envios são feitos via metodo HTTP
```bash
/hive/{idExterno}
```

```JSON

{
  "Moisture": {
    "Data": {
      "Temp": "valor_da_umidade",
      "Time": "timestamp"
    }
  },
  "OutsideTemperature": {
    "Data": {
      "Temp": "valor_da_temperatura_externa",
      "Time": "timestamp"
    }
  },
  "Temperature": {
    "Data": {
      "Temp": "valor_da_temperatura",
      "Time": "timestamp"
    }
  },
  "Weight": {
    "Value": "valor_do_peso",
    "Time": "timestamp"
  }
}



```