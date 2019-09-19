SET SYNCHRONOUS_COMMIT = 'off';


------------------------------------------------------------------------------------------------

CREATE TABLE presents (
    present_id      SERIAL PRIMARY KEY
);

CREATE TABLE categories (
    category_id    SERIAL PRIMARY KEY
);

CREATE TABLE present_categories (
    present_id      INT REFERENCES presents(present_id),
    category_id    INT REFERENCES categories(category_id),

    PRIMARY KEY (present_id,category_id)
);

CREATE TABLE users (
    user_id         SERIAL PRIMARY KEY
);

CREATE TABLE user_presents (
    present_id      INT REFERENCES presents(present_id),
    user_id         INT REFERENCES users(user_id),
    date            TIMESTAMPTZ DEFAULT now(),
    category_id    INT REFERENCES categories(category_id)

    PRIMARY KEY (present_id,user_id,date)
);

------------------------------------------------------------------------------------------------

-- 2.2 Вывести строки с 20 по 40 из топа категорий в которых больше всего подарков.

SELECT *
FROM present_categories
WHERE category_id IN (
	SELECT category_id
	FROM
	   present_categories
	GROUP BY
	   category_id
	ORDER BY COUNT(category_id) DESC
	LIMIT 20 OFFSET 20
);


-- 2.3 Вывести топ 20 категорий подарков, которые чаще всего дарят.

SELECT category_id, COUNT(category_id) AS cnt
FROM present_categories
GROUP BY category_id
ORDER BY cnt DESC
LIMIT 20;


-- 2.4 Вывести топ 20 категорий подарков, которые были подарены наибольшим количеством пользователей.

SELECT pc.category_id, COUNT(pc.category_id) AS cnt
FROM user_presents AS up
JOIN present_categories AS pc
USING (present_id)
GROUP BY pc.category_id
ORDER BY cnt DESC


-- 2.5 Вывести подарки, которые никогда не дарились.

SELECT presents.present_id
FROM presents
LEFT JOIN user_presents
ON presents.present_id = user_presents.present_id
WHERE user_presents.present_id IS NULL;


-- 2.6 Вывести таблицу – строки: подарки, колонки : дарился 1-10 раз, дарился 11-100 раз, дарился строго больше 100 раз.


-- 2.7 Вывести категории подарков, которые дарились в марте этого года чаще, чем в апреле этого года.

SELECT c1.category_id
FROM (
	SELECT category_id, count(category_id) AS cnt
	FROM user_presents
	WHERE date >= '2019-03-01 00:00:00' AND date < '2019-04-01 00:00:00'
	GROUP BY category_id
	) AS c1
JOIN (
	SELECT category_id, count(category_id) AS cnt
	FROM user_presents
	WHERE date >= '2019-04-01 00:00:00' AND date < '2019-05-01 00:00:00'
	GROUP BY category_id
) AS c2
ON c1.category_id = c2.category_id
WHERE c1.cnt > c2.cnt;


-- 2.8 По каждой категории вывести top 20 пользователей которые дарили больше всего подарки этой категории.

SELECT DISTINCT (user_id) user_id, category_id
FROM (
  SELECT ROW_NUMBER() OVER (PARTITION BY category_id ORDER BY user_id) AS r, up.*
  FROM user_presents AS up
) AS uc
WHERE uc.r <= 20
ORDER BY category_id;