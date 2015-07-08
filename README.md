###DUKE NUKEM

A set of tools for working with and migrating user data to the [10duke](http://www.10duke.com) system.

###MIGRATOR

Migrator is a tool to retrieve a user from Drupal and migrate them to 10duke.

* Navigate into the migrate/migrator folder.
* Run ```go install```
* Run ```migrator --uid=<uid>``` and check the output.

We also provide a shell script thta facilitates batch migrations.

* Ensure you have built the migrator binary as outlined above.
* Run ```./migrate.sh``` <start_of_range> <end_of_range> and feel the migration love.
