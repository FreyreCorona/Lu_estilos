CREATE TABLE IF NOT EXISTS products (
  id BIGSERIAL PRIMARY KEY,
  name TEXT NOT NULL,
  description TEXT,
  bar_code TEXT,
  category TEXT,
  initial_stock INTEGER NOT NULL DEFAULT 0,
  actual_stock INTEGER NOT NULL DEFAULT 0,
  price NUMERIC(10,2) NOT NULL,
  due_date TIMESTAMP DEFAULT NOW()
);
