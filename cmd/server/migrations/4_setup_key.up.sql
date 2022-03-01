CREATE TABLE IF NOT EXISTS setup_keys (
  id INTEGER PRIMARY KEY AUTOINCREMENT NOT NULL UNIQUE,
  user_id INT NOT NULL,
  key VARCHAR(800) NOT NULL UNIQUE,
  key_type INT NOT NULL,
  revoked BOOLEAN NOT NULL,
  created_at DATETIME NOT NULL,
  updated_at DATETIME NOT NULL,

  FOREIGN KEY (user_id) REFERENCES users(id)
);

CREATE INDEX idx_setup_keys_user_id ON setup_keys (user_id);

