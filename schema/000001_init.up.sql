CREATE TABLE IF NOT EXISTS books(
    id serial not null,
    title varchar not null,
    author varchar not null,
    pages int not null
);