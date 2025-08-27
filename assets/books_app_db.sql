-- 1. Buat database
CREATE DATABASE books_app_db;

-- 2. Buat tabel author
CREATE TABLE authors (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    bio TEXT
);

-- 3. Buat tabel category
CREATE TABLE categories (
    id SERIAL PRIMARY KEY,
    name VARCHAR(50) NOT NULL
);

-- 4. Buat tabel books
CREATE TABLE books (
    id SERIAL PRIMARY KEY,
    title VARCHAR(200) NOT NULL,
    author_id INT REFERENCES authors(id) ON DELETE SET NULL,
    category_id INT REFERENCES categories(id) ON DELETE SET NULL,
    published_year INT,
    price NUMERIC(10,2),
    stock INT DEFAULT 0
);

-- 5. Dummy data untuk authors
INSERT INTO authors (name, bio) VALUES
('J.K. Rowling', 'Author of the Harry Potter series'),
('George R.R. Martin', 'Author of A Song of Ice and Fire'),
('Haruki Murakami', 'Famous Japanese novelist');

-- 6. Dummy data untuk categories
INSERT INTO categories (name) VALUES
('Fantasy'),
('Science Fiction'),
('Literary Fiction'),
('Non-fiction');

-- 7. Dummy data untuk books
INSERT INTO books (title, author_id, category_id, published_year, price, stock) VALUES
('Harry Potter and the Philosopher''s Stone', 1, 1, 1997, 19.99, 10),
('Harry Potter and the Chamber of Secrets', 1, 1, 1998, 21.99, 7),
('A Game of Thrones', 2, 1, 1996, 25.50, 5),
('A Clash of Kings', 2, 1, 1998, 27.00, 4),
('Norwegian Wood', 3, 3, 1987, 15.75, 8),
('Kafka on the Shore', 3, 3, 2002, 18.20, 6),
('The Wind-Up Bird Chronicle', 3, 3, 1994, 20.00, 3);
