DROP TABLE IF EXISTS event;
CREATE TABLE event (
    id INTEGER PRIMARY KEY,
    title TEXT NOT NULL,
    content TEXT,
    costs TEXT,
    times TEXT,
    location TEXT,
    owner TEXT,
    tags TEXT,
    ticketURLS TEXT,
    restrictions TEXT
);

DROP TABLE IF EXISTS cost_type;
CREATE TABLE cost_type (
    name TEXT,
    display_name TEXT
);

INSERT INTO cost_type(name, display_name) values ("STANDARD","Standard"),("CONCESSION","Concession");

DROP TABLE IF EXISTS location;
CREATE TABLE location (
    name TEXT,
    display_name TEXT,
    address TEXT,
    latitude TEXT,
    longitude TEXT,
    url TEXT
);

DROP TABLE IF EXISTS member;
CREATE TABLE member (
    id INTEGER PRIMARY KEY,
    name TEXT,
    display_name TEXT,
    email TEXT,
    phone TEXT
);