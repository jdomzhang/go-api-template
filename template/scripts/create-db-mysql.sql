-- in bash window, run below command to create linux user
-- mysql -u root -p

-- CREATE DATABASE {{dbname}};
CREATE DATABASE `{{dbname}}` CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;

CREATE USER '{{dbname}}'@'localhost' IDENTIFIED BY '{{dbpassword}}';

FLUSH PRIVILEGES;

GRANT ALL PRIVILEGES ON {{dbname}}.* to {{dbname}}@localhost;

FLUSH PRIVILEGES;

-- test it in bash
-- mysql -u {{dbname}} -p
