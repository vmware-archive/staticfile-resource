package acceptance_test

import (
	"io"
	"os"
	"os/exec"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Out", func() {
	It("returns an empty version", func() {
		out := exec.Command(pathToOut)
		out.Stderr = os.Stderr

		stdin, err := out.StdinPipe()
		Expect(err).NotTo(HaveOccurred())

		_, err = io.WriteString(stdin, `{
			"source": {
				"filename": "some-filename",
				"data": "some-data"
			},
			"version": {}
		}`)
		Expect(err).NotTo(HaveOccurred())

		output, err := out.Output()
		Expect(err).NotTo(HaveOccurred())
		Expect(output).To(MatchJSON(`{
			"version": {
				"version": "latest"
			}
		}`))
	})
})
