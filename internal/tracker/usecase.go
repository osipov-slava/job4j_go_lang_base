package tracker

import (
	"fmt"

	"github.com/google/uuid"
)

type Usecase interface {
	Done(in Input, out Output, tracker *Tracker)
}

type AddUsecase struct{}

func (u AddUsecase) Done(in Input, out Output, tracker *Tracker) {
	out.Out("enter name:")
	name := in.Get()
	id := uuid.New().String()
	tracker.AddItem(Item{Name: name, ID: id})
}

type GetUsecase struct{}

func (u GetUsecase) Done(_ Input, out Output, tracker *Tracker) {
	for _, item := range tracker.GetItems() {
		out.Out(item.toString())
	}
}

type UpdateUsecase struct{}

func (u UpdateUsecase) Done(in Input, out Output, tracker *Tracker) {
	out.Out("enter ID:")
	id := in.Get()
	out.Out("enter NEW name:")
	name := in.Get()
	if tracker.UpdateItem(id, name) {
		fmt.Println("Update success")
	} else {
		fmt.Println("Update failed")
	}
}

type DeleteUsecase struct{}

func (u DeleteUsecase) Done(in Input, out Output, tracker *Tracker) {
	out.Out("enter ID:")
	id := in.Get()
	if tracker.DeleteItem(id) {
		fmt.Println("Delete success")
	} else {
		fmt.Println("Delete failed")
	}
}

type FindUsecase struct{}

func (u FindUsecase) Done(in Input, out Output, tracker *Tracker) {
	out.Out("enter name(or fragment):")
	name := in.Get()
	items := tracker.FindItemsByFragment(name)
	if len(items) == 0 {
		fmt.Println("No items found")
	}
	for _, item := range items {
		out.Out(item.toString())
	}
}
