-- +migrate Up
CREATE TABLE product (
    id bigserial NOT NULL,
    sku varchar NOT NULL,
    slug varchar NOT NULL,
    "name" varchar NOT NULL,
    created_at timestamptz NOT NULL,
    updated_at timestamptz NOT NULL,
    CONSTRAINT product_pk PRIMARY KEY (id),
    CONSTRAINT product_unique UNIQUE (sku),
    CONSTRAINT product_unique_1 UNIQUE (slug)
);

-- +migrate Down