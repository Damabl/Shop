CREATE TABLE IF NOT EXISTS users (
                                     id SERIAL PRIMARY KEY,
                                     created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                                     updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                                     deleted_at TIMESTAMP,
                                     email TEXT UNIQUE NOT NULL,
                                     name TEXT NOT NULL,
                                     gender TEXT,
                                     age INT DEFAULT 18 CHECK (age >= 18 AND age <= 100),
                                     password TEXT NOT NULL,
                                     role TEXT NOT NULL
);


CREATE TABLE IF NOT EXISTS products (
                                        id SERIAL PRIMARY KEY,
                                        created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                                        updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                                        deleted_at TIMESTAMP,
                                        name TEXT NOT NULL,
                                        description TEXT,
                                        price DOUBLE PRECISION NOT NULL,
                                        image TEXT,
                                        quantity INTEGER NOT NULL
);


CREATE TABLE IF NOT EXISTS carts (
                                     id SERIAL PRIMARY KEY,
                                     created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                                     updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                                     deleted_at TIMESTAMP,
                                     user_id INTEGER REFERENCES users(id) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS cart_items (
                                          id SERIAL PRIMARY KEY,
                                          cart_id INTEGER NOT NULL REFERENCES carts(id) ON DELETE CASCADE,
                                          product_id INTEGER NOT NULL REFERENCES products(id),
                                          quantity INTEGER NOT NULL
);
