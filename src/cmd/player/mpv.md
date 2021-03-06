# MPV Commands

https://mpv.io/manual/stable/#properties

https://mpv.io/manual/stable/#options

https://mpv.io/manual/stable/#list-of-input-commands

http://manpages.ubuntu.com/manpages/bionic/man1/mpv.1.html

> Most options can be set as runtime via properties as well. Just remove the leading -- from the option name. These are not documented. Only properties which do not exist as option with the same name, or which have very different behavior from the options are documented

## GET

* `path`: Path to the file being played. Example: `./06 Flying Sorcery.flac`
* `media-title`: If the currently played file has a title tag, use that. Otherwise, return the filename property. Example `Flying Sorcery`
* `file-format`: Symbolic name of the file format. In some cases, this is a comma-separated list of format names. Example `flac`
* `stream-pos`: Raw byte position in source stream. Example `4.8200567e+07`
* `stream-end`: Raw end position in bytes in source stream. Example `9.541933e+07`
* `duration`: Duration of the current file in seconds. Example `262.956885`
* `percent-pos`: Position in current file (0-100). Example `90.589849`
* `time-pos`: Position in current file in seconds. Example `23.292501`
* `time-remaining`: Remaining length of the file in seconds. Example `181.657651`
* `metadata`: Metadata key/value pairs. Example `map[album:2112 (Remastered) album_artist:Rush artist:Rush compatible_brands:M4A mp42isom compilation:0 composer:Geddy Lee, Neil Peart & Alex Lifeson creation_time:2025-01-03T14:01:48.000000Z date:1976-04-01T08:00:00Z disc:1/1 genre:Rock language:eng major_brand:M4A  media_type:1 minor_version:0 purchase_date:2017-11-23 01:11:19 sort_album:2112 (Remastered) sort_artist:Rush sort_name:2112: I. Overture, II. The Temples of Syrinx, III. Discovery, IV. Presentation, V. Oracle: The Dream, VI. Soliloquy, VII. Grand Finale title:2112: I. Overture, II. The Temples of Syrinx, III. Discovery, IV. Presentation, V. Oracle: The Dream, VI. Soliloquy, VII. Grand Finale track:1/6]`
* `audio-codec`: Audio codec selected for decoding. Example `flac (FLAC (Free Lossless Audio Codec))`
* `audio-codec-name`: Audio codec. Example: `flac`
* `audio-params`: Audio format as output by the audio decoder. Example `map[channel-count:2 channels:stereo format:s32 hr-channels:stereo samplerate:96000]`
* `seekable`: Return whether it's generally possible to seek in the current file. Example `true`
* `audio-device-list`: Return the list of discovered audio devices. Example ` [map[description:Autoselect device name:auto] map[description:DisplayPort name:coreaudio/AppleGFXHDAEngineOutputDP:0:{6D9E-7721-0002D07E}] map[description:Logitech USB Headset name:coreaudio/AppleUSBAudioEngine:Logitech USB Headset:Logitech USB Headset:14600000:2] map[description:Mac mini Speakers name:coreaudio/BuiltInSpeakerDevice] map[description:Loopback RC/Mic name:coreaudio/com.rogueamoeba.Loopback:E8FAA1C1-BD9B-4C4D-B313-7B67303CF401]]`
* `protocol-list`: List of protocol prefixes potentially recognized by the player. Example `[rtmp rtsp http https mms mmst mmsh mmshttp rtp httpproxy rtmpe rtmps rtmpt rtmpte rtmpts srtp gopher data lavf ffmpeg udp ftp tcp tls unix sftp md5 concat avdevice av file bd br bluray bdnav brnav bluraynav archive memory hex null mf edl file fd fdclose appending]`
* `decoder-list`: List of decoders supported. Example `map[codec:pcm_u8 description:PCM unsigned 8-bit driver:pcm_u8] map[codec:pcm_u16be description:PCM unsigned 16-bit big-endian driver:pcm_u16be]]`
* `demuxer-lavf-list`: List of available libavformat demuxers' names. Example `[aa aac ac3 avi mpeg mp3]`
* `mpv-version`: Return the mpv version/copyright string. Example `mpv 0.32.0`
* `ffmpeg-version`: Return the contents of the av_version_info() API call. Example `ffmpeg-version value: 4.2.3`
* `eof-reached`: Returns `true` when EOF has reached `false` otherwise. Useful to know when a song starts and ends



## SET

* `pause`: Toggles pause mode. Values `true` or `false`
* `audio-device`: Sets the audio device. `name` fully qualified is the value which should be used. Switching between audio interfaces causes a `pause` event in any case
* `speed`: Slow down or speed up playback by the factor given as parameter. Values are numbers from `0.01` to `100`
* `volume`: Sets the volume. Values are numbers from `0` to `100`

## COMMANDS

In the golang wrapper we need to use `conn.Call()` function

* `loadfile`: Changes the playing track at runtime. Example `"loadfile", filePath, "replace"`
* `seek`: Relative to current position, for examples 
`conn.Call("seek", -15, "relative")` will go backwards 15 seconds
`conn.Call("seek", 200, "exact")` will go to position at 03:20 

## Useful Stuff

This starts `mpv` player without displaying any UI and connecting to unix socket
```
mpv --no-video --input-ipc-server=/tmp/mpv_socket .
```


Shows no output to stdout
```
--really-quiet
```

Shows the minimal output to stdout
```
--quiet
```

Does not play anything and waits for a command. Not sure how does this work though
```
--track-auto-selection=no
```

Forces `mpv` not to exit (hence having a zombie process) when media ends
```
--keep-open=yes
```

This should be the startup command to fork it
```
mpv --no-video --input-ipc-server=/tmp/mpv_socket --keep-open=yes . 
```

### Listeners
```
	events, _ := conn.NewEventListener()
	go func() {
		for event := range events {
			log.Println(event.Name, event.Text, event.Data)
		}
	}()
```

## Properties

Events Dispatched by MPV

*Track Started*
Happens only once when we load a track
m.ipc.Call("observe_property", 1, "filename")

*Track Position Changed*
Happens each ~ 100ms
m.ipc.Call("observe_property", 2, "stream-pos")

*Track End Detected*
Happens only once when file stream is calculated
m.ipc.Call("observe_property", 2, "stream-end")

*Position Percent Detected*
Happens each ~ 100ms unit is percent of the song played
m.ipc.Call("observe_property", 2, "percent-pos")

*Time Position Detected*
Happens each ~ 100ms unit is seconds
m.ipc.Call("observe_property", 2, "time-pos")

*Time Position Detected*
The opposite of `time-pos` property unit is seconds
m.ipc.Call("observe_property", 2, "time-remaining")


These properties look usable, but not yet

* `seeking`
* `ao-volume`