CREATE TABLE "users" (
  "user_id" varchar PRIMARY KEY,
  "username" varchar NOT NULL,
  "email" varchar NOT NULL,
  "password" varchar NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "updated_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "todos" (
  "todo_id" varchar PRIMARY KEY,
  "user_id" varchar NOT NULL,
  "title" varchar NOT NULL,
  "description" text NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "updated_at" timestamptz NOT NULL DEFAULT (now()),
  "is_complete" boolean NOT NULL DEFAULT (false)
);

ALTER TABLE "todos" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("user_id");
