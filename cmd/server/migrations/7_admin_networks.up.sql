CREATE TABLE IF NOT EXISTS admin_networks (
  id INTEGER PRIMARY KEY AUTOINCREMENT NOT NULL UNIQUE,
  -- it would be the prefix used for dns and company logins
  name VARCHAR(255) NOT NULL UNIQUE,
  -- auth0 organization id
  org_id VARCHAR(255) NOT NULL UNIQUE,
  created_at DATETIME NOT NULL,
  updated_at DATETIME NOT NULL
);

CREATE INDEX idx_admin_networks_id ON admin_networks(org_id);

