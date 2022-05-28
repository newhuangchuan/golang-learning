# 创建⽤户
## 语法：
```shell
create user ⽤户名[@主机名] [identified by '密码'];
```
- 主机名默认值为%，表⽰这个⽤户可以从任何主机连接mysql服务器
- 密码可以省略，表⽰⽆密码登录

# 修改密码【3种⽅式】
## ⽅式1：通过管理员修改密码
```shell
SET PASSWORD FOR '⽤户名'@'主机' = PASSWORD('密码');
```
## ⽅式2：create	user	⽤户名[@主机名]	[identi>ied	by	'密码'];
```shell
set password = password('密码');
```
## ⽅式3：通过修改mysql.user表修改密码
```shell
use mysql;
update user set authentication_string = password('321') where user = 
'test1' and host = '%';
flush privileges;
```
> 注意：通过表的⽅式修改之后，需要执⾏flush privileges;才能对⽤户⽣效。

# 给⽤户授权
> 创建⽤户之后，需要给⽤户授权，才有意义。
## 语法：
```shell
grant privileges ON database.table TO 'username'[@'host'] [with grant 
option]
```
### grant命令说明：
- priveleges	(权限列表)，可以是all，表⽰所有权限，也可以是select、update等权
  限，多个权限之间⽤逗号分开。
  
- ON	⽤来指定权限针对哪些库和表，格式为数据库.表名 ，点号前⾯⽤来指定数据库
  名，点号后⾯⽤来指定表名，*.* 表⽰所有数据库所有表。
  
- TO	表⽰将权限赋予某个⽤户,	格式为username@host，@前⾯为⽤户名，@后⾯接限
  制的主机，可以是IP、IP段、域名以及%，%表⽰任何地⽅。
  
- WITH	GRANT	OPTION	这个选项表⽰该⽤户可以将⾃⼰拥有的权限授权给别⼈。注
  意：经常有⼈在创建操作⽤户的时候不指定WITH	GRANT	OPTION选项导致后来该⽤
  户不能使⽤GRANT命令创建⽤户或者给其它⽤户授权。 备注：可以使⽤GRANT重复
  给⽤户添加权限，权限叠加，⽐如你先给⽤户添加⼀个select权限，然后又给⽤户添
  加⼀个insert权限，那么该⽤户就同时拥有了select和insert权限。
  
### ⽰例：
```shell
grant all on *.* to 'test1'@‘%’;
```
> 说明：给test1授权可以操作所有库所有权限，相当于dba	
```shell
grant select on seata.* to 'test1'@'%';
```
> 说明：test1可以对seata库中所有的表执⾏select	
```shell
grant select,update on seata.* to 'test1'@'%';
```
> 说明：test1可以对seata库中所有的表执⾏select、update	
```shell
grant select(user,host) on mysql.user to 'test1'@'localhost';
```
> 说明：test1⽤户只能查询mysql.user表的user,host字段

