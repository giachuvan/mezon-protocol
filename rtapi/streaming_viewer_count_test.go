package rtapi

import (
	"bytes"
	"testing"

	"google.golang.org/protobuf/proto"
)

// Envelope field 101 is length-delimited: tag (101<<3)|2 = 810 = varint 0xAA 0x06.
var streamingViewerCountGolden = []byte{0xAA, 0x06, 0x06, 0x08, 0x01, 0x10, 0x02, 0x18, 0x03}

func TestStreamingViewerCountEventWire(t *testing.T) {
	env := &Envelope{Message: &Envelope_StreamingViewerCountEvent{StreamingViewerCountEvent: &StreamingViewerCountEvent{
		ClanId: 1, ChannelId: 2, ViewerCount: 3,
	}}}

	vt, err := env.MarshalVT()
	if err != nil {
		t.Fatalf("MarshalVT: %v", err)
	}
	if !bytes.Equal(vt, streamingViewerCountGolden) {
		t.Fatalf("MarshalVT = % X, want % X", vt, streamingViewerCountGolden)
	}

	std, err := proto.Marshal(env)
	if err != nil {
		t.Fatalf("proto.Marshal: %v", err)
	}
	if !bytes.Equal(std, streamingViewerCountGolden) {
		t.Fatalf("proto.Marshal = % X, want % X", std, streamingViewerCountGolden)
	}

	var got Envelope
	if err := got.UnmarshalVT(streamingViewerCountGolden); err != nil {
		t.Fatalf("UnmarshalVT: %v", err)
	}
	if !proto.Equal(&got, env) {
		t.Fatalf("round trip = %v, want %v", &got, env)
	}
}
