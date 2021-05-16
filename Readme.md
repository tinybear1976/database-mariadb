---
module: hhyt/database/mariadb
function: 简单封装mariadb连接的基本内容
version: 0.1.0
path: hhyt/database/mariadb@v0.1.0
---

目录

[TOC]



# 引用

## go.mod

replace 的物理磁盘位置要根据物理目录的实际位置给定。其中 github.com/go-sql-driver/mysql 与 github.com/jmoiron/sqlx  可以在引用hhyt/database/mariadb后，使用go mod  tidy，让系统自动添加

```go
module mariadbtest

go 1.15

require (
	github.com/go-sql-driver/mysql v1.5.0 // indirect
	github.com/jmoiron/sqlx v1.2.0 // indirect
	hhyt/database/mariadb v0.1.0
)

replace hhyt/database/mariadb => ../../hhyt/database/mariadb@v0.1.0
```

## main.go

连接和基本操作。由例可见，数据库连接初始化可以单独进行，并将其连接保存至包内，直到使用Destroy()，手动销毁所有的连接指针。在进行后续的操作时，只需要使用连接标识来表明调用哪个MariaDB连接即可。

```go
package main

import (
	"fmt"
	"hhyt/database/mariadb"
)

func main() {

	MariadbInit()
	DbOptTest()
	mariadb.Destroy()
}

func MariadbInit() {
	err := mariadb.New("local", "127.0.0.1", "3306", "root", "123", "tcc")
	if err != nil {
		fmt.Println(err)
		return
	}
}

func DbOptTest() {
	db, err := mariadb.Connect("local")
	if err != nil {
		fmt.Println(err)
		return
	}
	var vals []string
	db.Select(&vals, "select username from users")
	fmt.Println(len(vals), vals)
}

```



# 函数

## New

建立并保存一个MariaDB连接，在之后的使用过程中，通过连接标识创建连接并进行数据库操作。在New的过程中模块会尝试连接数据库，如果发现服务器无法连接，则不会保存该数据库连接。

```go
func New(serverTag, ip, port, username, password, dbname string) error
```

入口参数：
| 参数名    | 类型   | 描述                             |
| --------- | ------ | -------------------------------- |
| serverTag | string | Mariadb数据库连接标识            |
| ip        | string | Mariadb数据库服务器地址          |
| port      | string | 数据库服务端口 （一般应为 3306） |
| username  | string | 数据库用户名                     |
| password  | string | 数据库用户密码                   |
| dbname    | string | 数据库名称                       |

返回值：正确返回nil，否则返回错误信息



## Destroy

销毁所有模块内保存的（由New创建的）连接池指针。

```go
func Destroy()
```

入口参数：无

返回值：无



## Connect

获得一个数据库连接。

```go
func Connect(serverTag string) (*sqlx.DB, error) 
```

入口参数：

| 参数名    | 类型   | 描述                  |
| --------- | ------ | --------------------- |
| serverTag | string | Mariadb数据库连接标识 |

返回值：

| 返回变量 | 类型     | 描述                                      |
| -------- | -------- | ----------------------------------------- |
|          | *sqlx.DB | 数据库连接指针                            |
|          | error    | 返回操作结果的错误信息，如果正确则返回nil |


