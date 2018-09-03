-- in bash window, run below command to create linux user
-- adduser {{name}}

-- su - {{name}}
-- psql

-- then run below cocmmand in psql

create user {{name}} with password '{{name}}';

create database {{name}} with encoding 'UTF-8';

grant all privileges on database {{name}} to {{name}};
