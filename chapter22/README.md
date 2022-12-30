https://gin-gonic.com/zh-cn/docs/

https://getbootstrap.com/

https://pkg.go.dev

### start mysql
```
docker run --name mysql57 -v /u01/data/mysql:/var/lib/mysql -e MYSQL_ROOT_PASSWORD=oracle -p 3306:3306 -d mysql:5.7.40
docker exec -it mysql57 bash

mysql> create user 'root'@'%' identified by 'oracle';
Query OK, 0 rows affected (0.00 sec)

mysql> grant all privileges on *.* to 'root'@'%';
Query OK, 0 rows affected (0.00 sec)

mysql> create database oracle;
Query OK, 1 row affected (0.01 sec)
```