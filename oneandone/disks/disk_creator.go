package disks

import (
	"fmt"
	boshlog "github.com/cloudfoundry/bosh-utils/logger"
	"github.com/oneandone/oneandone-cloudserver-sdk-go"
)

type diskCreator struct {
	connector client.Connector
	logger    boshlog.Logger
	vmId      string
}

func NewCreator(c client.Connector, l boshlog.Logger, vmId string) Creator {
	return &diskCreator{connector: c, logger: l, vmId: vmId}
}

type CreatorFactory func(client.Connector, boshlog.Logger, string) Creator

func (dc *diskCreator) CreateVolume(name string, sizeinMB int, vmId string) (*resource.Volume, error) {

	req := oneandone.ServerHdds{
		Hdds: []oneandone.Hdd{
			oneandone.Hdd{
				Size:   sizeinMB,
				IsMain: false,
			},
		},
	}

	res, err := dc.connector.Client().AddServerHdds(vmId, &req)

	if err != nil {
		return nil, fmt.Errorf("Error creating volume. Reason: %s", err)
	}

	var volume *resource.Volume
	dc.connector.Client().WaitForState(res,"POWERED_ON",10,90)
	var volume *resource.Volume
	if err = dc.waitUntilProvisioned(res.Payload, func(v *models.Volume) {
		volume = resource.NewVolume(*v.DisplayName, *v.ID)
	}); err != nil {
		return nil, err
	}
	dc.logger.Debug(diskOperationsLogTag, "Created volume %v", volume)
	return volume, nil
}

func (dc *diskCreator) waitUntilProvisioned(v *models.Volume, acceptor func(*models.Volume)) error {
	waiter := &volumeAvailableWaiter{connector: dc.connector, logger: dc.logger,
		availableHandler: acceptor}
	return waiter.WaitFor(v)
}
