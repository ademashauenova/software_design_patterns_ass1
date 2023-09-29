package main

import "fmt"

// Strategy Interface
type ArtisticToolAlgo interface {
	Use()
}

// Concrete Paint Strategy
type PaintTool struct{}

func (p *PaintTool) Use() {
	fmt.Println("Switched to the paint tool to draw")
}

// Concrete Smudge Strategy
type SmudgeTool struct{}

func (s *SmudgeTool) Use() {
	fmt.Println("Switched to the smudge tool to blend colors")
}

// Concrete Erase Strategy
type EraseTool struct{}

func (e *EraseTool) Use() {
	fmt.Println("Switched to the erase tool to remove content")
}

// Context
type DigitalArtistryApp struct {
	ActiveTool ArtisticToolAlgo
}

func (app *DigitalArtistryApp) SetActiveTool(tool ArtisticToolAlgo) {
	app.ActiveTool = tool
}

func (app *DigitalArtistryApp) UseTool() {
	fmt.Println("Digital artistry app is active")
	app.ActiveTool.Use()
}

func main() {
	app := &DigitalArtistryApp{}

	paint := &PaintTool{}
	app.SetActiveTool(paint)
	app.UseTool()

	smudge := &SmudgeTool{}
	app.SetActiveTool(smudge)
	app.UseTool()

	erase := &EraseTool{}
	app.SetActiveTool(erase)
	app.UseTool()
}
