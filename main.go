package main

import (
	"fmt"
	"io/ioutil"
	"strings"

	"gitlab.com/gomidi/midi/writer"
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
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

	values := make(map[string]plotter.Values)
	v := plotter.Values{}
	err = writer.WriteSMF("markov.mid", 1, func(wr *writer.SMF) error {
		last := ""
		for _, line := range lines {
			bits := []byte(line)
			if len(bits) != 2 {
				continue
			}
			i := float64((bits[1] - '0') + 2*(bits[0]-'0'))
			v = append(v, i)
			if last == "" {
				values["00"] = append(values["00"], i)
			} else {
				values[last] = append(values[last], i)
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

	for key, value := range values {
		p := plot.New()
		if err != nil {
			panic(err)
		}
		p.Title.Text = "histogram plot"

		histogram, err := plotter.NewHist(value, 4)
		if err != nil {
			panic(err)
		}
		p.Add(histogram)

		err = p.Save(8*vg.Inch, 8*vg.Inch, fmt.Sprintf("%s_historgram.png", key))
		if err != nil {
			panic(err)
		}
	}
	p := plot.New()
	if err != nil {
		panic(err)
	}
	p.Title.Text = "histogram plot"

	histogram, err := plotter.NewHist(v, 4)
	if err != nil {
		panic(err)
	}
	p.Add(histogram)

	err = p.Save(8*vg.Inch, 8*vg.Inch, "historgram.png")
	if err != nil {
		panic(err)
	}
}
