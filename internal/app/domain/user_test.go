package domain_test

import (
	"chatroom-demo/internal/app/domain"
	"fmt"
	"io/ioutil"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
	"gopkg.in/yaml.v2"
)

var usertestcases []userfixture

type userfixture struct {
	Input struct {
		Data struct {
			Name string `yaml:"name"`
		} `yaml:"data"`
	} `yaml:"input"`
	Expect struct {
		Error bool `yaml:"error"`
		Data  struct {
			ID   string `yaml:"id"`
			Name string `yaml:"name"`
		} `yaml:"data"`
	} `yaml:"expect"`
}

func setupUser(t *testing.T) {
	initTestCaseUser(t)
}

func initTestCaseUser(t *testing.T) {
	file, err := ioutil.ReadFile("../../../testdata/fixture/domainuser.fixture.yaml")
	if err != nil {
		panic(err)
	}
	err = yaml.Unmarshal([]byte(file), &usertestcases)
	if err != nil {
		panic(err)
	}
}

func TestNewUser(t *testing.T) {
	setupUser(t)

	Convey("NewUser", t, func() {
		for _, tc := range usertestcases {
			if !tc.Expect.Error {
				Convey(fmt.Sprintf("given a valid input: %v", tc.Input), func() {
					data, err := domain.NewUser(tc.Input.Data.Name)

					Convey("should return a new room with expected value", func() {
						So(err, ShouldBeNil)
						So(data.Name, ShouldEqual, tc.Expect.Data.Name)
						So(data.ID, ShouldNotBeZeroValue)
					})
				})
			} else {
				Convey(fmt.Sprintf("given an invalid input: %v", tc.Input), func() {
					_, err := domain.NewUser(tc.Input.Data.Name)

					Convey("should return error on domain.NewRoom", func() {
						So(err, ShouldNotBeNil)
					})
				})
			}
		}
	})
}
