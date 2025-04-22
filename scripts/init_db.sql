CREATE TABLE "users" (
  "id" VARCHAR(250) PRIMARY KEY NOT NULL UNIQUE,
  "name" VARCHAR(250) NOT NULL,
  "email" VARCHAR(250) NOT NULL UNIQUE,
  "username" VARCHAR(250) NOT NULL UNIQUE,
  "password" VARCHAR(250) NOT NULL
);

CREATE TABLE server (
    id VARCHAR(255) PRIMARY KEY,  -- O ID do servidor será uma string e servirá como chave primária
    name VARCHAR(255) NOT NULL    -- O nome do servidor, obrigatório
);

CREATE TABLE channel (
    id VARCHAR(255) PRIMARY KEY,  -- O ID do canal será uma string e servirá como chave primária
    name VARCHAR(255) NOT NULL,   -- O nome do canal, obrigatório
    type VARCHAR(255) NOT NULL,   -- O tipo do canal, obrigatório
    server_id VARCHAR(255),       -- Chave estrangeira para associar o canal a um servidor
    FOREIGN KEY (server_id) REFERENCES server(id) ON DELETE CASCADE
);

