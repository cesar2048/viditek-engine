package main

/*
// VideoProcessor is any implementation that makes transformations on videos
type VideoProcessor interface {
	ConvertVideo(*Video, VideoMeta) Video
}
*/

type AwsVideo string

func (AwsVideo) ConvertVideo(v *Video, meta VideoMeta) Video {
	newVid := Video{}
	newVid.Title = "Modified " + v.Title
	newVid.URL = "https://domain.com/path/to/resource&params=to&authenticate"
	newVid.properties = meta

	return newVid
}
