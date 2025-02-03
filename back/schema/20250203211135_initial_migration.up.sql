-- *** UP ***
CREATE TABLE containers (
    id SERIAL PRIMARY KEY,
    ip_address VARCHAR(15) NOT NULL,
    ping_time TIMESTAMP NOT NULL,
    last_successful_ping TIMESTAMP
);