-- migrations/001_create_robots_table.sql
DROP TABLE IF EXISTS robots;

CREATE TABLE robots (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    status VARCHAR(50),
    last_tested TIMESTAMP,
    test_results TEXT[]
);
