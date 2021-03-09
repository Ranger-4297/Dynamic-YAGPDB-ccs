{{/*
        Made by Ranger (765316548516380732)

    Trigger Type: `Command`
    Trigger: `Apply`
©️ Dynamic 2021
MIT License
*/}}


{{/* Configuration values start */}}
{{$FormLink := "https://realsite.co.uk/application"}} {{/* Your servers application link */}}
{{$logchannel := 794365711614345267}} {{/* Log channel */}}
{{/* Configuration values end */}}

{{/* Only edit below if you know what you're doing (: rawr */}}

{{$icon := ""}}
{{$name := printf "%s (%d)" .Guild.Name .Guild.ID}}
{{if .Guild.Icon}}
    {{$ext := "webp"}}
    {{if eq (slice .Guild.Icon 0 2) "a_"}} {{$ext = "gif"}} {{end}}
    {{$icon = printf "https://cdn.discordapp.com/icons/%d/%s.%s" .Guild.ID .Guild.Icon $ext}}
{{end}}

{{sendMessage nil "Check your DM's for info!"}}
{{$embed := cembed
            "author" (sdict "url" (print "https://discord.com/channels/" .Guild.ID) "name" (print .Guild.Name " staff applications") "icon_url" (.User.AvatarURL "1024"))
            "description" (print  .User.Mention ",\nHello " .User.Username "\nThank you for taking interest in our servers staff team!\nPlease keep in mind that asking about the status of your application may warrant it's status as `denied`.\n[Application form](" $FormLink ")")
            "thumbnail" (sdict "url" $icon)
            "timestamp" currentTime
            "color" 4645612
            }}
{{sendDM $embed}}


{{$logembed := cembed
            "author" (sdict "icon_url" (.User.AvatarURL "1024") "name" (print .User.String " (ID " .User.ID ")"))
            "description" (print "**✉️ Application command notification**\n" .User.Mention " Has requested a staff application.")
            "thumbnail" (sdict "url" $icon)
            "timestamp" currentTime
            "color" 4645612
            }}
