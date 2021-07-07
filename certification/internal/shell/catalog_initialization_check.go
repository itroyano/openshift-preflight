package shell

import (
	"github.com/komish/preflight/certification"
	"github.com/sirupsen/logrus"
)

type CatalogInitializationCheck struct{}

func (p *CatalogInitializationCheck) Validate(image string, logger *logrus.Logger) (bool, error) {
	// TODO

	return false, nil
}

func (p *CatalogInitializationCheck) Name() string {
	return "OperatorCatalogInitializationCheck"
}
func (p *CatalogInitializationCheck) Metadata() certification.Metadata {
	return certification.Metadata{
		Description:      "Validating Catalog initialization",
		Level:            "best",
		KnowledgeBaseURL: "https://connect.redhat.com/zones/containers/container-certification-policy-guide",
		CheckURL:         "https://connect.redhat.com/zones/containers/container-certification-policy-guide",
	}
}

func (p *CatalogInitializationCheck) Help() certification.HelpText {
	return certification.HelpText{
		Message:    "Catalog initialization test failed, this test checks if the operator bundle can be added to a Catalog",
		Suggestion: "Verify the integrity of the operator's catalog manifests dir",
	}
}
