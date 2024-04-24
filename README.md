# api
# só de rodar a api já sobe o banco de dados
- sudo docker-compose up --build


## Envio de dados para a colmeia
#### Os envios são feitos via metodo HTTP
```bash
/hive/create/{idExterno}
```

```JSON
{
	"moisture": "35",
	"temperature": "22",
	"outsidetemperature": "28",
	"weight": "150"
}



```