package acceptance_test

import (
	"testing"

	"github.com/onsi/gomega/gexec"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestAcceptance(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "acceptance")
}

var (
	pathToCheck string
	pathToIn    string
	pathToOut   string
)

var _ = BeforeSuite(func() {
	var err error

	pathToCheck, err = gexec.Build("github.com/christianang/staticfile-resource/cmd/check")
	Expect(err).NotTo(HaveOccurred())

	pathToIn, err = gexec.Build("github.com/christianang/staticfile-resource/cmd/in")
	Expect(err).NotTo(HaveOccurred())

	pathToOut, err = gexec.Build("github.com/christianang/staticfile-resource/cmd/out")
	Expect(err).NotTo(HaveOccurred())
})

var _ = AfterSuite(func() {
	gexec.CleanupBuildArtifacts()
})
