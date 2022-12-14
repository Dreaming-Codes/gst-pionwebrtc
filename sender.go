package main

import (
	"github.com/pion/webrtc/v3"
	"github.com/tinyzimmer/go-glib/glib"
	"github.com/tinyzimmer/go-gst/gst"
	"github.com/tinyzimmer/go-gst/gst/base"
)

type MediaType int8

const (
	Audio MediaType = 0
	Video           = 1
)

type state struct {
	mediaType MediaType
	track     webrtc.TrackLocalStaticSample
}

type senderSink struct {
	// The current state of the element
	state *state
}

func (s *senderSink) ClassInit(klass *glib.ObjectClass) {
	CAT.Log(gst.LevelLog, "Initializing new senderSink class")
	class := gst.ToElementClass(klass)
	class.SetMetadata(
		"WebRTC Broadcast Engine (Internal sender)",
		"Sink/Video/Audio",
		"Internal WebRtcRedux sender",
		"Lorenzo Rizzotti <dev@dreaming.codes>",
	)
	CAT.Log(gst.LevelLog, "Adding sink pad template and properties to class")
	class.AddPadTemplate(gst.NewPadTemplate(
		"sink",
		gst.PadDirectionSink,
		gst.PadPresenceAlways,
		gst.NewAnyCaps(),
	))
	class.InstallProperties(properties)
}

func (s *senderSink) New() glib.GoObjectSubclass {
	CAT.Log(gst.LevelLog, "Initializing new senderSink object")
	return &senderSink{
		state: &state{},
	}
}

func (s *senderSink) Render(self *base.GstBaseSink, buffer *gst.Buffer) gst.FlowReturn {
	CAT.Log(gst.LevelLog, "Rendering buffer")
	return gst.FlowOK
}
