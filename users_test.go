// Copyright (c) 2016 The gotick developers. All rights reserved.
// Project site: https://github.com/questrail/gotick
// Use of this source code is governed by a MIT-style license that
// can be found in the LICENSE.txt file for the project.

package gotick

import (
	"fmt"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestGetUsers(t *testing.T) {
	Convey("Given I want to get the users", t, func() {
		Convey("When the /users.json URI receives a GET method\n"+
			"And the header contains valid authorization",
			func() {
				var fake FakeSession
				data, _ := GetUsers(&fake)
				Convey("It should return five users", func() {
					So(len(data), ShouldEqual, 5)
				})
				Convey("The first user ID should be 444444", func() {
					So(data[0].ID, ShouldEqual, 444444)
				})
				Convey("The second user ID should be 444445", func() {
					So(data[1].ID, ShouldEqual, 444445)
				})
				lastNames := []string{"Smith", "Sowell", "von Hayek", "Graham", "Friedman"}
				ordinals := []string{"first", "second", "third", "fourth", "fifth"}
				for index, lastName := range lastNames {
					ordinal := ordinals[index]
					conveyance := fmt.Sprintf("The %s user's last name should be %s", ordinal, lastName)
					Convey(conveyance, func() {
						So(data[index].LastName, ShouldEqual, lastName)
					})
				}
			})
	})
}
