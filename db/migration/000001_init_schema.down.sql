ALTER TABLE IF EXISTS "articles" DROP CONSTRAINT IF EXISTS "articles_author_username_fkey";
DROP TABLE IF EXISTS "articles";
DROP TABLE IF EXISTS "users";