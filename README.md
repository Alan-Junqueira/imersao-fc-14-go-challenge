Rodei usando o dev container, como não tenho go instalado na minha máquina, e não consegui adicionar ele no meu wsl.

Ao rodar o container, é necessário criar a tabela no banco de dados sql.

Na sua máquina, e não dentro do container gerado pelo dev container, siga os seguintes passos.

```
docker exec -it <id-do-seu-container> bash

mysql -uroot -proot

USE routes;

CREATE TABLE routes (
    id INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(255),
    source_lat FLOAT,
    source_lng FLOAT,
    dest_lat FLOAT,
    dest_lng FLOAT
);
```

Com o banco de dados configurado, navegue ate a pasta src, e rode o projeto.

```
go run main.go
```

Na raiz do projeto, tem o api.http, já está configurado para fazer as chamadas.