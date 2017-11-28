package api_test

import (
	"io/ioutil"
	"path/filepath"

	"github.com/christianang/staticfile-resource/api"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("In", func() {
	Describe("WriteDataToDestination", func() {
		var (
			in      api.In
			tempDir string
		)

		BeforeEach(func() {
			in = api.NewIn()

			var err error
			tempDir, err = ioutil.TempDir("", "")
			Expect(err).NotTo(HaveOccurred())
		})

		It("writes data to file", func() {
			err := in.WriteDataToDestination(
				filepath.Join(tempDir, "example.json"),
				[]byte("some-data"),
			)
			Expect(err).NotTo(HaveOccurred())

			data, err := ioutil.ReadFile(filepath.Join(tempDir, "example.json"))
			Expect(err).NotTo(HaveOccurred())
			Expect(string(data)).To(Equal("some-data"))
		})

		Context("when a failure occurs", func() {
			Context("when it fails to write the file", func() {
				It("returns an error", func() {
					err := in.WriteDataToDestination(
						"/some/fake/dir",
						[]byte("some-data"),
					)
					Expect(err).To(MatchError(ContainSubstring("no such file or directory")))
				})
			})
		})
	})
})
