package backend

import (
	"bytes"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/blucard/blucard/internal/pkg/dynamo/dao"
	"github.com/blucard/blucard/internal/pkg/dynamo/dao/daofakes"
	"github.com/gin-gonic/gin"
	. "github.com/smartystreets/goconvey/convey"
)

func TestRecordBackend(t *testing.T) {
	Convey("Test Backend", t, func() {
		gin.SetMode(gin.TestMode)
		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)

		mockRecordDB := new(daofakes.FakeRecordDao)
		mockBackend := &backendImpl{
			recordDB: mockRecordDB,
		}

		Convey("Test GetRecord", func() {
			Convey("with success", func() {
				mockRecordDB.GetRecordReturns(&dao.Record{
					UUID: "123",
					Data: "test",
				}, nil)

				mockBackend.GetRecord(c)
				So(w.Code, ShouldEqual, http.StatusOK)
				So(w.Body, ShouldNotBeNil)
			})

			Convey("with error", func() {
				mockRecordDB.GetRecordReturns(nil, errors.New("test error"))

				mockBackend.GetRecord(c)
				So(w.Code, ShouldEqual, http.StatusInternalServerError)
			})
		})

		Convey("Test SetRecord", func() {
			c.Request = httptest.NewRequest("POST", "/", bytes.NewBufferString("{\"data\": \"test\"}"))
			c.Params = []gin.Param{
				{
					Key:   "uuid",
					Value: "123",
				},
			}

			Convey("with success", func() {
				mockRecordDB.SetRecordReturns(nil)

				mockBackend.SetRecord(c)
				So(w.Code, ShouldEqual, http.StatusAccepted)
				So(w.Body, ShouldNotBeNil)
			})

			Convey("with error", func() {
				mockRecordDB.SetRecordReturns(errors.New("test error"))

				mockBackend.SetRecord(c)
				So(w.Code, ShouldEqual, http.StatusInternalServerError)
			})
		})
	})
}
