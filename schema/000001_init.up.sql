CREATE TABLE property
(
    id serial not null unique primary key ,
    name varchar(255) not null
);

CREATE TABLE parameter
(
    id serial not null unique primary key ,
    parameter varchar(255) not null,
    value varchar(255) not null,
    property_id int not null
)

