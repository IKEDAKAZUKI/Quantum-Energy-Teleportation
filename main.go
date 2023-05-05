package main

import (
	"fmt"
	"io/ioutil"
	"sort"
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

	type Value struct {
		Name   string
		Values plotter.Values
	}
	values := make(map[string]Value)
	values["00"] = Value{
		Name: "00",
	}
	values["01"] = Value{
		Name: "01",
	}
	values["10"] = Value{
		Name: "10",
	}
	values["11"] = Value{
		Name: "11",
	}
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
				v := values["00"]
				v.Values = append(values["00"].Values, i)
				values["00"] = v
			} else {
				v := values[last]
				v.Values = append(values[last].Values, i)
				values[last] = v
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

		histogram, err := plotter.NewHist(value.Values, 4)
		if err != nil {
			panic(err)
		}
		p.Add(histogram)

		err = p.Save(8*vg.Inch, 8*vg.Inch, fmt.Sprintf("%s_historgram.png", key))
		if err != nil {
			panic(err)
		}
	}
	sorted := make([]Value, 0, 8)
	for _, value := range values {
		sorted = append(sorted, value)
	}
	sort.Slice(sorted, func(i, j int) bool {
		return sorted[i].Name < sorted[j].Name
	})
	for _, value := range sorted {
		key := value.Name
		sum, total := make([]int, 4), 0
		for _, v := range value.Values {
			sum[int(v)]++
			total++
		}
		fmt.Println(key)
		for key, value := range sum {
			fmt.Printf("%d %f\n", key, float64(value)/float64(total))
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
	sum, total := make([]int, 4), 0
	for _, v := range v {
		sum[int(v)]++
		total++
	}
	fmt.Println("total")
	for key, value := range sum {
		fmt.Printf("%d %f\n", key, float64(value)/float64(total))
	}
}
