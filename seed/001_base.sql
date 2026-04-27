INSERT INTO destinations (id, name, country_code, city, lat, lng) VALUES
  (1, 'Barcelona Food Escape', 'ES', 'Barcelona', 41.3851, 2.1734),
  (2, 'Tokyo Night Bites', 'JP', 'Tokyo', 35.6762, 139.6503),
  (3, 'New Orleans Flavor Run', 'US', 'New Orleans', 29.9511, -90.0715)
ON CONFLICT (id) DO NOTHING;

INSERT INTO users (id, email, display_name) VALUES
  (1, 'demo@gladiatortravel.app', 'Demo User')
ON CONFLICT (id) DO NOTHING;

SELECT setval(pg_get_serial_sequence('destinations', 'id'), COALESCE((SELECT MAX(id) FROM destinations), 1), true);
SELECT setval(pg_get_serial_sequence('users', 'id'), COALESCE((SELECT MAX(id) FROM users), 1), true);
