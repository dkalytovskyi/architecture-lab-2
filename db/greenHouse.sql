-- Create tables.
DROP TABLE IF EXISTS "plants";
CREATE TABLE "plants"
(
    "id"   SERIAL PRIMARY KEY,
    "soilmoisturelevel" FLOAT NOT NULL,
    "soildatatimestamp" VARCHAR(50) NOT NULL
);

-- Insert demo data.
INSERT INTO "plants" VALUES (356, 0.64, "2019-10-20T19:02:10+03:00");
INSERT INTO "plants" VALUES (43, 0.12, "2019-10-20T23:50:05+03:00");
