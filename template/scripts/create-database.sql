-- in bash window, run below command to create linux user
-- adduser {{dbname}}

-- su - {{dbname}}
-- psql

-- then run below cocmmand in psql

create user {{dbname}} with password '{{dbpassword}}';

-- 支持中文排序
-- CREATE DATABASE {{dbname}} WITH ENCODING 'UTF-8' lc_collate='zh_CN.utf8' template=template0;

create database {{dbname}} with encoding 'UTF-8';

grant all privileges on database {{dbname}} to {{dbname}};