# 查看⽤户有哪些权限
> show	grants	for	'⽤户名'[@'主机']
## 主机可以省略，默认值为%，⽰例：
```shell
mysql> show grants for 'test1'@'localhost';
+--------------------------------------------------------------------+
| Grants for test1@localhost |
+--------------------------------------------------------------------+
| GRANT USAGE ON *.* TO 'test1'@'localhost' |
| GRANT SELECT (host, user) ON `mysql`.`user` TO 'test1'@'localhost' |
+--------------------------------------------------------------------+
2 rows in set (0.00 sec)
```
## show	grants;
> 查看当前⽤户的权限，如：
```shell
mysql> show grants;
+---------------------------------------------------------------------
+
| Grants for root@localhost 
|
+---------------------------------------------------------------------
+
| GRANT ALL PRIVILEGES ON *.* TO 'root'@'localhost' WITH GRANT OPTION 
|
| GRANT ALL PRIVILEGES ON `test`.* TO 'root'@'localhost' 
|
| GRANT DELETE ON `seata`.* TO 'root'@'localhost' 
|
| GRANT PROXY ON ''@'' TO 'root'@'localhost' WITH GRANT OPTION 
|
+---------------------------------------------------------------------
+4 rows in set (0.00 sec)
```
# 撤销⽤户的权限
## 语法
> revoke privileges ON database.table FROM '⽤户名'[@'主机'];
> 可以先通过show grants命令查询⼀下⽤户对于的权限，然后使⽤revoke命令撤销⽤户
对应的权限，⽰例：
```shell
mysql> show grants for 'test1'@'localhost';
+--------------------------------------------------------------------+
| Grants for test1@localhost |
+--------------------------------------------------------------------+
| GRANT USAGE ON *.* TO 'test1'@'localhost' |
| GRANT SELECT (host, user) ON `mysql`.`user` TO 'test1'@'localhost' |
+--------------------------------------------------------------------+
2 rows in set (0.00 sec)
mysql> revoke select(host) on mysql.user from test1@localhost;
Query OK, 0 rows affected (0.00 sec)
mysql> show grants for 'test1'@'localhost';
+--------------------------------------------------------------+
| Grants for test1@localhost |
+--------------------------------------------------------------+
| GRANT USAGE ON *.* TO 'test1'@'localhost' |
| GRANT SELECT (user) ON `mysql`.`user` TO 'test1'@'localhost' |
+--------------------------------------------------------------+
2 rows in set (0.00 sec)
```
上⾯我们先通过grants命令查看test1的权限，然后调⽤revoke命令撤销对mysql.user表
host字段的查询权限，最后又通过grants命令查看了test1的权限，和预期结果⼀致。
# 删除⽤户【2种⽅式】
## ⽅式1：
> drop	user	'⽤户名'[@‘主机’]，⽰例：
```shell
mysql> drop user test1@localhost;
Query OK, 0 rows affected (0.00 sec)
```
drop的⽅式删除⽤户之后，⽤户下次登录就会起效。
## ⽅式2：
> 通过删除mysql.user表数据的⽅式删除，如下：
```shell
delete from user where user='⽤户名' and host='主机';
flush privileges;
```
注意通过表的⽅式删除的，需要调⽤flush privileges;刷新权限信息（权限启动的时
候在内存中保存着，通过表的⽅式修改之后需要刷新⼀下）。
## 授权原则说明
- 只授予能满⾜需要的最⼩权限，防⽌⽤户⼲坏事，⽐如⽤户只是需要查询，那就只给
  select权限就可以了，不要给⽤户赋予update、insert或者delete权限
  
- 创建⽤户的时候限制⽤户的登录主机，⼀般是限制成指定IP或者内⽹IP段
- 初始化数据库的时候删除没有密码的⽤户，安装完数据库的时候会⾃动创建⼀些⽤
  户，这些⽤户默认没有密码
  
- 为每个⽤户设置满⾜密码复杂度的密码
- 定期清理不需要的⽤户，回收权限或者删除⽤户

# DDL常⻅操作汇总
> DDL：Data	De[ine	Language数据定义语⾔，主要⽤来对数据库、表进⾏⼀些管理操作。
如：建库、删库、建表、修改表、删除表、对列的增删改等等。
##  库的管理
### 创建库
```shell
create database [if not exists] 库名;
```
### 删除库
```shell
drop databases [if exists] 库名;
```
### 建库通⽤的写法
```shell
drop database if exists 旧库名;
create database 新库名;
```
### 示例
```shell
mysql> show databases like 'javacode2018';
+-------------------------+
| Database (javacode2018) |
+-------------------------+
| javacode2018 |
+-------------------------+
1 row in set (0.00 sec)
mysql> drop database if exists javacode2018;
Query OK, 0 rows affected (0.00 sec)
mysql> show databases like 'javacode2018';
Empty set (0.00 sec)
mysql> create database javacode2018;
Query OK, 1 row affected (0.00 sec)
# show databases like 'javacode2018';列出javacode2018库信息。
```
## 表管理
### 创建表
```shell
create table 表名(
				字段名1 类型[(宽度)] [约束条件] [comment '字段说明'],
				字段名2 类型[(宽度)] [约束条件] [comment '字段说明'],
				字段名3 类型[(宽度)] [约束条件] [comment '字段说明']
)[表的⼀些设置];
```

