package calendar_test

import (
	"fmt"

	"github.com/cvblood/qywxapi/corporation"
	"github.com/cvblood/qywxapi/corporation/apis/efficiency/calendar"
)

func ExampleCalendarAdd() {
	var ctx *corporation.App

	payload := []byte("{}")
	resp, err := calendar.CalendarAdd(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleCalendarUpdate() {
	var ctx *corporation.App

	payload := []byte("{}")
	resp, err := calendar.CalendarUpdate(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleCalendarGet() {
	var ctx *corporation.App

	payload := []byte("{}")
	resp, err := calendar.CalendarGet(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleCalendarDel() {
	var ctx *corporation.App

	payload := []byte("{}")
	resp, err := calendar.CalendarDel(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleScheduleAdd() {
	var ctx *corporation.App

	payload := []byte("{}")
	resp, err := calendar.ScheduleAdd(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleScheduleUpdate() {
	var ctx *corporation.App

	payload := []byte("{}")
	resp, err := calendar.ScheduleUpdate(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleScheduleGet() {
	var ctx *corporation.App

	payload := []byte("{}")
	resp, err := calendar.ScheduleGet(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleScheduleDel() {
	var ctx *corporation.App

	payload := []byte("{}")
	resp, err := calendar.ScheduleDel(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleScheduleGetByCalendar() {
	var ctx *corporation.App

	payload := []byte("{}")
	resp, err := calendar.ScheduleGetByCalendar(ctx, payload)

	fmt.Println(resp, err)
}
