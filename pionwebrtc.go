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

func main() {}

type gobin struct{}

func (g *gobin) New() glib.GoObjectSubclass { return &gobin{} }

func (g *gobin) ClassInit(klass *glib.ObjectClass) {}
