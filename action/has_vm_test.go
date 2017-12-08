package action

import (
	"errors"
	boshlog "github.com/cloudfoundry/bosh-utils/logger"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	clientfakes "github.com/bosh-oneandone-cpi/oneandone/client/fakes"
	vmfakes "github.com/bosh-oneandone-cpi/oneandone/vm/fakes"

	"github.com/bosh-oneandone-cpi/oneandone/client"
	"github.com/bosh-oneandone-cpi/oneandone/resource"
	"github.com/bosh-oneandone-cpi/oneandone/vm"
)

var _ = Describe("HasVM", func() {
	var (
		connector *clientfakes.FakeConnector
		logger    boshlog.Logger
		finder    *vmfakes.FakeVMFinder

		found bool
		err   error

		hasVM HasVM
	)

	BeforeEach(func() {

		connector = &clientfakes.FakeConnector{}
		logger = boshlog.NewLogger(boshlog.LevelNone)

		finder = &vmfakes.FakeVMFinder{}
		installVMFinderFactory(func(c client.Connector, l boshlog.Logger) vm.Finder {
			return finder
		})

		hasVM = NewHasVM(connector, logger)
	})
	AfterEach(func() { resetAllFactories() })

	Describe("Run", func() {
		It("returns true if an instance exists", func() {
			finder.FindInstanceResult = resource.NewInstance("fake-ocid")

			found, err = hasVM.Run("fake-ocid")
			Expect(finder.FindInstanceCalled).To(BeTrue())
			Expect(err).NotTo(HaveOccurred())
			Expect(found).To(BeTrue())
		})

		It("returns false if an instance does not exist", func() {
			found, err = hasVM.Run("fake-ocid")
			Expect(finder.FindInstanceCalled).To(BeTrue())
			Expect(err).NotTo(HaveOccurred())
			Expect(found).To(BeFalse())
		})

		It("Returns error if finder call returns an error", func() {
			finder.FindInstanceError = errors.New("fake-finder-error")

			found, err = hasVM.Run("fake-ocid")
			Expect(finder.FindInstanceCalled).To(BeTrue())
			Expect(err).To(HaveOccurred())
			Expect(found).To(BeFalse())
			Expect(err.Error()).To(ContainSubstring("fake-finder-error"))
		})
	})
})
