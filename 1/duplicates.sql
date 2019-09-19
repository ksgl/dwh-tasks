SET SYNCHRONOUS_COMMIT = 'off';

CREATE TABLE presents (
    ID              INT PRIMARY KEY,
    DELETED         INT,
    ID_PresentPaid  INT
)

-- 1.1 (self join) leaves rows with the lowest id

DELETE FROM presents AS p1 USING presents p2
WHERE p1.ID < p2.ID AND p1.DELETED = p2.DELETED AND p1.ID_PresentPaid = p2.ID_PresentPaid;

--  1.2 (self join) leaves rows with the highest id

DELETE FROM presents AS p1 USING presents p2
WHERE p1.ID > p2.ID AND p1.DELETED = p2.DELETED AND p1.ID_PresentPaid = p2.ID_PresentPaid;

-- 2

DELETE FROM presents
WHERE ID IN
    (SELECT ID FROM
        (SELECT ID, ROW_NUMBER() OVER (PARTITION BY DELETED, ID_PresentPaid
            ORDER BY ID) AS row_num
        FROM presents) p
        WHERE p.row_num > 1);

-- 3

CREATE TABLE presents_temp (LIKE presents);

INSERT INTO presents_temp(DELETED, ID_PresentPaid, ID)
SELECT DISTINCT ON (DELETED) DELETED, (ID_PresentPaid) ID_PresentPaid, ID
FROM presents;

DROP TABLE presents;

ALTER TABLE presents_temp RENAME TO presents;


