CREATE TABLE IF NOT EXISTS peers (
  id INTEGER PRIMARY KEY AUTOINCREMENT NOT NULL UNIQUE,
  user_id INT NOT NULL,
  setup_key_id INT NOT NULL UNIQUE,
  admin_network_id INT NOT NULL,
  user_group_id INT NOT NULL,
  network_id INT NOT NULL,
  client_pub_key VARCHAR(800) NOT NULL,
  wg_pub_key VARCHAR(800) NOT NULL UNIQUE,
  ip VARCHAR(18) NOT NULL,
  cidr INT NOT NULL,
  created_at DATETIME NOT NULL,
  updated_at DATETIME NOT NULL,

  FOREIGN KEY (user_id) REFERENCES users(id),
  FOREIGN KEY (setup_key_id) REFERENCES setup_keys(id),
  FOREIGN KEY (admin_network_id) REFERENCES admin_networks(id),
  FOREIGN KEY (user_group_id) REFERENCES user_groups(id),
  FOREIGN KEY (network_id) REFERENCES networks(id)
);
