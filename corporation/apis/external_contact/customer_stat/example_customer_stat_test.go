package customer_stat_test

import (
	"fmt"

	"github.com/cvblood/qywxapi/corporation"
	"github.com/cvblood/qywxapi/corporation/apis/external_contact/customer_stat"
)

func ExampleGetUserBehaviorData() {
	var ctx *corporation.App

	payload := []byte("{}")
	resp, err := customer_stat.GetUserBehaviorData(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleStatistic() {
	var ctx *corporation.App

	payload := []byte("{}")
	resp, err := customer_stat.Statistic(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleStatisticGroupByDay() {
	var ctx *corporation.App

	payload := []byte("{}")
	resp, err := customer_stat.StatisticGroupByDay(ctx, payload)

	fmt.Println(resp, err)
}
