-- PasteIndex
CREATE TABLE IF NOT EXISTS ENTRIES (
    ID            BIGSERIAL CONSTRAINT ENTRIES_PK PRIMARY KEY,
    RAW_TEXT      TEXT,
    LANGUAGE      TEXT,
    EXPIRES_AT    BIGINT,
    MAX_DOWNLOADS INT,
    DOWNLOADS     INT,
    AUTO_DELETE   BOOLEAN,
    PASSWORD      TEXT
);
-- im about to leave so lets finish up real quick 
CREATE UNIQUE INDEX IF NOT EXISTS ENTRIES_ID_UINDEX ON ENTRIES (ID);

-- We'll work on file index once we have paste index 100%