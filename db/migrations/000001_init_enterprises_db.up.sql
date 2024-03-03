START TRANSACTION;

CREATE SCHEMA IF NOT EXISTS "enterprises";

CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE
    "enterprises".types_of_ownership
(
    id   uuid NOT NULL DEFAULT (uuid_generate_v4()),
    name text NOT NULL,
    CONSTRAINT pk_types_of_ownership PRIMARY KEY (id)
);

CREATE TABLE
    "enterprises".enterprises
(
    id                   uuid                     NOT NULL DEFAULT (uuid_generate_v4()),
    name                 text                     NOT NULL,
    country              text                     NOT NULL,
    maintenanceYear      integer                  NOT NULL,
    phone                text                     NOT NULL,
    fax                  text                     NOT NULL,
    type_of_ownership_id uuid                     NOT NULL,
    created              timestamp with time zone NOT NULL default (now()),
    updated              timestamp with time zone NULL,
    CONSTRAINT pk_enterprises PRIMARY KEY (id),
    CONSTRAINT fk_enterprises_type_of_ownership_id FOREIGN KEY (type_of_ownership_id) REFERENCES "enterprises".types_of_ownership (id)
);

CREATE UNIQUE INDEX idx_enterprises_id ON "enterprises".enterprises (id);

CREATE UNIQUE INDEX idx_type_of_ownership_id ON "enterprises".types_of_ownership (id);

CREATE UNIQUE INDEX idx_enterprises_type_of_ownership_id ON "enterprises".enterprises (type_of_ownership_id);

COMMIT;