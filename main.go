package main

import (
	"bufio"
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"

	"entgo.io/ent/dialect/sql"
	"github.com/a8m/entraffle/ent"
	"github.com/a8m/entraffle/ent/user"
	_ "github.com/mattn/go-sqlite3"
)

var (
	count    = flag.Int("count", 5, "number of retweeters to choose")
	filename = flag.String("filename", "retweeters", "name of the file holds the users who retweet ent")
)

func main() {
	ctx := context.Background()
	// Open a client and assume migration is managed by Atlas.
	client, err := ent.Open("sqlite3", "file:ent.db?_fk=1")
	if err != nil {
		log.Fatalf("failed opening connection to sqlite: %v", err)
	}
	defer client.Close()
	f, err := os.Open(*filename)
	if err != nil {
		log.Fatalf("failed opening file: %v", err)
	}
	defer f.Close()
	for sr := bufio.NewScanner(f); sr.Scan(); {
		client.User.Create().
			SetName(strings.TrimSpace(sr.Text())).
			ExecX(ctx)
	}
	names := client.User.Query().
		Order(sql.OrderByRand()).
		Unique(true).
		Limit(*count).
		Select(user.FieldName).
		StringsX(ctx)
	fmt.Println("Winners are:")
	for i, n := range names {
		fmt.Printf("%d. %s\n", i+1, n)
	}
}
