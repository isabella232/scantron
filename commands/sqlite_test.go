package commands_test

import (
	"database/sql"
	"io/ioutil"
	"os"
	"path/filepath"
	"time"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/pivotal-cf/scantron"
	"github.com/pivotal-cf/scantron/commands"
	"github.com/pivotal-cf/scantron/scanner"
)

var _ = Describe("Sqlite", func() {
	var tmpdir string

	BeforeEach(func() {
		var err error
		tmpdir, err = ioutil.TempDir("", "scantron_db")
		Expect(err).NotTo(HaveOccurred())
	})

	AfterEach(func() {
		os.RemoveAll(tmpdir)
	})

	Describe("NewDatabase", func() {
		It("creates a new database file", func() {
			dbPath := filepath.Join(tmpdir, "database.db")

			db, err := commands.NewDatabase(dbPath)
			Expect(err).NotTo(HaveOccurred())
			defer db.Close()

			Expect(dbPath).To(BeAnExistingFile())
		})

		It("creates the required tables", func() {
			dbPath := filepath.Join(tmpdir, "database.db")

			db, err := commands.NewDatabase(dbPath)
			Expect(err).NotTo(HaveOccurred())
			defer db.Close()

			sqliteDB, err := sql.Open("sqlite3", dbPath)
			Expect(err).NotTo(HaveOccurred())
			defer sqliteDB.Close()

			tables := []string{}

			rows, err := sqliteDB.Query(`
				SELECT name
				FROM sqlite_master
				WHERE type = 'table'
					AND name NOT LIKE 'sqlite_%'`,
			)
			Expect(err).NotTo(HaveOccurred())
			defer rows.Close()

			var table string
			for rows.Next() {
				err = rows.Scan(&table)
				Expect(err).NotTo(HaveOccurred())

				tables = append(tables, table)
			}

			Expect(tables).To(ConsistOf(
				"hosts",
				"processes",
				"ports",
				"files",
				"tls_informations",
				"env_vars",
			))
		})
	})

	Describe("SaveReport", func() {
		var (
			db             *commands.Database
			hosts          []scanner.ScannedHost
			host           scanner.ScannedHost
			dbPath         string
			sqliteDB       *sql.DB
			certExpiration time.Time
		)

		JustBeforeEach(func() {
			dbPath = filepath.Join(tmpdir, "database.db")

			var err error
			db, err = commands.NewDatabase(dbPath)
			Expect(err).NotTo(HaveOccurred())

			err = db.SaveReport(hosts)
			Expect(err).NotTo(HaveOccurred())

			sqliteDB, err = sql.Open("sqlite3", dbPath)
			Expect(err).NotTo(HaveOccurred())
		})

		AfterEach(func() {
			db.Close()
			sqliteDB.Close()
		})

		Context("with a single host", func() {
			BeforeEach(func() {
				var err error
				certExpiration, err = time.Parse(time.RFC3339, "2012-11-01T22:08:41+00:00")
				Expect(err).NotTo(HaveOccurred())

				host = scanner.ScannedHost{
					IP:  "10.0.0.1",
					Job: "custom_name/0",
					Services: []scanner.ScannedService{{
						Name: "server-name",
						PID:  213,
						User: "root",
						Ports: []scantron.Port{
							{
								Protocol: "TCP",
								Address:  "123.0.0.1",
								Number:   123,
							},
						},
						TLSInformation: scanner.TLSInformation{
							Presence: true,
							Certificate: &scanner.Certificate{
								Expiration: certExpiration,
								Bits:       234,
								Subject: scanner.CertificateSubject{
									Country:      "some-country",
									Province:     "some-province",
									Locality:     "some-locality",
									Organization: "some-organization",
									CommonName:   "some-common-name",
								},
							},
						},
						Cmd: scanner.Cmd{
							Cmdline: []string{"this", "is", "a", "cmd"},
							Env:     []string{"PATH=this", "OTHER=that"},
						},
					}},
					Files: []scantron.File{
						{Path: "some-file-path"},
					},
				}

				hosts = []scanner.ScannedHost{host}
			})

			It("records a process", func() {
				rows, err := sqliteDB.Query(`
				SELECT hosts.name,
							 hosts.ip,
							 processes.pid,
							 processes.user,
							 processes.cmdline,
							 ports.protocol,
							 ports.address,
							 ports.number,
							 tls_informations.cert_expiration,
							 tls_informations.cert_bits,
							 tls_informations.cert_country,
							 tls_informations.cert_province,
							 tls_informations.cert_locality,
							 tls_informations.cert_organization,
							 tls_informations.cert_common_name,
							 env_vars.var,
							 files.path
				FROM   hosts,
							 processes,
							 ports,
							 tls_informations,
							 env_vars,
							 files
				WHERE  hosts.id = processes.host_id
							 AND ports.process_id = processes.id
							 AND ports.id = tls_informations.port_id
							 AND files.host_id = hosts.id
						   AND env_vars.process_id=processes.id`)
				Expect(err).NotTo(HaveOccurred())
				defer rows.Close()

				hasRows := rows.Next()
				Expect(hasRows).To(BeTrue())

				var (
					name, ip, user, cmdline, env, portProtocol, portAddress string
					pid, portNumber                                         int

					tlsCertCountry,
					tlsCertProvince,
					tlsCertLocality,
					tlsCertOrganization,
					tlsCertCommonName string
					tlsCertBits int
					tlsCertExp  time.Time
					file_path   string
				)

				err = rows.Scan(&name, &ip, &pid, &user, &cmdline, &portProtocol,
					&portAddress, &portNumber, &tlsCertExp, &tlsCertBits,
					&tlsCertCountry, &tlsCertProvince, &tlsCertLocality,
					&tlsCertOrganization, &tlsCertCommonName, &env, &file_path)
				Expect(err).NotTo(HaveOccurred())

				Expect(name).To(Equal("custom_name/0"))
				Expect(ip).To(Equal("10.0.0.1"))
				Expect(pid).To(Equal(213))
				Expect(user).To(Equal("root"))
				Expect(portAddress).To(Equal("123.0.0.1"))
				Expect(portNumber).To(Equal(123))
				Expect(tlsCertExp.Equal(certExpiration)).To(BeTrue())
				Expect(tlsCertBits).To(Equal(234))
				Expect(tlsCertCountry).To(Equal("some-country"))
				Expect(tlsCertProvince).To(Equal("some-province"))
				Expect(tlsCertLocality).To(Equal("some-locality"))
				Expect(tlsCertOrganization).To(Equal("some-organization"))
				Expect(tlsCertCommonName).To(Equal("some-common-name"))
				Expect(cmdline).To(Equal("this is a cmd"))
				Expect(env).To(Equal("PATH=this OTHER=that"))
				Expect(file_path).To(Equal("some-file-path"))
			})

			Context("when the service does not have a certificate", func() {
				BeforeEach(func() {
					service := host.Services[0]

					service.TLSInformation.Certificate = nil
					service.TLSInformation.Presence = false

					host.Services[0] = service

				})

				It("records a process", func() {
					rows, err := sqliteDB.Query(`
						SELECT hosts.NAME,
									 hosts.ip,
									 processes.pid,
									 processes.USER,
									 ports.protocol,
									 ports.address,
									 ports.number
						FROM   hosts,
									 processes,
									 ports
						WHERE  hosts.id = processes.host_id
									 AND ports.process_id = processes.id`)

					Expect(err).NotTo(HaveOccurred())
					defer rows.Close()

					hasRows := rows.Next()
					Expect(hasRows).To(BeTrue())

					var (
						name, ip, user, portProtocol, portAddress string
						pid, portNumber                           int
					)

					err = rows.Scan(&name, &ip, &pid, &user, &portProtocol, &portAddress, &portNumber)
					Expect(err).NotTo(HaveOccurred())

					Expect(name).To(Equal("custom_name/0"))
					Expect(ip).To(Equal("10.0.0.1"))
					Expect(pid).To(Equal(213))
					Expect(user).To(Equal("root"))
					Expect(portProtocol).To(Equal("TCP"))
					Expect(portAddress).To(Equal("123.0.0.1"))
					Expect(portNumber).To(Equal(123))
				})

				It("does not store any tls information", func() {
					rows, err := sqliteDB.Query(`SELECT count(1) FROM tls_informations`)
					Expect(err).NotTo(HaveOccurred())
					defer rows.Close()

					hasRows := rows.Next()
					Expect(hasRows).To(BeTrue())

					var count int
					err = rows.Scan(&count)
					Expect(err).NotTo(HaveOccurred())

					Expect(count).To(BeZero())
				})
			})
		})

		Context("with a multiple services that have the same host and job", func() {
			BeforeEach(func() {
				hosts = []scanner.ScannedHost{
					{
						IP:  "10.0.0.1",
						Job: "custom_name/0",
					},
					{
						IP:  "10.0.0.1",
						Job: "custom_name/0",
					},
				}
			})

			It("records only a single host", func() {
				rows, err := sqliteDB.Query(`SELECT COUNT(*) FROM hosts`)
				Expect(err).NotTo(HaveOccurred())
				defer rows.Close()

				hasRows := rows.Next()
				Expect(hasRows).To(BeTrue())

				var count int
				err = rows.Scan(&count)
				Expect(err).NotTo(HaveOccurred())

				Expect(count).To(Equal(1))
			})
		})

		Context("with a multiple services on different hosts", func() {
			BeforeEach(func() {
				hosts = []scanner.ScannedHost{
					{
						IP:  "10.0.0.1",
						Job: "custom_name/0",
					},
					{
						IP:  "10.0.0.2",
						Job: "custom_name/0",
					},
				}
			})

			It("records both hosts", func() {
				rows, err := sqliteDB.Query(`SELECT COUNT(*) FROM hosts`)
				Expect(err).NotTo(HaveOccurred())
				defer rows.Close()

				hasRows := rows.Next()
				Expect(hasRows).To(BeTrue())

				var count int
				err = rows.Scan(&count)
				Expect(err).NotTo(HaveOccurred())

				Expect(count).To(Equal(2))
			})
		})
	})
})