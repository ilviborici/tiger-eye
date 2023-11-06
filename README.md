## TigerEye


  
  <img align="center" src="assets/tiger-eye.png">

A simplistic backdoor built to work with Windows machines using Telegram as a C2 server. At this time it is limited because it doesnt support multi client yet and only works with one client and API key for each session. This means that if you want to control another machine, you need to build a new executable which includes a new Telegram bot API.

## Features
   
	/help - Show all commands.
	/kill - Kill current running backdoor.
	/remove - Self-Remove current backdoor.	
	/shell, /sh <cmd> - Reverse shell from client machine.
	/download <path> - Download file from client machine.
	/webload <filename> <url> - Download file from website.
	/screen <filename> - Screenshot client display.
	/geoip <city-mmdb> - Retrieve the client geolocation from IPv4.

	/registry <method> <regname> <value> <path>

		--set		Set registry key to client machine.
		--del		Delete registry key from client machine.
	
	/browser <browser>

		--dc		Dumping Chrome browser database files.
		--dm		Dumping MS Edge browser database files.

	/clipboard - Retrieve Clipboard data from client.



## Bot API
To build the Telegram bot and get its API key, you can go to <a href="https://t.me/botfather">BotFather</a> and create a new bot.



