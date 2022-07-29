BEGIN;
CREATE SEQUENCE users_id_seq
INCREMENT 1
START 1;

CREATE TABLE users (
    id INT NOT NULL DEFAULT nextval('users_id_seq') PRIMARY KEY,
    full_name VARCHAR(255) NOT NULL,
    email VARCHAR (255) NOT NULL UNIQUE ,
    password VARCHAR (255) NOT NULL,
    created_on TIMESTAMP NOT NULL DEFAULT now()
);

ALTER SEQUENCE users_id_seq
    OWNED BY users.id;
COMMIT;
