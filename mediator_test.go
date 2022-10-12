package sugar

import (
	"fmt"
	"strings"
	"testing"
)

type CDDriver struct {
	Data string
}

func (c *CDDriver) GetTransData() any {
	return c.Data
}

func (c *CDDriver) Process(a any) {
	c.Data = "music,image"
	fmt.Printf("CDDriver: reading data %s\n", c.Data)
	GetMediator().Changed(c)
}

type CPU struct {
	Video string
	Sound string
}

func (c *CPU) Process(data any) {
	sp := strings.Split(data.(string), ",")
	c.Sound = sp[0]
	c.Video = sp[1]

	fmt.Printf("CPU: split data with Sound %s, Video %s\n", c.Sound, c.Video)
	GetMediator().Changed(c)
}

func (c *CPU) GetTransData() any {
	return c
}

type VideoCard struct {
	Data string
}

func (v *VideoCard) Process(data any) {
	v.Data = data.(*CPU).Video
	fmt.Printf("VideoCard: display %s\n", v.Data)
	GetMediator().Changed(v)
}

func (v *VideoCard) GetTransData() any {
	return v.Data
}

type SoundCard struct {
	Data string
}

func (s *SoundCard) Process(data any) {
	s.Data = data.(*CPU).Sound
	fmt.Printf("SoundCard: play %s\n", s.Data)
	GetMediator().Changed(s)
}

func (s *SoundCard) GetTransData() any {
	return s.Data
}

func TestMediator(t *testing.T) {
	mediator := GetMediator()
	cd := &CDDriver{Data: "music,image"}
	cpu := &CPU{}
	mediator.RegisterMember(cd, cpu)
	mediator.RegisterMember(cpu, &VideoCard{})
	mediator.RegisterMember(cpu, &SoundCard{})
	cd.Process(nil)
}
