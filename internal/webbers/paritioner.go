package webbers

import (
	"strings"
)

// Takes in lines of text and sorts them into regions representing words.
type Partitioner struct {
	Regions []Region
	region  Region
}

func NewPartitioner() Partitioner {
	return Partitioner{
		region:  NewEmptyRegion(),
		Regions: make([]Region, 0),
	}
}

func (p *Partitioner) ReadLine(i int, line string) {
	//fmt.Printf("%d: %s\n", i, line)
	// Open new region if none active
	if p.region.IsNul() && line != "" && line != "\n" {
		p.region = NewOpenRegion(i)
	} else {
		// If we hit a new region, close old one and open a new one
		if strings.ToUpper(line) == line && line != "" && line != "\n" {
			p.region.Close(i - 1)
			p.Regions = append(p.Regions, p.region)
			p.region = NewOpenRegion(i)
		}
	}
	// Add the line to current region
	p.region.AddLine(line)
}

// Called when no more lines are to be read.
func (p *Partitioner) FinishedReading(i int) {
	p.region.End = i
	p.Regions = append(p.Regions, p.region)
}
