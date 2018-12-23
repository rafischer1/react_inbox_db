-- +goose Up
CREATE TABLE messages (
 id SERIAL PRIMARY KEY,
 read boolean,
 starred boolean,
 selected boolean,
 subject text UNIQUE NOT NULL,
 body text UNIQUE NOT NULL,
 labels text,
);

INSERT INTO messages
  (id, read, starred, selected, subject, body, labels)
VALUES
  (1, false, true, false, 'You cannot input the protocol without calculating the mobile RSS protocol!', 'Hey, this is Virginia Mosby: The littlest thing can cause a ripple effect that changes your life.', '{"dev", "personal"}'),
  (2, false, false, false, 'connecting the system will not do anything, we need to input the mobile AI panel!', 'Hey, it is Robin Scherbatsky here, Look, you cannot design your life like a building. It will not work that way. You just have to live it… and itll design itself.', '{}'),
  (3, false, false, false, 'We need to program the primary TCP hard drive!', 'Hey, Gary Blauman, Definitions are important.', '{}'),
  (4, false, true, false, 'If we override the interface, we can get to the HTTP feed through the virtual EXE interface!', 'Hey, Gary Blauman again: We are going to get older whether we like it or not, so the only question is whether we get on with our lives, or desperately cling to the past.', '{"personal"}'),
  (5, true, false, true, 'We need to back up the wireless GB driver...', 'Hello everyone it is Brad Morris: Because sometimes even if you know how something is gonna end that does not mean you cannot enjoy the ride.', '{}'),
  (6, false, false, false, 'We need to index the mobile PCI bus!', 'Hey, it is Derek Forkanger here, Look, you cannot design your life like a building. It will not work that way. You just have to live it… and itll design itself.', '{"dev", "personal"}'),
  (7, true, true, false, 'If we connect the sensor, we can get to the HDD port through the redundant IB firewall!', 'Hey, Bilson: The future is scary but you cannot just run back to the past because it is familiar.', '{}');

-- +goose Down
DROP TABLE messages;