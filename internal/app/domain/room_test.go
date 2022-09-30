package domain_test

import (
	"chatroom-demo/internal/app/domain"
	"chatroom-demo/internal/app/domain/vo"
	"fmt"
	"io/ioutil"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
	"gopkg.in/yaml.v2"
)

var roomtestcases []roomfixture

type roomfixture struct {
	Input struct {
		Data struct {
			Name string `yaml:"name"`
			Type string `yaml:"type"`
		} `yaml:"data"`
	} `yaml:"input"`
	Expect struct {
		Error bool `yaml:"error"`
		Data  struct {
			ID       string       `yaml:"id"`
			Code     string       `yaml:"code"`
			Name     string       `yaml:"name"`
			Type     string       `yaml:"type"`
			Messages []vo.Message `yaml:"messages"`
		} `yaml:"data"`
	} `yaml:"expect"`
}

func setupRoom(t *testing.T) {
	initTestCaseRoom(t)
}

func initTestCaseRoom(t *testing.T) {
	file, err := ioutil.ReadFile("../../../testdata/fixture/domainroom.fixture.yaml")
	if err != nil {
		panic(err)
	}
	err = yaml.Unmarshal([]byte(file), &roomtestcases)
	if err != nil {
		panic(err)
	}
}

func TestNewRoom(t *testing.T) {
	setupRoom(t)

	Convey("NewRoom", t, func() {
		for _, tc := range roomtestcases {
			if !tc.Expect.Error {
				Convey(fmt.Sprintf("given a valid input: %v", tc.Input), func() {
					data, err := domain.NewRoom(tc.Input.Data.Name, tc.Input.Data.Type)

					Convey("should return a new room with expected value", func() {
						So(err, ShouldBeNil)
						So(data.Name, ShouldEqual, tc.Expect.Data.Name)
						So(data.Type, ShouldEqual, tc.Expect.Data.Type)
						So(data.ID, ShouldNotBeZeroValue)
						So(data.Code, ShouldNotBeZeroValue)
						So(data.Messages, ShouldBeEmpty)
					})
				})
			} else {
				Convey(fmt.Sprintf("given an invalid input: %v", tc.Input), func() {
					_, err := domain.NewRoom(tc.Input.Data.Name, tc.Input.Data.Type)

					Convey("should return error on domain.NewRoom", func() {
						So(err, ShouldNotBeNil)
					})
				})
			}
		}
	})
}
