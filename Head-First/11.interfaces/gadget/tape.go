package gadget

import "fmt"

type Taper struct {
	Batteries string
}

func (t Taper) Play(song string) {
	fmt.Println("playing", song)
}

func (t Taper) Stop() {
	fmt.Println("stopped!")
}

type Recorder struct {
	Microphones int
}

func (r Recorder) Play(song string) {
	fmt.Println("playing", song)
}

func (r Recorder) Stop() {
	fmt.Println("stopped!")
}

func (r Recorder) Record() {
	fmt.Println("Recording")
}
