# ffmpeg_overlay

Allow to automate the process of overlaying videos. (overlay video at bottom right corner of the background video)

The duration of the output video is the duraction of the background video.

## Usage

```
$ go build ffmpeg_overlay.go
$ ./ffmpeg_overlay background_video overlay_video //outputs out.chosen_video_codec
```

The overlay video duration must be equal or longer than the background video. If you want to synchronize the videos, just make sure to cut the overlay video where you want to start according to the background video, the programm handles for you the remaining end synchronization.

If your files have blank spaces, just make sure to put them like `directory_path/{my file}`.

You can concatenate overlayed videos like:

```
$ ./ffmpeg_overlay [background_video1,background_video2] [overlay_video1,overlay_video2] 
```

This will create an `out.mp3` as concatenated videos from overlayed `background_video1` and `overlay_video1`, and overlayed `background_video2` and `overlay_video2`

## Versions

Built with `go1.24.2` and `ffmpeg/ffprobe7.1.1`

# Web utils

This directory provides an example of an `POST` request to upload file to a webserver and an example of a `GET` request.


