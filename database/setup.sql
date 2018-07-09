-- create micro-message database
\echo 'Installing micro-message database!'

alter user postgres password 'PG_ROOT_PASSWORD';

create user PG_PRIMARY_USER with REPLICATION  PASSWORD 'PG_PRIMARY_PASSWORD';
create table primarytable (key varchar(20), value varchar(20));
grant all on primarytable to PG_PRIMARY_USER;

CREATE ROLE PG_USER LOGIN ENCRYPTED PASSWORD 'PG_USER'
        NOSUPERUSER INHERIT NOCREATEDB NOCREATEROLE NOREPLICATION;

CREATE DATABASE PG_DATABASE 
    WITH OWNER = postgres 
         ENCODING = 'UTF8' 
         TABLESPACE = pg_default
         LC_COLLATE = 'en_US.UTF-8'
         LC_CTYPE = 'en_US.UTF-8'
         CONNECTION LIMIT = -1
         TEMPLATE template0;

--grant privileges
grant all privileges on database PG_DATABASE to PG_PRIMARY_USER;
ALTER USER PG_PRIMARY_USER CREATEROLE;

--database admin user
\c PG_DATABASE PG_PRIMARY_USER;

GRANT PG_USER TO PG_PRIMARY_USER;
CREATE SCHEMA PG_USER AUTHORIZATION PG_USER;
ALTER ROLE PG_USER SET search_path = PG_USER;

\c PG_DATABASE PG_USER;
create table message (
    id  SERIAL PRIMARY KEY,
    title varchar,
    content text, 
    palindrome boolean
);

create table users (
    id  SERIAL PRIMARY KEY, 
    username varchar, 
    password varchar, 
    email varchar
);

