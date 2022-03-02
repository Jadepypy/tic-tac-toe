CREATE TABLE users (
                         id int PRIMARY KEY,
                         name VARCHAR ( 50 ) NOT NULL
);

CREATE TABLE "states" (
                          "id" int PRIMARY KEY,
                          "description" VARCHAR (50) NOT NULL
);

CREATE TABLE "records" (
                           "id" SERIAL PRIMARY KEY,
                           "state_id" int NOT NULL,
                           "user_id" int NULL
);

ALTER TABLE "records" ADD FOREIGN KEY ("state_id") REFERENCES "states" ("id");

ALTER TABLE "records" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

INSERT INTO users (id, name) VALUES (0, 'Player1'), (1, 'Player2');

INSERT INTO states (id, description) VALUES (0, 'Ongoing'),
                                            (1, 'Draw'),
                                            (2, 'Win'),
                                            (3, 'Invalid');
