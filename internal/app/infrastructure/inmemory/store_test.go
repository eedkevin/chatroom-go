package inmemory_test

import (
	"chatroom-demo/internal/app/infrastructure/inmemory"
	"fmt"
	"io/ioutil"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
	"gopkg.in/yaml.v2"
)

var store *inmemory.InmemoryStorage
var testcases []fixture

type fixture struct {
	Input struct {
		Data struct {
			ID     string      `yaml:"id"`
			Object interface{} `yaml:"object"`
		} `yaml:"data"`
	} `yaml:"input"`
	Expect struct {
		Error bool `yaml:"error"`
		Data  struct {
			ID     string      `yaml:"id"`
			Object interface{} `yaml:"object"`
		} `yaml:"data"`
	}
}

func setup(t *testing.T) {
	store = inmemory.NewStorage()
	initTestCase(t)
}

func initTestCase(t *testing.T) {
	file, err := ioutil.ReadFile("../../../../testdata/fixture/inmemoerystore.fixture.yaml")
	if err != nil {
		panic(err)
	}
	err = yaml.Unmarshal([]byte(file), &testcases)
	if err != nil {
		panic(err)
	}
}

func TestInmemoryStoreBasic(t *testing.T) {
	setup(t)

	Convey("Save", t, func() {
		for _, tc := range testcases {
			if !tc.Expect.Error {
				Convey(fmt.Sprintf("given a valid input: %v", tc.Input), func() {
					err := store.Save(tc.Input.Data.ID, tc.Input.Data.Object)

					Convey("should return no error on store.Save", func() {
						So(err, ShouldBeNil)
					})
				})
			} else {
				Convey(fmt.Sprintf("given an invalid input: %v", tc.Input), func() {
					_, err := store.Get(tc.Input.Data.ID)
					Convey("should return error on store.Save", func() {
						So(err, ShouldNotBeNil)
					})
				})
			}
		}
	})

	Convey("Get", t, func() {
		for _, tc := range testcases {
			if !tc.Expect.Error {
				Convey(fmt.Sprintf("given a valid input: %v", tc.Input), func() {
					obj, err := store.Get(tc.Input.Data.ID)

					Convey("should return the expected value from store.Get", func() {
						So(err, ShouldBeNil)
						So(obj, ShouldEqual, tc.Expect.Data.Object)
					})
				})
			} else {
				Convey(fmt.Sprintf("given an invalid input: %v", tc.Input), func() {
					_, err := store.Get(tc.Input.Data.ID)
					Convey("should return error from store.Get", func() {
						So(err, ShouldNotBeNil)
					})
				})
			}
		}
	})

	Convey("List", t, func() {
		list, err := store.List()

		for _, tc := range testcases {
			if !tc.Expect.Error {
				Convey(fmt.Sprintf("given a valid input: %v", tc.Input), func() {
					obj := list[tc.Expect.Data.ID]

					Convey("should contain the expected value from store.List", func() {
						So(err, ShouldBeNil)
						So(obj, ShouldEqual, tc.Expect.Data.Object)
					})
				})
			} else {
				Convey(fmt.Sprintf("given an invalid input: %v", tc.Input), func() {
					obj := list[tc.Expect.Data.ID]

					Convey("should not contain this input value from store.List", func() {
						So(obj, ShouldBeNil)
					})
				})
			}
		}
	})

	Convey("Update", t, func() {
		for _, tc := range testcases {
			if !tc.Expect.Error {
				Convey(fmt.Sprintf("given a valid input: %v", tc.Input), func() {
					err := store.Update(tc.Input.Data.ID, tc.Input.Data.Object)

					Convey("should return no error on store.Update", func() {
						So(err, ShouldBeNil)
					})
				})
			} else {
				Convey(fmt.Sprintf("given an invalid input: %v", tc.Input), func() {
					err := store.Update(tc.Input.Data.ID, tc.Input.Data.Object)

					Convey("should return error on store.Update", func() {
						So(err, ShouldNotBeNil)
					})
				})
			}
		}
	})

	Convey("Delete", t, func() {
		for _, tc := range testcases {
			if !tc.Expect.Error {
				Convey(fmt.Sprintf("given a valid input: %v", tc.Input), func() {
					err := store.Delete(tc.Input.Data.ID)

					Convey("should return no error on store.Delete", func() {
						So(err, ShouldBeNil)
					})
				})
			} else {
				Convey(fmt.Sprintf("given an invalid input: %v", tc.Input), func() {
					_, err := store.Get(tc.Input.Data.ID)
					Convey("should return error on store.Delete", func() {
						So(err, ShouldNotBeNil)
					})
				})
			}
		}
	})
}
