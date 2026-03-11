CREATE TABLE IF NOT EXISTS users (
    "id" SERIAL PRIMARY KEY,
    "user_name" VARCHAR(255) NOT NULL,
    "first_name" VARCHAR(255) NOT NULL,
    "last_name" VARCHAR(255),
    "created_at" TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    "updated_at" TIMESTAMPTZ,
    "deleted_at" TIMESTAMPTZ
);