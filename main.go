package main

import (
	"io/ioutil"
	"strings"

	"gitlab.com/gomidi/midi/writer"
)

func main() {
	data, _ := ioutil.ReadFile("data.bits")
	lines := strings.Split(string(data), "\n")

	err := writer.WriteSMF("notes.mid", 1, func(wr *writer.SMF) error {
		for _, line := range lines {
			bits := []byte(line)
			if len(bits) != 2 {
				continue
			}
			wr.SetChannel(0)
			if bits[0] == '1' {
				writer.NoteOn(wr, 48, 100)
			}
			if bits[1] == '1' {
				writer.NoteOn(wr, 52, 100)
			}
			wr.SetDelta(120)
			if bits[0] == '1' {
				writer.NoteOff(wr, 48)
			}
			if bits[1] == '1' {
				writer.NoteOff(wr, 52)
			}
			wr.SetDelta(240)
		}

		writer.EndOfTrack(wr)

		return nil
	})
	if err != nil {
		panic(err)
	}

	err = writer.WriteSMF("markov.mid", 1, func(wr *writer.SMF) error {
		last := ""
		for _, line := range lines {
			bits := []byte(line)
			if len(bits) != 2 {
				continue
			}
			if last == "" || last == "00" {
				wr.SetDelta(120)
				wr.SetDelta(240)
			} else if last == "01" {
				wr.SetChannel(0)
				if bits[0] == '1' {
					writer.NoteOn(wr, 48, 100)
				}
				if bits[1] == '1' {
					writer.NoteOn(wr, 52, 100)
				}
				wr.SetDelta(120)
				if bits[0] == '1' {
					writer.NoteOff(wr, 48)
				}
				if bits[1] == '1' {
					writer.NoteOff(wr, 52)
				}
				wr.SetDelta(240)
			} else if last == "10" {
				wr.SetChannel(0)
				if bits[0] == '1' {
					writer.NoteOn(wr, 53, 100)
				}
				if bits[1] == '1' {
					writer.NoteOn(wr, 57, 100)
				}
				wr.SetDelta(120)
				if bits[0] == '1' {
					writer.NoteOff(wr, 53)
				}
				if bits[1] == '1' {
					writer.NoteOff(wr, 57)
				}
				wr.SetDelta(240)
			} else if last == "11" {
				wr.SetChannel(0)
				if bits[0] == '1' {
					writer.NoteOn(wr, 55, 100)
				}
				if bits[1] == '1' {
					writer.NoteOn(wr, 59, 100)
				}
				wr.SetDelta(120)
				if bits[0] == '1' {
					writer.NoteOff(wr, 55)
				}
				if bits[1] == '1' {
					writer.NoteOff(wr, 59)
				}
				wr.SetDelta(240)
			}
			last = line
		}

		writer.EndOfTrack(wr)

		return nil
	})
	if err != nil {
		panic(err)
	}
}
