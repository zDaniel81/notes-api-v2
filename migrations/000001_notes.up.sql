CREATE TABLE notes (
    ID UUID PRIMARY KEY DEFAULT gen_random_uuid() NOT NULL,
    Title text NOT NULL,
    Content text NOT NULL
);