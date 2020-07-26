# Messages

## List of Message Broker Subjects/Actions

| Message | Handler | Sender |Action |
| ------------- |:-------------:| :-------------:| -----|
| MESSAGE_ARTWORK_SCAN| artwork | audioscan  | looks up for artwork and makes a copy to `artwork` directory |
| MESSAGE_AUDIO_SCAN | audioscan | restapi | scans a directory for media new or updated |
| MESSAGE_MEDIA_INFO| mediainfo | audioscan | extracts information from a media file |
| MESSAGE_MEDIA_UPDATE_ARTWORK | mediainfo | artwork | updates artwork for a media in db |
| MESSAGE_MEDIA_UPDATE_SHA | mediainfo | audioscan | updates sha information for a media in db |
| MESSAGE_TAG_UPDATE | artwork | N/A | writes idv2 tags to media file |
| MESSAGE_MEDIA_UPDATE | mediainfo | N/A | writes media information to db |
| MESSAGE_IPC | player | restapi | performs different actions based on the command provided |
| MESSAGE_REST_API_READY | restapi | artwork, audioscan, mediainfo | asks for a ping to know if restapi is up |
| MESSAGE_MPV_EOF_REACHED | player | player | mpv internal event (file has finished playing |
| MESSAGE_MPV_PERCENT_POS | player | player | mpv internal event (player has changed position) |
| MESSAGE_MPV_TIME_POS | player | player | mpv internal event (player has changed position) |
| MESSAGE_MPV_TIME_REMAINING | player | player | mpv internal event (player has changed position) |
| MESSAGE_EVENT_ON_TRACK_STARTED | player | N/A | mpv event when track starts playing |
| MESSAGE_EVENT_ON_TRACK_FINISHED | N/A | N/A | mpv event when track ends playing |
| MESSAGE_EVENT_ON_TRACK_PLAY_COUNT_INCREASE | mediainfo | player | triggered when player determines the song has been played  
| MESSAGE_EVENT_ON_TRACK_SKIP_COUNT_INCREASE | mediainfo | player | triggered when player determines the song has been skipped