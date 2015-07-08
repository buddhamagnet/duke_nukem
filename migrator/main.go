package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/buddhamagnet/duke_nukem"
	_ "github.com/joho/godotenv/autoload"

	"os"
	"regexp"
	"strconv"
	"strings"
)

const (
	MIGRATE_SINGLE = 1
	MIGRATE_RANGE  = 2
	MIGRATE_SET    = 3
)

var (
	mode             int
	uids             string
	migrateUser      *regexp.Regexp
	migrateUserRange *regexp.Regexp
	migrateUserSet   *regexp.Regexp
)

func init() {
	flag.StringVar(&uids, "uids", "", "Drupal UID(s) to migrate")
	migrateUser = regexp.MustCompile("^[0-9]+$")
	migrateUserRange = regexp.MustCompile("^[0-9]+-[0-9]+$")
	migrateUserSet = regexp.MustCompile("^([0-9]+,)+")
}

func usage() {
	fmt.Println("usage: migrator --uids=<uids> (single: <uid>, range: <uid-uid>, set: <uid,uid,uid>")
	os.Exit(1)
}

func main() {

	flag.Parse()
	if uids == "" {
		usage()
	}
	// Determine whether we are migrating a single user, a range, or a set.
	switch true {
	case migrateUser.MatchString(uids):
		mode = MIGRATE_SINGLE
	case migrateUserRange.MatchString(uids):
		mode = MIGRATE_RANGE
	case migrateUserSet.MatchString(uids):
		mode = MIGRATE_SET
	default:
		usage()
	}

	switch mode {
	case MIGRATE_SINGLE:
		migrate(uids)
	case MIGRATE_RANGE:
		users := strings.Split(uids, "-")
		// UID range formats are already enforced by the constant regular expressions.
		start, _ := strconv.ParseInt(users[0], 10, 0)
		end, _ := strconv.ParseInt(users[1], 10, 0)
		for i := start; i <= end; i++ {
			migrate(strconv.Itoa(int(i)))
		}
	case MIGRATE_SET:
		users := strings.Split(uids, ",")
		for _, user := range users {
			migrate(user)
		}
	}
}

func migrate(uids string) {
	msg, err := runMigration(uids)
	if err != nil {
		log.Println(err)
	} else {
		log.Println(msg)
	}
}

func runMigration(uids string) (msg string, err error) {
	user, err := duke_nukem.RunUserQuery(uids)
	if err != nil {
		return "", fmt.Errorf("error with query for user %s: %s\n", uids, err)
	}
	err = duke_nukem.Migrate(user)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("user migrated to 10duke, visit %s to check status.", duke_nukem.BaseURI), nil
}
