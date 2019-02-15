-- Migration: brags
-- Created at: 2019-02-15 16:21:50
-- ====  UP  ====

BEGIN;
  CREATE TABLE brags (
    brag_id bigserial primary key,
    message text NOT NULL
  );
COMMIT;

-- ==== DOWN ====

BEGIN;
  DROP TABLE IF EXISTS brags;
COMMIT;
