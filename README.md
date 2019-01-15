# ytcli - POSIX sh

## Installation:

	just include in your path. 

## Usage: 
	
	ytcli [search|playlist|user|channel] [ -t ] KEYWORD
	ytcli [channel-id] KEYWORD

# ytcli - Go

## Building: 

	`go get github.com/halfwit/ytcli`

## Installation:

	`go install github.com/halfwit/ytcli`
	

Requests:

```
# Usage
 ytcli [ -trpucfa ] KEYWORD
	
# Search for a list of videos matching string
ytcli <string>

# Search for videos related to url
ytcli -r 'https://youtube.com/someurl'

# Search for playlists matching strings
ytcli -p <string>

# List of videos by user
ytcli -u <user>

# List of playlits by user
ytcli -u -p <user>

# List of videos by channel
ytcli -c <channel>


# Link to channels' RSS feed
ytcli -f <channel>

# Normally ytcli will query the factotum for an API key (this requires plan9port, plan9, etc)
# To set a key explicitely:
ytcli -a myapikey

# Include thumbnails
-t 


