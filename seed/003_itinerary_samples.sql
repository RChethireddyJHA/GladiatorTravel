INSERT INTO user_preferences (user_id, budget_level, prefers_late_night, avoid_tags) VALUES
  (1, 3, true, 'quiet_only')
ON CONFLICT (user_id) DO UPDATE
SET budget_level = EXCLUDED.budget_level,
    prefers_late_night = EXCLUDED.prefers_late_night,
    avoid_tags = EXCLUDED.avoid_tags;

INSERT INTO user_venue_feedback (user_id, venue_id, feedback) VALUES
  (1, 2, 'up'),
  (1, 6, 'up'),
  (1, 12, 'down');

INSERT INTO trips (id, user_id, destination_id, start_date, end_date) VALUES
  (1, 1, 1, '2026-06-01', '2026-06-04')
ON CONFLICT (id) DO NOTHING;

INSERT INTO itineraries (trip_id, day_number, slot, venue_id, notes) VALUES
  (1, 1, 'lunch', 1, 'seed itinerary sample'),
  (1, 1, 'dinner', 3, 'seed itinerary sample'),
  (1, 1, 'bar', 2, 'seed itinerary sample'),
  (1, 2, 'lunch', 1, 'seed itinerary sample'),
  (1, 2, 'dinner', 3, 'seed itinerary sample'),
  (1, 2, 'bar', 4, 'seed itinerary sample')
ON CONFLICT DO NOTHING;
