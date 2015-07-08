package duke_nukem

// User represents a collection of user data
// extracted from Drupal for migration  to 10duke.
type User struct {
	UID       string `db:"uid"`
	UserName  string `db:"name"`
	FirstName string `db:"given_name"`
	LastName  string `db:"family_name"`
	Password  string `db:"pass"`
	Email     string `db:"mail"`
	Created   string `db:"created"`
	Status    int    `db:"status"`
	Country   string `db:"country_cd"`
}
