## 需求
- 一张表里有数据重复，需要删除重复的行，保留其中一条数据

## 实现
```sql
DELETE 
FROM
    half_demon 
WHERE
    ( name, weapon ) IN (
SELECT
    t.name,
    t.weapon 
FROM
    ( SELECT name, weapon FROM half_demon GROUP BY name, weapon HAVING count( 1 ) > 1 ) t 
    ) 
    AND id NOT IN ( SELECT hd.minid FROM ( SELECT min( id ) AS minid FROM half_demon GROUP BY name, weapon HAVING count( 1 ) > 1 ) hd ) 
```
