package controller_test

import (
	"context"
	"net/http"
	"testing"
	"time"

	"github.com/amartha-shorty/pkg/model"
	"github.com/amartha-shorty/pkg/shorty/controller"
	"github.com/amartha-shorty/pkg/shorty/repository"
	. "github.com/smartystreets/goconvey/convey"
)

func TestShortyController(t *testing.T) {
	Convey("Prepare statement", t, func() {
		shortyRepository := repository.NewShortyRepository()
		shortyController := controller.NewShortyController(time.Second*2, shortyRepository)

		Convey("Shorten test", func() {
			Convey("Given Incorrect Parameters on url", func() {
				shortCode, httpStatusCode, err := shortyController.Shorten(context.TODO(), "google.com", "google")
				So(shortCode, ShouldEqual, "")
				So(httpStatusCode, ShouldEqual, http.StatusBadRequest)
				So(err, ShouldNotBeNil)
			})

			Convey("Given Incorrect Parameters on shortcode already used", func() {
				shortCode, httpStatusCode, err := shortyController.Shorten(context.TODO(), "http://google.com", "google")
				shortCode, httpStatusCode, err = shortyController.Shorten(context.TODO(), "http://google.com", "google")
				So(shortCode, ShouldEqual, "")
				So(httpStatusCode, ShouldEqual, http.StatusConflict)
				So(err, ShouldNotBeNil)
			})

			Convey("Given Incorrect Parameters on shortcode (less than 6 character)", func() {
				shortCode, httpStatusCode, err := shortyController.Shorten(context.TODO(), "http://google.com", "goog")
				So(shortCode, ShouldEqual, "")
				So(httpStatusCode, ShouldEqual, http.StatusUnprocessableEntity)
				So(err, ShouldNotBeNil)
			})

			Convey("Given Incorrect Parameters on shortcode (more than 6 character)", func() {
				shortCode, httpStatusCode, err := shortyController.Shorten(context.TODO(), "http://google.com", "googles")
				So(shortCode, ShouldEqual, "")
				So(httpStatusCode, ShouldEqual, http.StatusUnprocessableEntity)
				So(err, ShouldNotBeNil)
			})

			Convey("Given Correct Parameters without short code", func() {
				shortCode, httpStatusCode, err := shortyController.Shorten(context.TODO(), "http://google.com", "")
				So(shortCode, ShouldNotBeNil)
				So(httpStatusCode, ShouldEqual, http.StatusCreated)
				So(err, ShouldBeNil)
			})

			Convey("Given Correct Parameters with short code", func() {
				shortCode, httpStatusCode, err := shortyController.Shorten(context.TODO(), "http://google.com", "google")
				So(shortCode, ShouldEqual, "google")
				So(httpStatusCode, ShouldEqual, http.StatusCreated)
				So(err, ShouldBeNil)
			})
		})

		Convey("ShortCode test", func() {
			Convey("Given Incorrect Parameters with unavailable data", func() {
				url, httpStatusCode, err := shortyController.ShortCode(context.TODO(), "abcdef")
				So(url, ShouldEqual, "")
				So(httpStatusCode, ShouldEqual, http.StatusNotFound)
				So(err, ShouldNotBeNil)
			})

			Convey("Given Correct Parameters with available data", func() {
				shortCode, _, err := shortyController.Shorten(context.TODO(), "http://ngantuk.com", "abcdef")
				url, httpStatusCode, err := shortyController.ShortCode(context.TODO(), shortCode)
				So(url, ShouldEqual, "http://ngantuk.com")
				So(httpStatusCode, ShouldEqual, http.StatusFound)
				So(err, ShouldBeNil)
			})
		})

		Convey("ShortCodeStats test", func() {
			Convey("Given Incorrect Parameters with unavailable data", func() {
				stats, httpStatusCode, err := shortyController.ShortCodeStats(context.TODO(), "fedcba")
				So(stats, ShouldResemble, model.ShortySpec{})
				So(httpStatusCode, ShouldEqual, http.StatusNotFound)
				So(err, ShouldNotBeNil)
			})

			Convey("Given Correct Parameters with available data", func() {
				shortCode, _, err := shortyController.Shorten(context.TODO(), "http://ngantuk.com", "fedcba")
				_, _, err = shortyController.ShortCode(context.TODO(), shortCode)
				_, _, err = shortyController.ShortCode(context.TODO(), shortCode)
				stats, httpStatusCode, err := shortyController.ShortCodeStats(context.TODO(), shortCode)
				So(stats.RedirectCount, ShouldEqual, 2)
				So(httpStatusCode, ShouldEqual, http.StatusOK)
				So(err, ShouldBeNil)
			})
		})
	})
}
