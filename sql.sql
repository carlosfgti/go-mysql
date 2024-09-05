use goproject;

CREATE TABLE users (
    id INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    email VARCHAR(100) NOT NULL,
    age INT NOT NULL
);

INSERT INTO users (name, email, age) VALUES
('John Doe', 'john.doe@example.com', 28),
('Jane Smith', 'jane.smith@example.com', 34),
('Michael Johnson', 'michael.johnson@example.com', 22),
('Emily Davis', 'emily.davis@example.com', 30),
('William Brown', 'william.brown@example.com', 27),
('Olivia Garcia', 'olivia.garcia@example.com', 31),
('James Miller', 'james.miller@example.com', 26),
('Isabella Martinez', 'isabella.martinez@example.com', 29),
('Ethan Anderson', 'ethan.anderson@example.com', 35),
('Sophia Thomas', 'sophia.thomas@example.com', 24);
