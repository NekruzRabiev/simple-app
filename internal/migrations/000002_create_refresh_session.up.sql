BEGIN;

CREATE TABLE refresh_sessions
(
    id         serial                    not null unique,
    token      varchar(128)              not null,
    expires_at timestamp                 not null,
    user_id    int references users (id) not null
);

COMMIT;