package acceptance_test

import (
	"io"
	"os"
	"os/exec"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Check", func() {
	It("returns an empty version", func() {
		check := exec.Command(pathToCheck)
		check.Stderr = os.Stderr

		stdin, err := check.StdinPipe()
		Expect(err).NotTo(HaveOccurred())

		_, err = io.WriteString(stdin, `{
			"source": {
				"filename": "some-filename",
				"data": "some-data"
			},
			"version": {}
		}`)
		Expect(err).NotTo(HaveOccurred())

		output, err := check.Output()
		Expect(err).NotTo(HaveOccurred())
		Expect(output).To(MatchJSON("[{}]"))
	})
})
