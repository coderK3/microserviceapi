package main

import (
	"fmt"

	"github.com/looplab/fsm"
)

type Content struct {
	User string
	FSM  *fsm.FSM
}

func NewUser(user string) *Content {
	urs := &Content{
		User: user,
	}

	urs.FSM = fsm.NewFSM(
		"unread",
		fsm.Events{
			{Name: "read", Src: []string{"unread"}, Dst: "read"},
			{Name: "uread", Src: []string{"read"}, Dst: "unread"},
		},
		fsm.Callbacks{
			"enter_state": func(e *fsm.Event) {
				urs.enterState(e)
			},
		},
	)
	return urs
}
func (urs *Content) enterState(e *fsm.Event) {
	fmt.Printf(" %s is %s \n", urs.User, e.Dst)

}
func main() {
	users := NewUser("test")

	err := users.FSM.Event("read")
	if err != nil {
		fmt.Println(err)
	}
	err = users.FSM.Event("uread")
	if err != nil {
		fmt.Println(err)
	}
}
