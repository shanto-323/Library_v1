-- +goose Up

-- Insert genres
INSERT INTO genres (id, name) VALUES
(1, 'Science Fiction'),
(2, 'Fantasy'),
(3, 'Mystery'),
(4, 'Thriller'),
(5, 'Romance'),
(6, 'Historical Fiction'),
(7, 'Non-Fiction'),
(8, 'Biography'),
(9, 'Self-Help'),
(10, 'Horror');

-- Insert Authors
INSERT INTO authors (id , name, birth_year, nationality, created_at) VALUES
(1 ,'George Orwell', 1903, 'British', NOW()),
(2,'J.K. Rowling', 1965, 'British', NOW()),
(3,'Stephen King', 1947, 'American', NOW()),
(4,'Agatha Christie', 1890, 'British', NOW()),
(5,'J.R.R. Tolkien', 1892, 'British', NOW());

-- Insert Books
INSERT INTO books (id,isbn, title, published_year, total_copies, available_copies, created_at) VALUES
(1,'ISBN0000000001', '1984', 1949, 10, 10, NOW()),
(2,'ISBN0000000002', 'Harry Potter and the Philosopher''s Stone', 1997, 15, 15, NOW()),
(3,'ISBN0000000003', 'The Shining', 1977, 8, 8, NOW()),
(4,'ISBN0000000004', 'Murder on the Orient Express', 1934, 12, 12, NOW()),
(5,'ISBN0000000005', 'The Hobbit', 1937, 20, 20, NOW());

-- Link Books to Authors
INSERT INTO book_authors (book_id, author_id) VALUES
(1, 1), -- 1984 -> George Orwell
(2, 2), -- Harry Potter -> J.K. Rowling
(3, 3), -- The Shining -> Stephen King
(4, 4), -- Murder on the Orient Express -> Agatha Christie
(4, 1),
(5, 5); -- The Hobbit -> J.R.R. Tolkien

 -- Insert book_genres (AFTER books and genres are inserted)
INSERT INTO book_genres (book_id, genre_id) VALUES
(1, 1), -- 1984 -> Science Fiction
(2, 2), -- Harry Potter -> Fantasy
(3, 3), -- The Shining -> Mystery
(4, 3), -- Murder on the Orient Express -> Mystery
(5, 2); -- The Hobbit -> Fantasy

-- Insert Students
INSERT INTO students (id,name, email, phone, is_active, created_at) VALUES
(1,'Alice Johnson', 'alice.johnson@example.com', '123-456-7890', true, NOW()),
(2,'Bob Smith', 'bob.smith@example.com', '234-567-8901', true, NOW()),
(3,'Charlie Brown', 'charlie.brown@example.com', '345-678-9012', true, NOW());

-- Insert Borrowed Books
INSERT INTO borrowed_books (id, student_id, book_isbn, taken_at, return_at, returned) VALUES
(1, 1, 'ISBN0000000001', NOW() - INTERVAL '10 days', NOW() + INTERVAL '14 days', false),
(2, 2, 'ISBN0000000002', NOW() - INTERVAL '5 days', NOW() + INTERVAL '14 days', false),
(3, 3, 'ISBN0000000003', NOW() - INTERVAL '2 days', NOW() + INTERVAL '14 days', false);

-- +goose Down
TRUNCATE TABLE borrowed_books CASCADE;
TRUNCATE TABLE students CASCADE;
TRUNCATE TABLE book_genres CASCADE;
TRUNCATE TABLE books CASCADE;
TRUNCATE TABLE authors CASCADE;
TRUNCATE TABLE genres CASCADE;
