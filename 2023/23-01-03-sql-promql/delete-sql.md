# MySQL 中 delete 语句的子查询限制

```sql
delete from student where id = (select max(id) from student);
 
[Err] 1093 - You can't specify target table 'student' for update in FROM clause
```
描述: 如果子查询的 from 子句和更新、删除对象使用同一张表，会出现上述错误。

解决方法: 通过给 from 子句中的结果集起别名。
```sql
delete from student where id = (select n.max_id from (select max(id) as max_id from student) as n);
```

上述情况对于 in 子句也适用

```sql
 
delete from student where id in (select id from student where id > 30);
 
[Err] 1093 - You can't specify target table 'student' for update in FROM clause
```
解决方法同上:

```sql
 
delete from student where id in (select n.id from (select id from student where id > 30) as n);
````