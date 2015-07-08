package duke_nukem

import (
	"os"
	"strings"

	_ "github.com/go-sql-driver/mysql" // blank import for mysql driver.
	"github.com/jmoiron/sqlx"
)

var dbCreds = []string{
	os.Getenv("DATABASE_USER"),
	":",
	os.Getenv("DATABASE_PWD"),
	"@tcp(",
	os.Getenv("DATABASE_HOST"),
	":3306)/",
	os.Getenv("DATABASE_NAME"),
}

var queryUser = `SELECT u.uid, u.name, pass, mail, created, status, p.given_name, p.family_name, c.country_cd FROM users u
                 INNER JOIN ec_profile p ON u.uid = p.uid
                 INNER JOIN ec_country c ON c.country_id = p.country_id
                 WHERE u.uid = ?`

// dbConnect connects to the database via sqlx.
func dbConnect() (*sqlx.DB, error) {
	db, err := sqlx.Connect("mysql", strings.Join(dbCreds, ""))
	if err != nil {
		return nil, err
	}
	return db, nil
}

// RunUserQuery retrieves a Drupal user from the
// database and marshalls it into a struct.
func RunUserQuery(args ...string) (User, error) {

	var user User
	db, err := dbConnect()

	if err != nil {
		return user, err
	}
	defer db.Close()

	err = db.Get(&user, queryUser, strings.Join(args, ""))

	if err != nil {
		return user, err
	}
	user.Password = generatePassword()
	return user, nil
}
