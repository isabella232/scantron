package main_test

import (
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"

	yaml "gopkg.in/yaml.v2"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/gexec"
	"github.com/pivotal-cf/scantron/commands"
	"github.com/pivotal-cf/scantron/manifest"
	"github.com/pivotal-cf/scantron/scanner"
)

var _ = Describe("Main", func() {
	Describe("audit", func() {
		var (
			tmpdir       string
			manifestPath string
			databasePath string

			mani  manifest.Manifest
			hosts []scanner.ScanResult
		)

		BeforeEach(func() {
			var err error

			tmpdir, err = ioutil.TempDir("", "scantron-main-test")
			Expect(err).NotTo(HaveOccurred())

			manifestPath = filepath.Join(tmpdir, "manifest.yml")
			mani = manifest.Manifest{
				Specs: []manifest.Spec{
					{
						Prefix: "prefix",
					},
				},
			}

			databasePath = filepath.Join(tmpdir, "database.db")
			hosts = []scanner.ScanResult{
				{
					Job: "prefix-name-1",
				},
				{
					Job: "prefix-name-2",
				},
			}
		})

		JustBeforeEach(func() {
			manifestBytes, err := yaml.Marshal(mani)
			Expect(err).NotTo(HaveOccurred())

			err = ioutil.WriteFile(manifestPath, manifestBytes, 0600)
			Expect(err).NotTo(HaveOccurred())

			db, err := commands.CreateDatabase(databasePath)
			Expect(err).NotTo(HaveOccurred())
			defer db.Close()

			err = db.SaveReport(hosts)
			Expect(err).NotTo(HaveOccurred())
		})

		AfterEach(func() {
			os.RemoveAll(tmpdir)
		})

		Context("when the manifest does not exist", func() {
			JustBeforeEach(func() {
				err := os.RemoveAll(manifestPath)
				Expect(err).NotTo(HaveOccurred())
			})

			It("exits 1", func() {
				command := exec.Command(scantronPath, "audit", "--database", databasePath, "--manifest", manifestPath)
				process, err := gexec.Start(command, GinkgoWriter, GinkgoWriter)
				Expect(err).NotTo(HaveOccurred())

				Eventually(process).Should(gexec.Exit(1))
			})
		})

		Context("when the manifest is malformed", func() {
			JustBeforeEach(func() {
				err := ioutil.WriteFile(manifestPath, []byte("not-yaml"), 0600)
				Expect(err).NotTo(HaveOccurred())
			})

			It("exits 1", func() {
				command := exec.Command(scantronPath, "audit", "--database", databasePath, "--manifest", manifestPath)
				process, err := gexec.Start(command, GinkgoWriter, GinkgoWriter)
				Expect(err).NotTo(HaveOccurred())

				Eventually(process).Should(gexec.Exit(1))
			})
		})

		Context("when the audit fails", func() {
			BeforeEach(func() {
				hosts = []scanner.ScanResult{
					{
						Job: "not-the-right-prefix-name",
					},
				}
			})

			It("exits 3", func() {
				command := exec.Command(scantronPath, "audit", "--database", databasePath, "--manifest", manifestPath)
				process, err := gexec.Start(command, GinkgoWriter, GinkgoWriter)
				Expect(err).NotTo(HaveOccurred())

				Eventually(process).Should(gexec.Exit(3))
			})
		})

		It("shows ok for each host", func() {
			command := exec.Command(scantronPath, "audit", "--database", databasePath, "--manifest", manifestPath)
			process, err := gexec.Start(command, GinkgoWriter, GinkgoWriter)
			Expect(err).NotTo(HaveOccurred())

			Eventually(process).Should(gexec.Exit(0))

			output := process.Buffer().Contents()

			Expect(output).To(ContainSubstring("ok  prefix-name-1"))
			Expect(output).To(ContainSubstring("ok  prefix-name-2"))
		})
	})
})
