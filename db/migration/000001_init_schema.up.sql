CREATE TABLE "users" (
  "id" bigserial PRIMARY KEY,
  "username" varchar NOT NULL,
  "first_name" varchar NOT NULL,
  "second_name" varchar NOT NULL,
  "email" varchar NOT NULL,
  "email_confirm" BOOLEAN NOT NULL DEFAULT 'false',
  "roles" TEXT[] DEFAULT '{"user", "seller"}',
  "hashed_password" varchar NOT NULL,
  "account" bigserial,
  "address" bigserial,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "account" (
  "id" bigserial PRIMARY KEY,
  "balance" int NOT NULL
);

CREATE TABLE "address" (
  "id" bigserial PRIMARY KEY,
  "city" varchar NOT NULL,
  "line1" varchar NOT NULL,
  "line2" varchar NOT NULL,
  "postcode" varchar NOT NULL
);

CREATE TABLE "sessions" (
  "id" uuid PRIMARY KEY,
  "username" varchar NOT NULL,
  "token" varchar NOT NULL,
  "user_agent" varchar NOT NULL,
  "client_ip" varchar NOT NULL,
  "is_blocked" boolean NOT NULL DEFAULT false,
  "expires_at" timestamptz NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

ALTER TABLE "users" ADD FOREIGN KEY ("account") REFERENCES "account" ("id");

ALTER TABLE "users" ADD FOREIGN KEY ("address") REFERENCES "address" ("id");