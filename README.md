# dwh-tasks

Время для выполнения заданий – примерно час и 30 минут. Ответы для пунктов 1 и 2 можно писать на любом диалекте SQL.
## 0.
У Вас ситуация: таблица T, несколько миллиардов записей. В колонке А - 5 разных возможных значений,  в колонке B - 50 разных возможных значений, в колонке C - 100 000 разных возможных значений. Распределение количества строк по знaчениям A, B, C - равномерное. О производительности INSERT, UPDATE,  DELETE - беспокоиться не надо. Но надо сделать макс.  возможно хорошие условия для
```
SELECT * FROM T WHERE A=? AND B=? AND C=?
```
Какой индекс (индексы) сделаете?  С коротким обьяснением почему такой индекс.
## 1.
```
Таблица Present
ID              int PK
DELETED         int
ID_PresentPaid  int
```
Написать скрипт удаления дублирующих записей из этой таблицы. Можно указать несколько способов.
## 2.
В системе есть подарки. Подарки группируются по категориям. Один подарок может войти в одну или несколько категорий. Так-же фиксируются факты (дата с временем и user_id) дарения этих подарков(c фиксацией категории). Один и тот же подарок можно дарить несколько раз. Подарок может  быть ни разу не подарен.
* 2.1 Нарисовать нормализированную схему баз данных для хранения этой информации.
Написать по этой базе данных SQL запросы(по желанию - как один Select или как скрипт из несколько команд):
* 2.2 Вывести строки с 20 по 40 из топа категорий в которых больше всего подарков.
* 2.3 Вывести топ 20 категорий подарков, которые чаще всего дарят.
* 2.4 Вывести топ 20 категорий подарков,  которые были подарены наибольшим количеством пользователей.
* 2.5 Вывести подарки, которые никогда не дарились.
* 2.6 Вывести таблицу – строки: подарки, колонки : дарился 1-10 раз, дарился 11-100 раз, дарился строго больше 100 раз.
* 2.7 Вывести категории подарков, которые дарились в марте этого года чаще, чем в апреле этого года.
* 2.8 По каждой категории вывести top 20 пользователей которые дарили больше всего подарки этой категории.

## 3.
Реализуйте на любом известном вам языке следующую программу, которая на вход получит файл со строками вида:
```
[DWH] name@domain||FirstName||LastName||Age||PhoneNumber [\DWH]
```
На выходе она должна возвращать файл формате CSV c полями:
```
"name","domain","FirstName","LastName","Age","PhoneNumber"
```
При этом в выходном файле все номера телефонов должны быть приведены к одному формату.
Отметьте, пожалуйста язык на котором реализована программа, и при использовании сторонних библиотек, перечислите их и обоснуйте необходимость их использования.
