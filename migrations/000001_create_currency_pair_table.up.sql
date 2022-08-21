CREATE TABLE currency_pair
(
    currency_from VARCHAR(3)     NOT NULL,
    currency_to   VARCHAR(3)     NOT NULL,
    well          NUMERIC(20, 6) NOT NULL,
    updated_at    TIMESTAMP WITH TIME ZONE,
    PRIMARY KEY (currency_from, currency_to)
);