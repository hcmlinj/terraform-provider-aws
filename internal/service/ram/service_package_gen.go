// Code generated by internal/generate/servicepackages/main.go; DO NOT EDIT.

package ram

import (
	"context"

	"github.com/hashicorp/terraform-provider-aws/internal/types"
	"github.com/hashicorp/terraform-provider-aws/names"
)

type servicePackage struct{}

func (p *servicePackage) FrameworkDataSources(ctx context.Context) []*types.ServicePackageFrameworkDataSource {
	return []*types.ServicePackageFrameworkDataSource{}
}

func (p *servicePackage) FrameworkResources(ctx context.Context) []*types.ServicePackageFrameworkResource {
	return []*types.ServicePackageFrameworkResource{}
}

func (p *servicePackage) SDKDataSources(ctx context.Context) []*types.ServicePackageSDKDataSource {
	return []*types.ServicePackageSDKDataSource{
		{
			Factory:  DataSourceResourceShare,
			TypeName: "aws_ram_resource_share",
		},
	}
}

func (p *servicePackage) SDKResources(ctx context.Context) []*types.ServicePackageSDKResource {
	return []*types.ServicePackageSDKResource{
		{
			Factory:  ResourcePrincipalAssociation,
			TypeName: "aws_ram_principal_association",
		},
		{
			Factory:  ResourceResourceAssociation,
			TypeName: "aws_ram_resource_association",
		},
		{
			Factory:  ResourceResourceShare,
			TypeName: "aws_ram_resource_share",
		},
		{
			Factory:  ResourceResourceShareAccepter,
			TypeName: "aws_ram_resource_share_accepter",
		},
	}
}

func (p *servicePackage) ServicePackageName() string {
	return names.RAM
}

var ServicePackage = &servicePackage{}
