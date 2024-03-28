-- +migrate Up
CREATE TABLE stock (
    id bigserial NOT NULL,
    product_id int8 NOT NULL,
    stock int4 NOT NULL,
    created_at timestamptz NOT NULL,
    updated_at timestamptz NOT NULL,
    CONSTRAINT stock_pk PRIMARY KEY (id),
    CONSTRAINT stock_product_fk FOREIGN KEY (product_id) REFERENCES product(id) ON DELETE CASCADE
);

-- +migrate Down