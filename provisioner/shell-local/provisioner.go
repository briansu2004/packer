package shell

import (
	"context"

	"github.com/hashicorp/hcl/v2/hcldec"
	"github.com/hashicorp/packer/packer"
	sl "github.com/hashicorp/packer/packer-plugin-sdk/shell-local"
)

type Provisioner struct {
	config sl.Config
}

func (p *Provisioner) ConfigSpec() hcldec.ObjectSpec { return p.config.FlatMapstructure().HCL2Spec() }

func (p *Provisioner) Prepare(raws ...interface{}) error {
	err := sl.Decode(&p.config, raws...)
	if err != nil {
		return err
	}

	err = sl.Validate(&p.config)
	if err != nil {
		return err
	}

	return nil
}

func (p *Provisioner) Provision(ctx context.Context, ui packer.Ui, _ packer.Communicator, generatedData map[string]interface{}) error {
	_, retErr := sl.Run(ctx, ui, &p.config, generatedData)

	return retErr
}
