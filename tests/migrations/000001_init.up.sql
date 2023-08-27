CREATE TABLE IF NOT EXISTS data
(
    id serial not null unique primary key,
    name varchar(255) not null,
    parameters jsonb not null
);
