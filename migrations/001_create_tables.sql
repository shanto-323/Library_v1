-- +goose Up
CREATE TABLE IF NOT EXISTS genres (
  id SERIAL PRIMARY KEY,
  name VARCHAR(150) NOT NULL UNIQUE
);

CREATE TABLE IF NOT EXISTS author (
  id SERIAL PRIMARY KEY,
  name VARCHAR(150) NOT NULL UNIQUE,
  birth_year INT CHECK (birth_year > 0),
  nationality VARCHAR(100),
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS books (
  id SERIAL PRIMARY KEY,
  isbn VARCHAR(30) NOT NULL UNIQUE,
  title VARCHAR(200) NOT NULL,
  published_year INT NOT NULL,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  total_copies INT NOT NULL CHECK (total_copies >= 0),
  available_copies INT NOT NULL CHECK (available_copies >= 0)
);

CREATE TABLE IF NOT EXISTS students (
  id SERIAL PRIMARY KEY,
  name VARCHAR(100) NOT NULL UNIQUE,
  email VARCHAR(150) UNIQUE,
  phone VARCHAR(15) NOT NULL UNIQUE,
  is_active BOOL DEFAULT TRUE,
  created_at DATE DEFAULT CURRENT_DATE
);

CREATE TABLE IF NOT EXISTS book_genres (
  book_id INT REFERENCES books(id) ON DELETE CASCADE,
  genre_id INT REFERENCES genres(id) ON DELETE CASCADE,
  PRIMARY KEY (book_id, genre_id)
);

CREATE TABLE IF NOT EXISTS book_author (
  book_isbn VARCHAR(30),
  author_id INT,
  PRIMARY KEY (book_isbn, author_id),
  CONSTRAINT fk_book FOREIGN KEY (book_isbn) REFERENCES books(isbn) ON DELETE CASCADE,
  CONSTRAINT fk_author FOREIGN KEY (author_id) REFERENCES author(id) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS borrowed_book (
  id SERIAL PRIMARY KEY,
  student_id INT NOT NULL,
  book_isbn VARCHAR(20) NOT NULL,
  taken_at DATE DEFAULT CURRENT_DATE,
  return_at DATE NOT NULL,
  returned BOOL DEFAULT FALSE,
  CONSTRAINT fk_book FOREIGN KEY (book_isbn) REFERENCES books(isbn) ON DELETE CASCADE,
  CONSTRAINT fk_student FOREIGN KEY (student_id) REFERENCES students(id) ON DELETE CASCADE
);

-- +goose Down
DROP TABLE IF EXISTS borrowed_book;
DROP TABLE IF EXISTS book_author;
DROP TABLE IF EXISTS book_genres;
DROP TABLE IF EXISTS students;
DROP TABLE IF EXISTS books;
DROP TABLE IF EXISTS author;
DROP TABLE IF EXISTS genres;
