package duke_nukem

import (
	"log"
	"os"
)

var (
	c     = NewClient()
	debug = os.Getenv("DEBUG") == "1"
)

// Migrate takes user data extracted it from Drupal and
// POSTs the data to the 10duke registration form.
func Migrate(user User) (err error) {
	displayUser(user)
	err = c.Login()
	if err != nil {
		return err
	}
	err = c.RegisterUser(user)
	if err != nil {
		return err
	}
	return err
}

func displayUser(user User) {
	log.Printf("migrating user %s (%s %s %s from %s) to 10duke.\n",
		user.UID,
		user.FirstName,
		user.LastName,
		user.Email,
		user.Country)
	if debug {
		log.Printf("password: %s\n", user.Password)
	}
}
