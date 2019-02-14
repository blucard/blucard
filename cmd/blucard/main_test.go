package main

import (
	"testing"

	"github.com/blucard/blucard/internal/backend"
	"github.com/blucard/blucard/internal/pkg/dynamo/dao/daofakes"
	"github.com/blucard/blucard/internal/server"

	. "github.com/smartystreets/goconvey/convey"
)

func TestCmdMain(t *testing.T) {
	Convey("Testing main", t, func() {
		Convey("when configuring backend", func() {
			recordDao := new(daofakes.FakeRecordDao)
			testBackend := backend.NewBackend(&backend.NewBackendInput{
				RecordDB: recordDao,
			})
			So(testBackend, ShouldNotBeNil)
		})

		Convey("when configuring server", func() {
			recordDao := new(daofakes.FakeRecordDao)
			testBackend := backend.NewBackend(&backend.NewBackendInput{
				RecordDB: recordDao,
			})
			server, err := server.NewServer(testBackend)
			So(err, ShouldBeNil)
			So(server, ShouldNotBeNil)
		})
	})
}
