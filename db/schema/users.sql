CREATE TABLE users(
  id        int BIGSERIAL PRIMARY KEY,
  email     varchar(255) NOT NULL,
  name      varchar(255) NOT NULL,
  password  varchar(255) NOT NULL
);