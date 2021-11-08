package main

import (
	"github.com/abyu/stats/internal"
	log "github.com/sirupsen/logrus"
	"os"
)

func main() {
	context, err := internal.BuildContext()
	fatalErrCheck(err, "Failed to initialize application")

	err = context.DBMigration.AutoMigrate()
	fatalErrCheck(err, "Failed to run db migrations")

	context.FetchTask.RunBlocking()
}

// fatalErrCheck fails fatally upon an error.
func fatalErrCheck(err error, msg string) {
	if err != nil {
		log.WithError(err).Fatal(msg)
		os.Exit(1)
	}
}
