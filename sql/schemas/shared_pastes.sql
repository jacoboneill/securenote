CREATE TABLE shared_pastes ( 
  id INTEGER PRIMARY KEY AUTOINCREMENT,
  paste_id TEXT REFERENCES pastes(id),
  shared_with INTEGER REFERENCES users(id),
  permission TEXT CHECK(permission IN ('read', 'write'))
)
