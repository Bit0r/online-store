package model

import "golang.org/x/crypto/bcrypt"

func VerifyUser(name, passwd string) (id uint64) {
	var hash string
	query := `select id, passwd
	from user
	where name = ?`

	if db.QueryRow(query, name).Scan(&id, &hash) != nil {
		// 验证用户名是否正确
		return 0
	}

	if bcrypt.CompareHashAndPassword([]byte(hash), []byte(passwd)) != nil {
		// 验证密码是否正确
		return 0
	}

	return
}

func AddUser(name, passwd string) (id uint64) {
	query := `insert ignore into user(name, passwd) values(?, ?)`
	hash, _ := bcrypt.GenerateFromPassword([]byte(passwd), 0)
	result, err := db.Exec(query, name, hash)
	if count, _ := result.RowsAffected(); err != nil || count == 0 {
		return 0
	} else {
		id, _ := result.LastInsertId()
		return uint64(id)
	}
}
