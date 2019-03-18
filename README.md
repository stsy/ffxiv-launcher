# FFXIV Launcher
An unofficial FFXIV cli launcher written in Go

This is a auto-launcher if password is stored, but it's not safe to store passwords in plain text **(not recommended)**

![v01](https://user-images.githubusercontent.com/4086225/54465737-f7c65d80-477c-11e9-81f9-6df033e2bff5.PNG)


# Installation

#### Build from source
1) Download and install Go https://golang.org/dl/
2) ```go install github.com/stsy/ffxiv-launcher```
3) Create a folder named config in the same folder as the binary
4) Store [config.json](https://raw.githubusercontent.com/stsy/ffxiv-launcher/master/config/config.example.json) file in it

#### Download
https://github.com/stsy/ffxiv-launcher/releases


# Config

<sub>Remember SqEx ID?</sub>
```javascript
"user_id": "YOUR_ID"
```
#### Important!

<sub>If you are using token (and you should):</sub>
```javascript
"token": false //set this to true if you are using one-time password / token.
```
<sub>Set expansion value:</sub>
```javascript
"expansion":"2"
//"0": A Realm Reborn
//"1": Heavensward
//"2": Stormblood
```
<sub>Set the correct game/boot path:</sub>
```javascript
"boot":"PATH_TO_BOOT",
"game":"PATH_TO_GAME"
// eg. C:\\Program Files (x86)\\SquareEnix\\FINAL FANTASY XIV - A Realm Reborn\\game\\
```
<sub>config.example.json</sub>
```json
{
    "auth": {
        "user_id": "",
        "password": "",
        "token": false,
        "session": {
            "date": "",
            "id": ""
        }
    },
    "launcher": {
        "user_agent": "SQEXAuthor/2.0.0(Windows 6.2; ja-jp; 1064d87a30)",
        "oauth": {
            "regex_stored": "<input type=\"hidden\" name=\"_STORED_\" value=\"(?P<_STORED_>.*)\"",
            "regex_sid": "login=auth,ok,sid,(?P<SID>.+?),",
            "get": "https://ffxiv-login.square-enix.com/oauth/ffxivarr/login/top",
            "post": "https://ffxiv-login.square-enix.com/oauth/ffxivarr/login/login.send"
        }
    },
    "game":{
        "expansion":"2",
        "dx11":true,
        "path":{
            "boot":"C:\\Program Files (x86)\\SquareEnix\\FINAL FANTASY XIV - A Realm Reborn\\boot\\",
            "game":"C:\\Program Files (x86)\\SquareEnix\\FINAL FANTASY XIV - A Realm Reborn\\game\\"
        },
        "files": [
             "ffxivboot.exe",
             "ffxivboot64.exe",
             "ffxivlauncher.exe",
             "ffxivlauncher64.exe",
             "ffxivupdater.exe",
             "ffxivupdater64.exe"
		    ],
        "gamever_url":"https://patch-gamever.ffxiv.com/http/win32/ffxivneo_release_game/%s/%s"
    }
}
```

# Todo
- [ ] Check if launcher / game is up to date
- [ ] Maintenance check
- [ ] Setup / auto-gen config
- [ ] TestSID for auto-login

---
**This launcher is not affiliated with SQUARE ENIX CO., LTD. in any way**
