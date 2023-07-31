CREATE TABLE "user" (
    "user_id" SERIAL PRIMARY KEY,
    "username" varchar(50) NOT NULL,
    "email" varchar(50) NOT NULL,
    "password" varchar(255) NOT NULL,
    "age" int4 NOT NULL,
    "created_at" timestamp NOT NULL,
    "updated_at" timestamp NOT NULL,
    "deleted" boolean default false,
    UNIQUE("username", "email")
);

CREATE TABLE "socialmedia" (
    "socialmedia_id" SERIAL PRIMARY KEY,
    "user_id" int4 NOT NULL,
    "name" varchar(255) NOT NULL,
    "socialmedia_url" varchar(500) NOT NULL,
    "created_at" timestamp NOT NULL,
    "updated_at" timestamp NOT NULL,
    "deleted" boolean default false
);

CREATE TABLE "photo" (
    "photo_id" SERIAL PRIMARY KEY,
    "user_id" int4 NOT NULL,
    "title" varchar(255) NOT NULL,
    "caption" varchar(255),
    "photo_url" varchar(255) NOT NULL,
    "created_at" timestamp NOT NULL,
    "updated_at" timestamp NOT NULL,
    "deleted" boolean default false
);

CREATE TABLE "comment" (
    "comment_id" SERIAL PRIMARY KEY,
    "photo_id" int4 NOT NULL,
    "user_id" int4 NOT NULL,
    "message" varchar(1000) NOT NULL,
    "created_at" timestamp NOT NULL,
    "updated_at" timestamp NOT NULL,
    "deleted" boolean default false
);
