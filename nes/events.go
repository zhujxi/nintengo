package nes

import "fmt"

type Event interface {
	Process(nes *NES)
}

type FrameEvent struct {
	colors []uint8
}

func (e *FrameEvent) Process(nes *NES) {
	if nes.state != Running {
		return
	}

	if nes.recorder != nil {
		nes.recorder.Input() <- e.colors
	}

	nes.video.Input() <- e.colors
}

type ControllerEvent struct {
	controller int
	down       bool
	button     Button
}

func (e *ControllerEvent) Process(nes *NES) {
	if nes.state != Running {
		return
	}

	if e.down {
		nes.controllers.KeyDown(e.controller, e.button)
	} else {
		nes.controllers.KeyUp(e.controller, e.button)
	}
}

type PauseEvent struct{}

func (e *PauseEvent) Process(nes *NES) {
	switch nes.state {
	case Running:
		nes.state = Paused
		nes.paused <- 1
	case Paused:
		nes.state = Running
		nes.paused <- 0
	}
}

type ResetEvent struct{}

func (e *ResetEvent) Process(nes *NES) {
	if nes.state != Running {
		return
	}

	nes.Reset()
}

type RecordEvent struct{}

func (e *RecordEvent) Process(nes *NES) {
	if nes.recorder != nil {
		nes.recorder.Record()
	}
}

type StopEvent struct{}

func (e *StopEvent) Process(nes *NES) {
	if nes.recorder != nil {
		nes.recorder.Stop()
	}
}

type QuitEvent struct{}

func (e *QuitEvent) Process(nes *NES) {
	nes.state = Quitting
}

type SaveEvent struct{}

func (e *SaveEvent) Process(nes *NES) {
	nes.SaveState()
}

type LoadEvent struct{}

func (e *LoadEvent) Process(nes *NES) {
	nes.LoadState()
}

type ShowBackgroundEvent struct{}

func (e *ShowBackgroundEvent) Process(nes *NES) {
	nes.ppu.ShowBackground = !nes.ppu.ShowBackground
	fmt.Println("*** Toggling show background = ", nes.ppu.ShowBackground)
}

type ShowSpritesEvent struct{}

func (e *ShowSpritesEvent) Process(nes *NES) {
	nes.ppu.ShowSprites = !nes.ppu.ShowSprites
	fmt.Println("*** Toggling show sprites = ", nes.ppu.ShowSprites)
}

type CPUDecodeEvent struct{}

func (e *CPUDecodeEvent) Process(nes *NES) {
	fmt.Println("*** Toggling CPU decode = ", nes.cpu.ToggleDecode())
}

type PPUDecodeEvent struct{}

func (e *PPUDecodeEvent) Process(nes *NES) {
	fmt.Println("*** Toggling PPU decode = ", nes.ppu.ToggleDecode())
}

type FastForwardEvent struct{}

func (e *FastForwardEvent) Process(nes *NES) {
	nes.fps.SetRate(DEFAULT_FPS * 2.00)
	fmt.Println("*** Setting fps to fast forward (2x)")
}

type FPS100Event struct{}

func (e *FPS100Event) Process(nes *NES) {
	nes.fps.SetRate(DEFAULT_FPS * 1.00)
	fmt.Println("*** Setting fps to 4/4")
}

type FPS75Event struct{}

func (e *FPS75Event) Process(nes *NES) {
	nes.fps.SetRate(DEFAULT_FPS * 0.75)
	fmt.Println("*** Setting fps to 3/4")
}

type FPS50Event struct{}

func (e *FPS50Event) Process(nes *NES) {
	nes.fps.SetRate(DEFAULT_FPS * 0.50)
	fmt.Println("*** Setting fps to 2/4")
}

type FPS25Event struct{}

func (e *FPS25Event) Process(nes *NES) {
	nes.fps.SetRate(DEFAULT_FPS * 0.25)
	fmt.Println("*** Setting fps to 1/4")
}

type SavePatternTablesEvent struct{}

func (e *SavePatternTablesEvent) Process(nes *NES) {
	fmt.Println("*** Saving PPU pattern tables")
	nes.ppu.SavePatternTables()
}