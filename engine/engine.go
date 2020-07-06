package main

// VideoCollection represents a collection of the original video and derivatives of it
type VideoCollection struct {
	Original    Video
	Derivatives map[string]Video
}

// GeneratePreview generates a derivative video from the original
func (v *VideoCollection) GeneratePreview(processor VideoProcessor, derivativeName string, newProperties VideoMeta) {
	newVideo := processor.ConvertVideo(&v.Original, newProperties)
	v.Derivatives[derivativeName] = newVideo
}

// ----- Structs and interfaces -----

// VideoProcessor is any implementation that makes transformations on videos
type VideoProcessor interface {
	ConvertVideo(*Video, VideoMeta) Video
}

// VideoMeta holds properties about dimensions
type VideoMeta struct {
	Width  int
	Height int
}

// Video represents a single video
type Video struct {
	Title      string
	URL        string
	properties VideoMeta
}
