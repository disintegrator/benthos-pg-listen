CREATE TABLE IF NOT EXISTS outbox_items (
  id varchar(20),
  created_at TIMESTAMP,
  payload JSON,
  attempt_count INT,
  PRIMARY KEY (id)
);

CREATE OR REPLACE FUNCTION notify_outbox_items_insert()
   RETURNS TRIGGER
   LANGUAGE plpgsql
AS $$
BEGIN
   PERFORM pg_notify('outbox_items__insert', '');
   RETURN NULL;
END;
$$;

CREATE TRIGGER outbox_items__insert
AFTER INSERT ON outbox_items
REFERENCING NEW TABLE AS outbox_items__inserted
FOR EACH STATEMENT
EXECUTE FUNCTION notify_outbox_items_insert();
