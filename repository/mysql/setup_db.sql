CREATE TABLE users (
    id int primary key AUTO_INCREMENT,
    name varchar(100) not null,
    phone varchar(15) not null unique,
    created_at timestamp CURRENT_TIMESTAMP ,
    updated_at timestamp CURRENT_TIMESTAMP
);