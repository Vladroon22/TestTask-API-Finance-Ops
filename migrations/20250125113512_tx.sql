-- Active: 1737805021434@@127.0.0.1@5422@postgres
-- +goose Up
-- +goose StatementBegin
CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    name VARCHAR(20) NOT NULL,
    balance DECIMAL(10, 2) NOT NULL DEFAULT 0.00
);

CREATE TABLE tx (
    id SERIAL PRIMARY KEY,
    user_id INT NOT NULL,
    sender_name VARCHAR(20) NOT NULL,
    receiver_name VARCHAR(20) NOT NULL,
    amount DECIMAL(10, 2) NOT NULL,
    type VARCHAR(50) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES users(id)
);

INSERT INTO users (name, balance) VALUES
('Alice', 1000.00),
('Bob', 500.00),
('Charlie', 750.00),
('David', 300.00);

INSERT INTO tx (user_id, sender_name, receiver_name, amount, type, created_at) VALUES
(1, 'Alice', 'Bob', 200.00, 'transfer', '2024-10-01 10:00:00'),
(2, 'Bob', 'Charlie', 100.00, 'transfer', '2021-10-01 11:00:00'),
(3, 'Charlie', 'David', 50.00, 'transfer', '2022-10-01 12:00:00'),
(4, 'David', 'Alice', 150.00, 'transfer', '2020-10-01 13:00:00'),
(1, 'Alice', 'Charlie', 75.00, 'transfer', '2017-10-01 14:00:00');
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE tx;
DROP TABLE users;
-- +goose StatementEnd