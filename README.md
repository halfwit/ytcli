# Small Script To Interact With Youtube

## Usage

```
ytcli [-t] [action]

Actions:
	search	<keyword>
		will return a list of URLs and descritions
	user	<username>
		will return a list of videos by user
	playlist
		will return a list of playlists
	channel	<channel name>
		will return a list of videos by channel
	channel-id <channel name>
		will return channel id of given channel
	-t 
		return thumbnail as well
```

## Configuration

ytcli assumes that you have an API token for using the v3 Youtube API

Example: 

```
# Configuration .local/cfg/ytcli.conf
# Lines starting with # will be ignored
key=4TNSHS543tn5hNSTH54354
max_results=50
```
