// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/v2/pkg/controller"

	providerconfig "github.com/STASiAN/crossplane-provider-vaultwarden/internal/controller/namespaced/providerconfig"
	accountregister "github.com/STASiAN/crossplane-provider-vaultwarden/internal/controller/namespaced/vaultwarden/accountregister"
	folder "github.com/STASiAN/crossplane-provider-vaultwarden/internal/controller/namespaced/vaultwarden/folder"
	itemattachment "github.com/STASiAN/crossplane-provider-vaultwarden/internal/controller/namespaced/vaultwarden/itemattachment"
	itemlogin "github.com/STASiAN/crossplane-provider-vaultwarden/internal/controller/namespaced/vaultwarden/itemlogin"
	itemsecurenote "github.com/STASiAN/crossplane-provider-vaultwarden/internal/controller/namespaced/vaultwarden/itemsecurenote"
	organization "github.com/STASiAN/crossplane-provider-vaultwarden/internal/controller/namespaced/vaultwarden/organization"
	organizationapikeyrotation "github.com/STASiAN/crossplane-provider-vaultwarden/internal/controller/namespaced/vaultwarden/organizationapikeyrotation"
	organizationcollection "github.com/STASiAN/crossplane-provider-vaultwarden/internal/controller/namespaced/vaultwarden/organizationcollection"
	organizationcollectiongroup "github.com/STASiAN/crossplane-provider-vaultwarden/internal/controller/namespaced/vaultwarden/organizationcollectiongroup"
	organizationcollectionuser "github.com/STASiAN/crossplane-provider-vaultwarden/internal/controller/namespaced/vaultwarden/organizationcollectionuser"
	organizationgroup "github.com/STASiAN/crossplane-provider-vaultwarden/internal/controller/namespaced/vaultwarden/organizationgroup"
	organizationgroupmember "github.com/STASiAN/crossplane-provider-vaultwarden/internal/controller/namespaced/vaultwarden/organizationgroupmember"
	organizationpolicy "github.com/STASiAN/crossplane-provider-vaultwarden/internal/controller/namespaced/vaultwarden/organizationpolicy"
	organizationuser "github.com/STASiAN/crossplane-provider-vaultwarden/internal/controller/namespaced/vaultwarden/organizationuser"
	user "github.com/STASiAN/crossplane-provider-vaultwarden/internal/controller/namespaced/vaultwarden/user"
)

// Setup creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		providerconfig.Setup,
		accountregister.Setup,
		folder.Setup,
		itemattachment.Setup,
		itemlogin.Setup,
		itemsecurenote.Setup,
		organization.Setup,
		organizationapikeyrotation.Setup,
		organizationcollection.Setup,
		organizationcollectiongroup.Setup,
		organizationcollectionuser.Setup,
		organizationgroup.Setup,
		organizationgroupmember.Setup,
		organizationpolicy.Setup,
		organizationuser.Setup,
		user.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

// SetupGated creates all controllers with the supplied logger and adds them to
// the supplied manager gated.
func SetupGated(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		providerconfig.SetupGated,
		accountregister.SetupGated,
		folder.SetupGated,
		itemattachment.SetupGated,
		itemlogin.SetupGated,
		itemsecurenote.SetupGated,
		organization.SetupGated,
		organizationapikeyrotation.SetupGated,
		organizationcollection.SetupGated,
		organizationcollectiongroup.SetupGated,
		organizationcollectionuser.SetupGated,
		organizationgroup.SetupGated,
		organizationgroupmember.SetupGated,
		organizationpolicy.SetupGated,
		organizationuser.SetupGated,
		user.SetupGated,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

// SetupWebhookWithManager registers conversion webhooks for all resource kinds in the group.
func SetupWebhookWithManager(mgr ctrl.Manager) error {
	for _, setup := range []func(ctrl.Manager) error{
		providerconfig.SetupWebhookWithManager,
		accountregister.SetupWebhookWithManager,
		folder.SetupWebhookWithManager,
		itemattachment.SetupWebhookWithManager,
		itemlogin.SetupWebhookWithManager,
		itemsecurenote.SetupWebhookWithManager,
		organization.SetupWebhookWithManager,
		organizationapikeyrotation.SetupWebhookWithManager,
		organizationcollection.SetupWebhookWithManager,
		organizationcollectiongroup.SetupWebhookWithManager,
		organizationcollectionuser.SetupWebhookWithManager,
		organizationgroup.SetupWebhookWithManager,
		organizationgroupmember.SetupWebhookWithManager,
		organizationpolicy.SetupWebhookWithManager,
		organizationuser.SetupWebhookWithManager,
		user.SetupWebhookWithManager,
	} {
		if err := setup(mgr); err != nil {
			return err
		}
	}
	return nil
}
