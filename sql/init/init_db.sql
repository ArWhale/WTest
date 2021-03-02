DO $$
BEGIN
  IF NOT EXISTS (
           SELECT table_name
           FROM   information_schema.tables
           WHERE  table_schema = 'public'
             AND    table_name = 'goose_db_version'
             )
  THEN
    EXECUTE 'DROP SCHEMA IF EXISTS public CASCADE';
    EXECUTE 'CREATE SCHEMA public';
    EXECUTE 'GRANT ALL ON SCHEMA public TO postgres';
    EXECUTE 'GRANT ALL ON SCHEMA public TO public';
  END IF;

END
$$;
