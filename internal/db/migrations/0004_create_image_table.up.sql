CREATE TABLE IF NOT EXISTS images (
                                      id SERIAL PRIMARY KEY,
                                      url TEXT NOT NULL,
                                      public_id TEXT NOT NULL,
                                      product_id INTEGER REFERENCES products(id) ON DELETE CASCADE,
                                      created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                                      updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
