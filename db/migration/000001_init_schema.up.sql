CREATE TABLE "accounts" (
  "id" BIGSERIAL PRIMARY KEY,
  "name" varchar NOT NULL,
  "balance" bigint NOT NULL DEFAULT 0,
  "phone_number" varchar NOT NULL,
  "currency" varchar NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "transaction" (
  "id" BIGSERIAL PRIMARY KEY,
  "reciever_id" bigint NOT NULL,
  "sender_id" bigint NOT NULL,
  "amount" bigint NOT NULL,
  "status" varchar NOT NULL DEFAULT 'pending',
  "type" varchar NOT NULL,
  "currency" varchar NOT NULL,
  "message" varchar,
  "deadline" timestamptz,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE INDEX ON "accounts" ("phone_number");

CREATE INDEX ON "transaction" ("reciever_id");

CREATE INDEX ON "transaction" ("sender_id");

CREATE INDEX ON "transaction" ("reciever_id", "sender_id");

ALTER TABLE "transaction" ADD FOREIGN KEY ("reciever_id") REFERENCES "accounts" ("id");

ALTER TABLE "transaction" ADD FOREIGN KEY ("sender_id") REFERENCES "accounts" ("id");