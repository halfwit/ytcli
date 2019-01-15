package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/user"
	"net/http"

	"google.golang.org/api/googleapi/transport"
	"google.golang.org/api/youtube/v3"
	"bitbucket.org/mischief/libauth"
)

var (
	//username = flag.String("u", "", "List videos uploaded by user")
	apiKey	 = flag.String("a", "", "Use API key instead of factotum")
	nresults = flag.Int64("m", 50, "Number of results per query")
	playlist = flag.Bool("p", false, "List playlists that match query")
	channel  = flag.Bool("c", false, "List videos uploaded by channel")
	//thumbs   = flag.Bool("t", false, "Return link to thumbnail as well")
	related  = flag.Bool("r", false, "List videos related to <URL>")
	feed     = flag.Bool("f", false, "Return link to RSS feed of user/channel (Can be used with -c, -u, or a video URL)")
)



func main() {
	flag.Parse()
	if flag.Lookup("h") != nil {
		flag.Usage()
		os.Exit(1)
	}
	if *apiKey == "" {
		u, err := user.Current()
		if err != nil {
			log.Fatal(err)
		}
		key, err := libauth.Getuserpasswd( "proto=pass service=ytcli user=%s", 
			u.Username,
		)
		if err != nil {
			log.Fatal(err)
		}
		*apiKey = key.Password
	}
	client := &http.Client{
		Transport: &transport.APIKey{ Key: *apiKey },
	}
	service, err := youtube.New(client)
	if err != nil {
		log.Fatalf("Error with youtube: %v", err)
	}
	results := runCommand(service)

	for id, title := range results {
		fmt.Printf("%s - %v\n", title, id)
	}
}
