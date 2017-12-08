CREATE ROLE james LOGIN CREATEDB CREATEROLE PASSWORD 'SOMESECRETPASSWORD';
CREATE DATABASE hydro_routing OWNER james

-- \q will log you out if you are using psql