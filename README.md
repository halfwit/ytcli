#Small Script To Interact With Youtube

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

## Example

```
ytcli search -t koalas

Cutest Koala Videos EVER
https://www.youtube.com/watch?v=S4UG_l-CHF4
https://i.ytimg.com/vi/S4UG_l-CHF4/hqdefault.jpg
Koala Gets Kicked Out Of Tree and Cries!
https://www.youtube.com/watch?v=O0cAx1jLbJk
https://i.ytimg.com/vi/O0cAx1jLbJk/hqdefault.jpg
Cute Koalas Playing 🐨 Funny Koala Bears [Funny Pets]
https://www.youtube.com/watch?v=1R-QFQGWYuc
https://i.ytimg.com/vi/1R-QFQGWYuc/hqdefault.jpg
10 Fun Facts About Koalas
https://www.youtube.com/watch?v=hoataOsqfhc
https://i.ytimg.com/vi/hoataOsqfhc/hqdefault.jpg
Koala fight
https://www.youtube.com/watch?v=djK_ucSYpaw
https://i.ytimg.com/vi/djK_ucSYpaw/hqdefault.jpg
Koala Gives Stinky Hugs!
https://www.youtube.com/watch?v=pw-Njg0Hess
https://i.ytimg.com/vi/pw-Njg0Hess/hqdefault.jpg
Fighting Koalas Shooed From Highway by Concerned Motorist
https://www.youtube.com/watch?v=eSrpNwZUeQY
https://i.ytimg.com/vi/eSrpNwZUeQY/hqdefault.jpg
Cute Adorable Koala takes water from human
https://www.youtube.com/watch?v=RXDLodcCRuM
https://i.ytimg.com/vi/RXDLodcCRuM/hqdefault.jpg
koala joey's most adorable home video of all time
https://www.youtube.com/watch?v=cU8v4vZbFPc
https://i.ytimg.com/vi/cU8v4vZbFPc/hqdefault.jpg
La vida secreta de los koalas: LA ÉPOCA DE REPRODUCCIÓN | Grandes documentales
https://www.youtube.com/watch?v=9F9zhc8F6Go
https://i.ytimg.com/vi/9F9zhc8F6Go/hqdefault.jpg
[...]
```
