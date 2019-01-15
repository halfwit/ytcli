package main

import (
	"flag"
	"log"
	"regexp"
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
	call := service.Search.List("id,snippet").Q(query).MaxResults(1)
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

func feedSearch(query string, service *youtube.Service) map[string]string {
	results := make(map[string]string)
	feed := "https://youtube.com/feeds/video.xml?channel_id="
	if *channel {
		call := service.Search.List("id,snippet").Q(query).MaxResults(*nresults)
		response, err := call.Do()
		if err != nil {
			log.Fatal(err)
		}
		for _, item := range response.Items {
			if item.Id.Kind == "youtube#channel" {
 				results[feed + item.Id.ChannelId] = item.Snippet.Title
			}
		}
	} else {
		// TODO: Remove unwanted group 
		r := regexp.MustCompile(`^((?:https?:)?\/\/)?((?:www|m)\.)?((?:youtube\.com|youtu.be))(\/(?:[\w\-]+\?v=|embed\/|v\/)?)([\w\-]+)(\S+)?$`)
		vid := r.FindStringSubmatch(query)
		call := service.Videos.List("id,snippet").Id(vid[5])
		id, err := call.Do()
		if err != nil {
			log.Fatal(err)
		}
		for _, item := range id.Items {
			results[feed + item.Snippet.ChannelId] = item.Snippet.ChannelTitle
		}
	}
	return results
}

func relatedSearch(query string, service *youtube.Service) map[string]string{
	results := make(map[string]string)
	r := regexp.MustCompile(`^((?:https?:)?\/\/)?((?:www|m)\.)?((?:youtube\.com|youtu.be))(\/(?:[\w\-]+\?v=|embed\/|v\/)?)([\w\-]+)(\S+)?$`)
	vid := r.FindStringSubmatch(query)
	call := service.Search.List("id,snippet").RelatedToVideoId(vid[5]).MaxResults(*nresults).Type("video")
	id, err := call.Do()
	if err != nil {
		log.Fatal(err)
	}
	for _, item := range id.Items {
		results[item.Id.VideoId] = item.Snippet.Title
	}
	return results
}

func runCommand(service *youtube.Service) map[string]string {
	query := strings.Join(flag.Args(), " ")
	switch {
	case *feed:
		return feedSearch(query, service)
	case *playlist:
		return playlistSearch(query, service)
	case *channel:
		return channelSearch(query, service)
	case *related:
		return relatedSearch(query, service)

	}
	return normalSearch(query, service)
}

