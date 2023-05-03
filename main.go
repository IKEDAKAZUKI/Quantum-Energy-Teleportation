package main

import (
	"io/ioutil"
	"strings"

	"gitlab.com/gomidi/midi/writer"
)

func main() {
	data, _ := ioutil.ReadFile("data.bits")
	lines := strings.Split(string(data), "\n")

	err := writer.WriteSMF("notes.mid", 2, func(wr *writer.SMF) error {
		for _, line := range lines {
			bits := []byte(line)
			if len(bits) != 2 {
				continue
			}
			wr.SetChannel(0)
			if bits[0] == '1' {
				writer.NoteOn(wr, 50, 100)
			}
			if bits[1] == '1' {
				writer.NoteOn(wr, 53, 100)
			}
			wr.SetDelta(120)
			if bits[0] == '1' {
				writer.NoteOff(wr, 50)
			}
			if bits[1] == '1' {
				writer.NoteOff(wr, 53)
			}
			wr.SetDelta(240)
		}

		writer.EndOfTrack(wr)

		return nil
	})
	if err != nil {
		panic(err)
	}
}
