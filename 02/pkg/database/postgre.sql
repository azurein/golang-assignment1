CREATE TABLE "order" (
    "order_id" SERIAL PRIMARY KEY,
    "customer_name" varchar(100) NOT NULL,
    "ordered_at" timestamp NOT NULL,
    "created_at" timestamp NOT NULL,
    "updated_at" timestamp
);

CREATE TABLE "order_detail" (
    "order_detail_id" SERIAL PRIMARY KEY,
    "order_id" int4 NOT NULL,
    "item_code" varchar(50) NOT NULL,
    "description" varchar(255) NOT NULL,
    "quantity" int4 NOT NULL,
    "created_at" timestamp NOT NULL,
    "updated_at" timestamp
);
