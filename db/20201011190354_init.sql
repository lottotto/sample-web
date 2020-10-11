-- +goose Up
-- SQL in this section is executed when the migration is applied.
create table employees (
    department varchar,
    "group" varchar,
    name varchar,
    position varchar
);
insert into employees values ('Sales', 'Group 1', 'Alice', 'Assosiate');

-- +goose Down
-- SQL in this section is executed when the migration is rolled back.
drap table employees
