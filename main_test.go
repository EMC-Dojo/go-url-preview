package main_test

import (
	"fmt"
	"net/http"
	"net/http/httptest"

	"github.com/EMC-Dojo/go-url-preview/server"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/ghttp"
)

var _ = Describe("Main", func() {
	Describe("getTitle", func() {
		var s *ghttp.Server

		BeforeSuite(func() {
			s = ghttp.NewServer()
		})

		AfterEach(func() {
			//shut down the server between tests
			s.Close()
		})

		Context("When a URL is passed to getTitle", func() {
			It("Will respond with a title from the URL", func() {
				rr := httptest.NewRecorder()
				req, err := http.NewRequest("GET", fmt.Sprintf("/getTitle?url=%s/someurl", s.URL()), nil)
				Expect(err).ToNot(HaveOccurred())

				s.AppendHandlers(ghttp.CombineHandlers(
					ghttp.VerifyRequest("GET", "/someurl"),
					ghttp.RespondWith(http.StatusOK, `
					<HTML><HEAD><meta>
						<TITLE>MOCK-TITLE</TITLE></HEAD><BODY>
						<H1>lel</H1>
					</BODY></HTML>`),
				))

				handler := http.HandlerFunc(server.GetTitle)
				handler.ServeHTTP(rr, req)
				Expect(rr.Body.String()).To(Equal(`{"title": "MOCK-TITLE"}`))
			})
		})
	})

})
