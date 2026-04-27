INSERT INTO destination_scores (destination_id, food_score, bar_score, combined_score, venue_count) VALUES
  (1, 9.0, 8.4, 8.76, 12),
  (2, 8.7, 9.3, 8.94, 12),
  (3, 8.9, 8.1, 8.58, 12)
ON CONFLICT (destination_id) DO NOTHING;

INSERT INTO venues (destination_id, name, type, rating, price_tier, is_open_late, lat, lng, ambiance) VALUES
  (1, 'La Rambla Kitchen', 'restaurant', 4.8, 3, false, 41.3830, 2.1700, 'lively'),
  (1, 'Midnight Tapas Club', 'bar', 4.7, 2, true, 41.3862, 2.1780, 'group_friendly'),
  (1, 'Catalan Fire Grill', 'restaurant', 4.6, 3, false, 41.3890, 2.1650, 'romantic'),
  (1, 'Neon Vermut House', 'bar', 4.5, 2, true, 41.3842, 2.1810, 'lively'),

  (2, 'Shibuya Smoke House', 'restaurant', 4.8, 3, true, 35.6595, 139.7005, 'lively'),
  (2, 'Sake Alley 24', 'bar', 4.9, 2, true, 35.6580, 139.7020, 'group_friendly'),
  (2, 'Tsukiji Night Omakase', 'restaurant', 4.7, 4, false, 35.6650, 139.7700, 'quiet'),
  (2, 'Lantern Jazz Bar', 'bar', 4.6, 3, true, 35.6715, 139.7650, 'romantic'),

  (3, 'Bayou Spice Table', 'restaurant', 4.7, 2, false, 29.9550, -90.0700, 'lively'),
  (3, 'French Quarter Cocktails', 'bar', 4.8, 3, true, 29.9580, -90.0640, 'group_friendly'),
  (3, 'Creole Ember House', 'restaurant', 4.6, 3, false, 29.9480, -90.0750, 'romantic'),
  (3, 'Late Brass Lounge', 'bar', 4.5, 2, true, 29.9522, -90.0680, 'lively');

INSERT INTO accommodations (destination_id, name, price_tier, lat, lng, vibe) VALUES
  (1, 'Gothic Quarter Suites', 3, 41.3840, 2.1760, 'lively'),
  (1, 'Calm Eixample Stay', 2, 41.3900, 2.1600, 'quiet'),
  (2, 'Shinjuku Neon Hotel', 3, 35.6938, 139.7034, 'lively'),
  (2, 'Asakusa Garden Inn', 2, 35.7148, 139.7967, 'quiet'),
  (3, 'Bourbon Street Loft', 3, 29.9565, -90.0678, 'lively'),
  (3, 'Warehouse District Retreat', 2, 29.9468, -90.0731, 'quiet');

INSERT INTO travel_options (origin_city, destination_id, method, min_minutes, max_minutes, practicality_score) VALUES
  ('London', 1, 'plane', 135, 165, 92),
  ('London', 1, 'train', 780, 900, 35),
  ('Seoul', 2, 'plane', 120, 150, 95),
  ('Seoul', 2, 'boat', 1260, 1440, 10),
  ('Houston', 3, 'plane', 75, 95, 89),
  ('Houston', 3, 'car', 320, 390, 55);
