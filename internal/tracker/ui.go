package tracker

import "fmt"

type UI struct {
	In      Input
	Out     Output
	Tracker *Tracker
}

func (u UI) Run() {
	actions := map[string]Usecase{
		"add":    AddUsecase{},
		"get":    GetUsecase{},
		"update": UpdateUsecase{},
		"delete": DeleteUsecase{},
		"find":   FindUsecase{},
	}

	for {
		u.Out.Out("select action")
		for key := range actions {
			fmt.Println(key)
		}
		selected := u.In.Get()

		if selected == "exit" {
			break
		}

		action, ok := actions[selected]
		if !ok {
			u.Out.Out("not found action")
			continue
		}

		action.Done(u.In, u.Out, u.Tracker)
	}
}
