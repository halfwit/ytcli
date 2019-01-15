package main

import (
	"flag"
	"log"
	"strings"

	"google.golang.org/api/youtube/v3"
)

func playlistSearch(query string, service *youtube.Service) map[string]string {
	results := make(map[string]string)
	call := service.Search.List("id,snippet").Q(query).MaxResults(*nresults)
	response, err := call.Do()
	if err != nil {
		log.Fatal(err)
	}
	for _, item := range response.Items {
		if item.Id.Kind == "youtube#playlist" {
			results["https://youtu.be/" + item.Id.PlaylistId] = item.Snippet.Title
		}
	}
	return results
}

func channelSearch(query string, service *youtube.Service) map[string]string {
	// Fetch channel ID
	results := make(map[string]string)
	call := service.Search.List("id,snippet").Q(query).MaxResults(*nresults)
	id, err := call.Do()
	if err != nil {
		log.Fatal(err)
	}
	call2 := service.Search.List("id,snippet").MaxResults(*nresults).ChannelId(id.Items[0].Id.ChannelId)
	response, err := call2.Do()
	if err != nil {
		log.Fatal(err)
	}
	for _, item := range response.Items {
		if item.Id.Kind == "youtube#video" {
			results["https://youtu.be/" + item.Id.VideoId] = item.Snippet.Title
		}
	}
	return results
}

func normalSearch(query string, service *youtube.Service) map[string]string {
	results := make(map[string]string)
	call := service.Search.List("id,snippet").Q(query).MaxResults(*nresults)
	response, err := call.Do()
	if err != nil {
		log.Fatal(err)
	}
	for _, item := range response.Items {
		if item.Id.Kind == "youtube#video" {
			results["https://youtu.be/" + item.Id.VideoId] = item.Snippet.Title
		}
	}
	return results
}

func feedSearch(query string, service *youtube.Service) {
	//call := service.Search.Channel(query)
	//response, err := call.Do()
	//if err != nil {
	//	log.Fatal(err)
	//}
	//url := "https://youtube.com/feeds/videos.xml?channel_id"
	//fmt.Printf("%s=%v", url, item.Id.ChannelId)
}

func runCommand(service *youtube.Service) map[string]string {
	query := strings.Join(flag.Args(), " ")
	switch {
	case *feed:
		if *channel {
			feedSearch(query, service)
			return nil
		}
		//if user {
			//feedearch(query, service)
			//return nil
		//}
		//else {
			//parse URL for ID 
		//}
	case *playlist:
		return playlistSearch(query, service)
	case *channel:
		return channelSearch(query, service)

	}
	return normalSearch(query, service)
}
