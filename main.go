package main

import (
	"gitlab.com/gomidi/midi/writer"
)

func main() {
	err := writer.WriteSMF("notes.mid", 2, func(wr *writer.SMF) error {
		wr.SetChannel(0)
		writer.NoteOn(wr, 50, 100)
		wr.SetDelta(120)
		writer.NoteOff(wr, 50)

		wr.SetDelta(240)
		writer.NoteOn(wr, 53, 100)
		wr.SetDelta(20)
		writer.NoteOff(wr, 53)
		writer.EndOfTrack(wr)

		wr.SetChannel(1)
		writer.NoteOn(wr, 50, 100)
		wr.SetDelta(60)
		writer.NoteOff(wr, 50)
		writer.EndOfTrack(wr)
		return nil
	})
	if err != nil {
		panic(err)
	}
}
