package main

import (
	"github.com/pion/webrtc/v3"
	"github.com/tinyzimmer/go-glib/glib"
	"github.com/tinyzimmer/go-gst/gst"
)

var (
	peerConnectionConfig = webrtc.Configuration{}
)

var CAT = gst.NewDebugCategory(
	"pionwebrtc",
	gst.DebugColorNone,
	"PionWebRTC Element",
)

var properties = []*glib.ParamSpec{}

func main() {}

type gobin struct{}

func (g *gobin) New() glib.GoObjectSubclass { return &gobin{} }

func (g *gobin) ClassInit(klass *glib.ObjectClass) {
	class := gst.ToElementClass(klass)
	class.SetMetadata(
		"PionWebRTC Element",
		"Bin/WebRTC",
		"WebRTC Video and Audio Transmitter",
		"Lorenzo Rizzotti <dev@reaming.codes>",
	)
	class.AddPadTemplate(gst.NewPadTemplate(
		"src",
		gst.PadDirectionSource,
		gst.PadPresenceSometimes,
		gst.NewAnyCaps(),
	))
	class.AddPadTemplate(gst.NewPadTemplate(
		"sink",
		gst.PadDirectionSink,
		gst.PadPresenceRequest,
		gst.NewAnyCaps(),
	))
	class.InstallProperties(properties)
}
