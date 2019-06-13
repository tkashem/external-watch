package external

import (
	"github.com/tkashem/external-watch/lib/appregistry"
)

type Lister struct {
	
}

func (l *Lister) List(source *OperatorSource) ( items []*appregistry.RegistryMetadata, err error ) {	
	options := appregistry.Options{
		Source: source.Endpoint,
	}

	factory := appregistry.NewClientFactory()

	client, err := factory.New(options)
	if err != nil {
		return
	}

	list, err := client.ListPackages(source.Namespace)
	if err != nil {
		return
	}	

	items = list
	return
}
