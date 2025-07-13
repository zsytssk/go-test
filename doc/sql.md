```sql

BEGIN TRANSACTION;

-- 步骤 1~4 之间的语句写在这里
CREATE TABLE dir_new (
    id INTEGER PRIMARY KEY,
    name TEXT,
    priority INT,
    age INTEGER -- 修改了字段类型
);

INSERT INTO dir_new (id, name, priority)
SELECT id, name, priority  FROM dir;

DROP TABLE dir;
ALTER TABLE dir_new RENAME TO dir;

COMMIT;



SELECT COUNT(*)
FROM INFORMATION_SCHEMA.COLUMNS
WHERE
    TABLE_NAME = 'dir' AND
    COLUMN_NAME = 'name'


SELECT COUNT(*) FROM pragma_table_info('dir') WHERE name = 'name';

SELECT EXISTS (
  SELECT 1 FROM pragma_table_info('dir') WHERE name = 'name';
) AS has_column;


ALTER TABLE users ADD COLUMN age INTEGER;
```
