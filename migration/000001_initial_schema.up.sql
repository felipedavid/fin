CREATE TABLE users
(
    id         SERIAL PRIMARY KEY,
    first_name VARCHAR(50)                               NOT NULL,
    last_name  VARCHAR(50)                               NOT NULL,
    email      VARCHAR(50)                               NOT NULL,
    phone      VARCHAR(50)                               NOT NULL,
    password   VARCHAR(50)                               NOT NULL,
    activated  BOOLEAN     DEFAULT FALSE                 NOT NULL,
    closed     BOOLEAN     DEFAULT FALSE                 NOT NULL,
    closed_at  TIMESTAMPTZ DEFAULT '0001-01-01 00:00:00' NOT NULL,

    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP     NOT NULL,
    updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP     NOT NULL,
    version    INT         DEFAULT 0                     NOT NULL
);

CREATE TYPE asset_type AS ENUM(
    'cash',
    'stock',
    'bond',
    'fund',
    'etf',
    'forex',
    'crypto',
    'other'
);

CREATE TABLE assets
(
    id         SERIAL PRIMARY KEY,
    asset_type asset_type DEFAULT 'CASH' NOT NULL,

    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,
    version    INT         DEFAULT 0                 NOT NULL
);