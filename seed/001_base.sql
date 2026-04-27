INSERT INTO destinations (id, name, country_code, city, lat, lng) VALUES
  (1, 'Barcelona Food Escape', 'ES', 'Barcelona', 41.3851, 2.1734),
  (2, 'Tokyo Night Bites', 'JP', 'Tokyo', 35.6762, 139.6503),
  (3, 'New Orleans Flavor Run', 'US', 'New Orleans', 29.9511, -90.0715)
ON CONFLICT (id) DO NOTHING;

INSERT INTO users (id, email, display_name) VALUES
  (1, 'demo@gladiatortravel.app', 'Demo User')
ON CONFLICT (id) DO NOTHING;
