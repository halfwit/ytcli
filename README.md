# Small Script To Interact With Youtube

## Usage

```
ytcli [action]

Actions:
	search	<keyword>
		will return a list of URLs and descritions
		using -t will also return the path to a thumbnail (in /tmp)
	user	<username>
		will return a list of videos by user
	channel	<channel name>
		will return a list of videos by channel
	channel-id <channel name>
		will return channel id of given channel
```

## Configuration

ytcli assumes that you have an API token for using the v3 Youtube API

Example: 

```
# Configuration .local/cfg/ytcli.conf
# Lines starting with # will be ignored
key=4TNSHS543tn5hNSTH54354
max_results=50
tmpdir=/tmp/yt
```
