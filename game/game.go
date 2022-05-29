package game

import (
	"fmt"
	"sync"

	"github.com/gdamore/tcell"
)

type Config struct {
	Framerate int
}

var (
	_frame      = 0
	_isQuitting = false
	_mutex      = sync.Mutex{}
	_box        = Box{
		size: Cell{x: 1, y: 1},
	}
)

func Run(config Config) error {
	s, err := tcell.NewScreen()
	if err != nil {
		return err
	}

	if err := s.Init(); err != nil {
		return err
	}

	for !_isQuitting {
		//step()
		//time.Sleep(time.Second / time.Duration(config.Framerate))

		drawWhiteBox(s, _box)
		s.Show()

		event := s.PollEvent()

		quit := func() {
			s.Fini()
			_isQuitting = true
		}

		switch event := event.(type) {
		case *tcell.EventKey:
			switch event.Key() {
			case tcell.KeyCtrlC, tcell.KeyESC:
				quit()
			case tcell.KeyCtrlD:
				s.Clear()
			case tcell.KeyUp:
				_box = _box.Move(0, -1)
			case tcell.KeyDown:
				_box = _box.Move(0, 1)
			case tcell.KeyLeft:
				_box = _box.Move(-1, 0)
			case tcell.KeyRight:
				_box = _box.Move(1, 0)
			}
		}
	}

	return nil
}

func step() {
	if !_mutex.TryLock() {
		return
	}

	_frame++
	fmt.Printf("Game step %v\n", _frame)
	_mutex.Unlock()
}
