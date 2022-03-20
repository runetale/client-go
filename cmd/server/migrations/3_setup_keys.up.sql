CREATE TABLE IF NOT EXISTS setup_keys (
  id INTEGER PRIMARY KEY AUTOINCREMENT NOT NULL UNIQUE,
  admin_network_id INT NOT NULL,
  user_id INT NOT NULL,
  key VARCHAR(800) NOT NULL UNIQUE,
  key_type INT NOT NULL,
  created_at DATETIME NOT NULL,
  updated_at DATETIME NOT NULL,

  FOREIGN KEY (user_id) REFERENCES users(id)
  FOREIGN KEY (admin_network_id) REFERENCES admin_networks(id)
);

CREATE INDEX idx_setup_keys_user_id ON setup_keys (user_id);

