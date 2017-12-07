package stemcell

import (
	boshlog "github.com/cloudfoundry/bosh-utils/logger"
	"github.com/oracle/bosh-oracle-cpi/oci/client"
)

const stemCellLogTag = "OAOStemcell"

type Creator interface {
	CreateStemcell(imageSourceURL string, customImageName string) (stemcellId string, err error)
}
type CreatorFactory func(client.Connector, boshlog.Logger) Creator

type Destroyer interface {
	DeleteStemcell(stemcellId string) (err error)
}
type DestroyerFactory func(client.Connector, boshlog.Logger) Destroyer

type Finder interface {
	FindStemcell(imageOCID string) (stemcellId string, err error)
}
type FinderFactory func(client.Connector, boshlog.Logger) Finder

func NewCreator(c client.Connector, l boshlog.Logger) Creator {
	return stemcellOperations{connector: c, logger: l}
}

func NewDestroyer(c client.Connector, l boshlog.Logger) Destroyer {
	return stemcellOperations{connector: c, logger: l}
}

func NewFinder(c client.Connector, l boshlog.Logger) Finder {
	return stemcellOperations{connector: c, logger: l}
}
