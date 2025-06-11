package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql" // 匿名导入
)

// 查询多行数据
func queryMultipleRows(db *sql.DB) {
	// 执行查询
	rows, err := db.Query("SELECT id, username, email FROM users WHERE active = ?", true)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close() // 重要：确保关闭rows

	fmt.Println("用户列表:")
	for rows.Next() {
		var (
			id       int
			username string
			email    string
		)
		if err := rows.Scan(&id, &username, &email); err != nil {
			log.Fatal(err)
		}
		fmt.Printf("ID: %d, 用户名: %s, 邮箱: %s\n", id, username, email)
	}

	// 检查遍历过程中是否有错误
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}
}

// 插入数据
func insertData(db *sql.DB) {
	// 准备插入语句
	stmt, err := db.Prepare("INSERT INTO users(username, email) VALUES(?, ?)")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	// 执行插入
	res, err := stmt.Exec("newuser", "newuser@example.com")
	if err != nil {
		log.Fatal(err)
	}

	// 获取插入的ID
	id, err := res.LastInsertId()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("插入成功，ID为: %d\n", id)
}

// 更新数据
func updateData(db *sql.DB) {
	// 执行更新
	res, err := db.Exec("UPDATE users SET email = ? WHERE id = ?",
		"updated@example.com", 1)
	if err != nil {
		log.Fatal(err)
	}

	// 获取影响的行数
	rowsAffected, err := res.RowsAffected()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("更新成功，影响了 %d 行\n", rowsAffected)
}

// 删除数据
func deleteData(db *sql.DB) {
	// 执行删除
	res, err := db.Exec("DELETE FROM users WHERE id = ?", 1)
	if err != nil {
		log.Fatal(err)
	}

	// 获取影响的行数
	rowsAffected, err := res.RowsAffected()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("删除成功，删除了 %d 行\n", rowsAffected)
}

func main() {
	dbUser := "rongan"
	dbPass := "rongandb"
	dbHost := "183.221.243.20"
	dbPort := "3306"
	dbName := "mysql"

	// 创建DSN(Data Source Name)
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True",
		dbUser, dbPass, dbHost, dbPort, dbName)

	// 连接数据库
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("无法连接数据库: %v", err)
	}
	defer db.Close() // 确保程序退出前关闭连接

	// 测试连接是否成功
	err = db.Ping()
	if err != nil {
		log.Fatalf("数据库连接测试失败: %v", err)
	}

	fmt.Println("成功连接到MySQL数据库!")

}
