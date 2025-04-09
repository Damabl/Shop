ALTER TABLE products
    ADD COLUMN brand TEXT,
    ADD COLUMN category TEXT,
    ADD COLUMN size TEXT,
    ADD COLUMN color TEXT,
    ADD COLUMN material TEXT,
    ADD COLUMN gender TEXT,
    ADD COLUMN season TEXT,
    ADD COLUMN discount DOUBLE PRECISION DEFAULT 0.0;
