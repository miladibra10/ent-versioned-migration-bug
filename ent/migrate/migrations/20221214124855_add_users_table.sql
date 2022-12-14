-- create "users" table
CREATE TABLE "users" ("id" bigint NOT NULL GENERATED BY DEFAULT AS IDENTITY, "age" bigint NOT NULL, "name" character varying NOT NULL, PRIMARY KEY ("id"));
-- create "ent_types" table
CREATE TABLE "ent_types" ("id" bigint NOT NULL GENERATED BY DEFAULT AS IDENTITY, "type" character varying NOT NULL, PRIMARY KEY ("id"));
-- create index "ent_types_type_key" to table: "ent_types"
CREATE UNIQUE INDEX "ent_types_type_key" ON "ent_types" ("type");
-- add pk ranges for ('users') tables
INSERT INTO "ent_types" ("type") VALUES ('users');
