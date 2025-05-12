package main

import (
	"database/sql"
	"gator/internal/config"
	"gator/internal/database"
	"log"
	"os"

	_ "github.com/lib/pq"
)

type state struct {
	db        *database.Queries
	configPtr *config.Config
}

type RSSItem struct {
	Title       string `xml:"title"`
	Link        string `xml:"link"`
	Description string `xml:"description"`
	PubDate     string `xml:"pubDate"`
}

type RSSFeed struct {
	Channel struct {
		Title       string    `xml:"title"`
		Link        string    `xml:"link"`
		Description string    `xml:"description"`
		Item        []RSSItem `xml:"item"`
	} `xml:"channel"`
}

func main() {
	cfg, err := config.Read()
	if err != nil {
		log.Fatalf("error reading config: %v", err)
	}
	st := &state{
		configPtr: &cfg,
	}

	//setting the DB connections
	db, err := sql.Open("postgres", cfg.DbURL)
	if err != nil {
		log.Fatal(err)
	}
	dbQueries := database.New(db)
	st.db = dbQueries

	cmds := commands{ //create the commands map
		command_map: make(map[string]func(*state, command) error),
	}

	cmds.register("login", handlerLogin)
	cmds.register("register", handlerRegister) // registerign the new command
	cmds.register("reset", handleReset)
	cmds.register("users", handleUsers)
	cmds.register("agg", handlerAgg)
	curr_args := os.Args
	if len(curr_args) < 2 {
		log.Fatal("Usage: cli <command> [args...]")
		return
	}

	new_command := command{
		name: curr_args[1],
		args: curr_args[2:],
	}

	if err := cmds.run(st, new_command); err != nil {
		log.Fatal(err)
		return

	}

}
