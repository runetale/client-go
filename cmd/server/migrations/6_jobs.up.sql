CREATE TABLE IF NOT EXISTS jobs (
  id INTEGER PRIMARY KEY AUTOINCREMENT NOT NULL UNIQUE,
  admin_network_id INT NOT NULL,
  name VARCHAR(255) NOT NULL,
  created_at DATETIME NOT NULL,
  updated_at DATETIME NOT NULL,

  FOREIGN KEY (admin_network_id) REFERENCES admin_networks(id)
);

CREATE INDEX idx_jobs_name ON jobs(name);
CREATE INDEX idx_jobs_admin_network_id ON jobs(admin_network_id);
