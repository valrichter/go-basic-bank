CREATE TABLE
    "session" (
        "id" uuid PRIMARY KEY,
        "username" varchar NOT NULL,
        "refresh_token" varchar NOT NULL,
        "user_agent" varchar NOT NULL,
        "client_ip" varchar NOT NULL,
        "is_blocked" boolean NOT NULL DEFAULT false,
        "expires_at" timestamptz NOT NULL DEFAULT (now ())
    );

ALTER TABLE "session"
ADD
    FOREIGN KEY ("username") REFERENCES "users" ("username");