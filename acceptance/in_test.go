package acceptance_test

import (
	"io"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("In", func() {
	var (
		tempDir string
	)

	BeforeEach(func() {
		var err error
		tempDir, err = ioutil.TempDir("", "")
		Expect(err).NotTo(HaveOccurred())
	})

	Context("when given arbitrary data, a filename, and destination directory", func() {
		var (
			in *exec.Cmd
		)

		BeforeEach(func() {
			in = exec.Command(pathToIn, tempDir)
			in.Stderr = os.Stderr

			stdin, err := in.StdinPipe()
			Expect(err).NotTo(HaveOccurred())

			_, err = io.WriteString(stdin, `{
				"source": {
					"filename": "example.json",
					"data": "some-data"
				},
				"version": {}
			}`)
			Expect(err).NotTo(HaveOccurred())
		})

		It("writes the data to the destination directory with the given filename", func() {
			outputJSON, err := in.Output()
			Expect(err).NotTo(HaveOccurred())
			Expect(outputJSON).To(MatchJSON(`{
				"version": {}
			}`))

			data, err := ioutil.ReadFile(filepath.Join(tempDir, "example.json"))
			Expect(err).NotTo(HaveOccurred())
			Expect(string(data)).To(Equal("some-data"))
		})
	})
})
