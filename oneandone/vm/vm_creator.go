package vm

import (
	boshlog "github.com/cloudfoundry/bosh-utils/logger"
	"github.com/bosh-oneandone-cpi/oneandone/resource"
	"github.com/bosh-oneandone-cpi/oneandone/client"
	"github.com/oneandone/oneandone-cloudserver-sdk-go"
)

const logTag = "VMOperations"

type InstanceConfiguration struct {
	ImageId string
	Name    string
	ServerIp string
	Location string
	SecondaryIps []string
}

type Creator interface {
	CreateInstance(icfg InstanceConfiguration, md InstanceMetadata) (*resource.Instance, error)
}

type CreatorFactory func(client.Connector, boshlog.Logger, string) Creator

type creator struct {
	connector client.Connector
	logger    boshlog.Logger
	location  string
}

func NewCreator(c client.Connector, l boshlog.Logger) Creator {
	return &creator{connector: c, logger: l,
		location: "us/las",
	}
}

func (cv *creator) CreateInstance(icfg InstanceConfiguration,
	md InstanceMetadata) (*resource.Instance, error) {

	//if err := icfg.Network.validate(); err != nil {
	//	return nil, err
	//}
	return cv.launchInstance(icfg, md)
}
func (cv *creator) launchInstance(icfg InstanceConfiguration, md InstanceMetadata) (*resource.Instance, error) {

	req := oneandone.ServerRequest{}
	_, res, err := cv.connector.Client().CreateServer(&req)
	if err != nil {
		return nil, err
	}
	instance := resource.NewInstance(res.Id)

	//if icfg.Network.hasSecondaries() {
	//	err = cv.attachSecondaryVnics(instance, icfg.Network.secondaries())
	//	if err != nil {
	//		return nil, newLaunchInstanceError(err)
	//	}
	//}
	return instance, nil
}
