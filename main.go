package main

import (
	"fmt"
	"io/ioutil"
	"math"
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
	values["0000"] = Value{
		Name: "0000",
	}
	values["0001"] = Value{
		Name: "0001",
	}
	values["0010"] = Value{
		Name: "0010",
	}
	values["0011"] = Value{
		Name: "0011",
	}

	values["0100"] = Value{
		Name: "0100",
	}
	values["0101"] = Value{
		Name: "0101",
	}
	values["0110"] = Value{
		Name: "0110",
	}
	values["0111"] = Value{
		Name: "0111",
	}

	values["1000"] = Value{
		Name: "1000",
	}
	values["1001"] = Value{
		Name: "1001",
	}
	values["1010"] = Value{
		Name: "1010",
	}
	values["1011"] = Value{
		Name: "1011",
	}

	values["1100"] = Value{
		Name: "1100",
	}
	values["1101"] = Value{
		Name: "1101",
	}
	values["1110"] = Value{
		Name: "1110",
	}
	values["1111"] = Value{
		Name: "1111",
	}
	v := plotter.Values{}
	err = writer.WriteSMF("markov.mid", 1, func(wr *writer.SMF) error {
		context := [2]string{"00", "00"}
		for _, line := range lines {
			bits := []byte(line)
			if len(bits) != 2 {
				continue
			}
			i := float64((bits[1] - '0') + 2*(bits[0]-'0'))
			v = append(v, i)
			c := ""
			for _, value := range context {
				c += value
			}
			v := values[c]
			v.Values = append(values[c].Values, i)
			values[c] = v
			last := context[len(context)-1]
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
			copy(context[:len(context)-1], context[1:])
			context[len(context)-1] = line
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
	h := make([]float64, 4)
	sum, total := make([]int, 4), 0
	for _, v := range v {
		sum[int(v)]++
		total++
	}
	for key, value := range sum {
		h[key] = float64(value) / float64(total)
	}
	for _, value := range sorted {
		key := value.Name
		sum, total := make([]int, 4), 0
		for _, v := range value.Values {
			sum[int(v)]++
			total++
		}
		difference := 0.0
		for key, value := range h {
			difference += math.Abs(value - float64(sum[key])/float64(total))
		}
		difference /= float64(len(h))
		fmt.Println(key, difference, len(value.Values))
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
	fmt.Println("total")
	for key, value := range h {
		fmt.Printf("%d %f\n", key, value)
	}
}
