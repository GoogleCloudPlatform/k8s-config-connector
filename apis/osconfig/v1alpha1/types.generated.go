// Copyright 2025 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package v1alpha1


// +kcc:proto=google.cloud.osconfig.v1.Inventory
type Inventory struct {

	// Base level operating system information for the VM.
	// +kcc:proto:field=google.cloud.osconfig.v1.Inventory.os_info
	OsInfo *Inventory_OsInfo `json:"osInfo,omitempty"`

	// TODO: unsupported map type with key string and value message

}

// +kcc:proto=google.cloud.osconfig.v1.Inventory.Item
type Inventory_Item struct {
	// Identifier for this item, unique across items for this VM.
	// +kcc:proto:field=google.cloud.osconfig.v1.Inventory.Item.id
	ID *string `json:"id,omitempty"`

	// The origin of this inventory item.
	// +kcc:proto:field=google.cloud.osconfig.v1.Inventory.Item.origin_type
	OriginType *string `json:"originType,omitempty"`

	// When this inventory item was first detected.
	// +kcc:proto:field=google.cloud.osconfig.v1.Inventory.Item.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// When this inventory item was last modified.
	// +kcc:proto:field=google.cloud.osconfig.v1.Inventory.Item.update_time
	UpdateTime *string `json:"updateTime,omitempty"`

	// The specific type of inventory, correlating to its specific details.
	// +kcc:proto:field=google.cloud.osconfig.v1.Inventory.Item.type
	Type *string `json:"type,omitempty"`

	// Software package present on the VM instance.
	// +kcc:proto:field=google.cloud.osconfig.v1.Inventory.Item.installed_package
	InstalledPackage *Inventory_SoftwarePackage `json:"installedPackage,omitempty"`

	// Software package available to be installed on the VM instance.
	// +kcc:proto:field=google.cloud.osconfig.v1.Inventory.Item.available_package
	AvailablePackage *Inventory_SoftwarePackage `json:"availablePackage,omitempty"`
}

// +kcc:proto=google.cloud.osconfig.v1.Inventory.OsInfo
type Inventory_OsInfo struct {
	// The VM hostname.
	// +kcc:proto:field=google.cloud.osconfig.v1.Inventory.OsInfo.hostname
	Hostname *string `json:"hostname,omitempty"`

	// The operating system long name.
	//  For example 'Debian GNU/Linux 9' or 'Microsoft Window Server 2019
	//  Datacenter'.
	// +kcc:proto:field=google.cloud.osconfig.v1.Inventory.OsInfo.long_name
	LongName *string `json:"longName,omitempty"`

	// The operating system short name.
	//  For example, 'windows' or 'debian'.
	// +kcc:proto:field=google.cloud.osconfig.v1.Inventory.OsInfo.short_name
	ShortName *string `json:"shortName,omitempty"`

	// The version of the operating system.
	// +kcc:proto:field=google.cloud.osconfig.v1.Inventory.OsInfo.version
	Version *string `json:"version,omitempty"`

	// The system architecture of the operating system.
	// +kcc:proto:field=google.cloud.osconfig.v1.Inventory.OsInfo.architecture
	Architecture *string `json:"architecture,omitempty"`

	// The kernel version of the operating system.
	// +kcc:proto:field=google.cloud.osconfig.v1.Inventory.OsInfo.kernel_version
	KernelVersion *string `json:"kernelVersion,omitempty"`

	// The kernel release of the operating system.
	// +kcc:proto:field=google.cloud.osconfig.v1.Inventory.OsInfo.kernel_release
	KernelRelease *string `json:"kernelRelease,omitempty"`

	// The current version of the OS Config agent running on the VM.
	// +kcc:proto:field=google.cloud.osconfig.v1.Inventory.OsInfo.osconfig_agent_version
	OsconfigAgentVersion *string `json:"osconfigAgentVersion,omitempty"`
}

// +kcc:proto=google.cloud.osconfig.v1.Inventory.SoftwarePackage
type Inventory_SoftwarePackage struct {
	// Yum package info.
	//  For details about the yum package manager, see
	//  https://access.redhat.com/documentation/en-us/red_hat_enterprise_linux/6/html/deployment_guide/ch-yum.
	// +kcc:proto:field=google.cloud.osconfig.v1.Inventory.SoftwarePackage.yum_package
	YumPackage *Inventory_VersionedPackage `json:"yumPackage,omitempty"`

	// Details of an APT package.
	//  For details about the apt package manager, see
	//  https://wiki.debian.org/Apt.
	// +kcc:proto:field=google.cloud.osconfig.v1.Inventory.SoftwarePackage.apt_package
	AptPackage *Inventory_VersionedPackage `json:"aptPackage,omitempty"`

	// Details of a Zypper package.
	//  For details about the Zypper package manager, see
	//  https://en.opensuse.org/SDB:Zypper_manual.
	// +kcc:proto:field=google.cloud.osconfig.v1.Inventory.SoftwarePackage.zypper_package
	ZypperPackage *Inventory_VersionedPackage `json:"zypperPackage,omitempty"`

	// Details of a Googet package.
	//   For details about the googet package manager, see
	//   https://github.com/google/googet.
	// +kcc:proto:field=google.cloud.osconfig.v1.Inventory.SoftwarePackage.googet_package
	GoogetPackage *Inventory_VersionedPackage `json:"googetPackage,omitempty"`

	// Details of a Zypper patch.
	//  For details about the Zypper package manager, see
	//  https://en.opensuse.org/SDB:Zypper_manual.
	// +kcc:proto:field=google.cloud.osconfig.v1.Inventory.SoftwarePackage.zypper_patch
	ZypperPatch *Inventory_ZypperPatch `json:"zypperPatch,omitempty"`

	// Details of a Windows Update package.
	//  See https://docs.microsoft.com/en-us/windows/win32/api/_wua/ for
	//  information about Windows Update.
	// +kcc:proto:field=google.cloud.osconfig.v1.Inventory.SoftwarePackage.wua_package
	WuaPackage *Inventory_WindowsUpdatePackage `json:"wuaPackage,omitempty"`

	// Details of a Windows Quick Fix engineering package.
	//  See
	//  https://docs.microsoft.com/en-us/windows/win32/cimwin32prov/win32-quickfixengineering
	//  for info in Windows Quick Fix Engineering.
	// +kcc:proto:field=google.cloud.osconfig.v1.Inventory.SoftwarePackage.qfe_package
	QfePackage *Inventory_WindowsQuickFixEngineeringPackage `json:"qfePackage,omitempty"`

	// Details of a COS package.
	// +kcc:proto:field=google.cloud.osconfig.v1.Inventory.SoftwarePackage.cos_package
	CosPackage *Inventory_VersionedPackage `json:"cosPackage,omitempty"`

	// Details of Windows Application.
	// +kcc:proto:field=google.cloud.osconfig.v1.Inventory.SoftwarePackage.windows_application
	WindowsApplication *Inventory_WindowsApplication `json:"windowsApplication,omitempty"`
}

// +kcc:proto=google.cloud.osconfig.v1.Inventory.VersionedPackage
type Inventory_VersionedPackage struct {
	// The name of the package.
	// +kcc:proto:field=google.cloud.osconfig.v1.Inventory.VersionedPackage.package_name
	PackageName *string `json:"packageName,omitempty"`

	// The system architecture this package is intended for.
	// +kcc:proto:field=google.cloud.osconfig.v1.Inventory.VersionedPackage.architecture
	Architecture *string `json:"architecture,omitempty"`

	// The version of the package.
	// +kcc:proto:field=google.cloud.osconfig.v1.Inventory.VersionedPackage.version
	Version *string `json:"version,omitempty"`
}

// +kcc:proto=google.cloud.osconfig.v1.Inventory.WindowsApplication
type Inventory_WindowsApplication struct {
	// The name of the application or product.
	// +kcc:proto:field=google.cloud.osconfig.v1.Inventory.WindowsApplication.display_name
	DisplayName *string `json:"displayName,omitempty"`

	// The version of the product or application in string format.
	// +kcc:proto:field=google.cloud.osconfig.v1.Inventory.WindowsApplication.display_version
	DisplayVersion *string `json:"displayVersion,omitempty"`

	// The name of the manufacturer for the product or application.
	// +kcc:proto:field=google.cloud.osconfig.v1.Inventory.WindowsApplication.publisher
	Publisher *string `json:"publisher,omitempty"`

	// The last time this product received service. The value of this property
	//  is replaced each time a patch is applied or removed from the product or
	//  the command-line option is used to repair the product.
	// +kcc:proto:field=google.cloud.osconfig.v1.Inventory.WindowsApplication.install_date
	InstallDate *Date `json:"installDate,omitempty"`

	// The internet address for technical support.
	// +kcc:proto:field=google.cloud.osconfig.v1.Inventory.WindowsApplication.help_link
	HelpLink *string `json:"helpLink,omitempty"`
}

// +kcc:proto=google.cloud.osconfig.v1.Inventory.WindowsQuickFixEngineeringPackage
type Inventory_WindowsQuickFixEngineeringPackage struct {
	// A short textual description of the QFE update.
	// +kcc:proto:field=google.cloud.osconfig.v1.Inventory.WindowsQuickFixEngineeringPackage.caption
	Caption *string `json:"caption,omitempty"`

	// A textual description of the QFE update.
	// +kcc:proto:field=google.cloud.osconfig.v1.Inventory.WindowsQuickFixEngineeringPackage.description
	Description *string `json:"description,omitempty"`

	// Unique identifier associated with a particular QFE update.
	// +kcc:proto:field=google.cloud.osconfig.v1.Inventory.WindowsQuickFixEngineeringPackage.hot_fix_id
	HotFixID *string `json:"hotFixID,omitempty"`

	// Date that the QFE update was installed.  Mapped from installed_on field.
	// +kcc:proto:field=google.cloud.osconfig.v1.Inventory.WindowsQuickFixEngineeringPackage.install_time
	InstallTime *string `json:"installTime,omitempty"`
}

// +kcc:proto=google.cloud.osconfig.v1.Inventory.WindowsUpdatePackage
type Inventory_WindowsUpdatePackage struct {
	// The localized title of the update package.
	// +kcc:proto:field=google.cloud.osconfig.v1.Inventory.WindowsUpdatePackage.title
	Title *string `json:"title,omitempty"`

	// The localized description of the update package.
	// +kcc:proto:field=google.cloud.osconfig.v1.Inventory.WindowsUpdatePackage.description
	Description *string `json:"description,omitempty"`

	// The categories that are associated with this update package.
	// +kcc:proto:field=google.cloud.osconfig.v1.Inventory.WindowsUpdatePackage.categories
	Categories []Inventory_WindowsUpdatePackage_WindowsUpdateCategory `json:"categories,omitempty"`

	// A collection of Microsoft Knowledge Base article IDs that are associated
	//  with the update package.
	// +kcc:proto:field=google.cloud.osconfig.v1.Inventory.WindowsUpdatePackage.kb_article_ids
	KbArticleIds []string `json:"kbArticleIds,omitempty"`

	// A hyperlink to the language-specific support information for the update.
	// +kcc:proto:field=google.cloud.osconfig.v1.Inventory.WindowsUpdatePackage.support_url
	SupportURL *string `json:"supportURL,omitempty"`

	// A collection of URLs that provide more information about the update
	//  package.
	// +kcc:proto:field=google.cloud.osconfig.v1.Inventory.WindowsUpdatePackage.more_info_urls
	MoreInfoUrls []string `json:"moreInfoUrls,omitempty"`

	// Gets the identifier of an update package.  Stays the same across
	//  revisions.
	// +kcc:proto:field=google.cloud.osconfig.v1.Inventory.WindowsUpdatePackage.update_id
	UpdateID *string `json:"updateID,omitempty"`

	// The revision number of this update package.
	// +kcc:proto:field=google.cloud.osconfig.v1.Inventory.WindowsUpdatePackage.revision_number
	RevisionNumber *int32 `json:"revisionNumber,omitempty"`

	// The last published date of the update, in (UTC) date and time.
	// +kcc:proto:field=google.cloud.osconfig.v1.Inventory.WindowsUpdatePackage.last_deployment_change_time
	LastDeploymentChangeTime *string `json:"lastDeploymentChangeTime,omitempty"`
}

// +kcc:proto=google.cloud.osconfig.v1.Inventory.WindowsUpdatePackage.WindowsUpdateCategory
type Inventory_WindowsUpdatePackage_WindowsUpdateCategory struct {
	// The identifier of the windows update category.
	// +kcc:proto:field=google.cloud.osconfig.v1.Inventory.WindowsUpdatePackage.WindowsUpdateCategory.id
	ID *string `json:"id,omitempty"`

	// The name of the windows update category.
	// +kcc:proto:field=google.cloud.osconfig.v1.Inventory.WindowsUpdatePackage.WindowsUpdateCategory.name
	Name *string `json:"name,omitempty"`
}

// +kcc:proto=google.cloud.osconfig.v1.Inventory.ZypperPatch
type Inventory_ZypperPatch struct {
	// The name of the patch.
	// +kcc:proto:field=google.cloud.osconfig.v1.Inventory.ZypperPatch.patch_name
	PatchName *string `json:"patchName,omitempty"`

	// The category of the patch.
	// +kcc:proto:field=google.cloud.osconfig.v1.Inventory.ZypperPatch.category
	Category *string `json:"category,omitempty"`

	// The severity specified for this patch
	// +kcc:proto:field=google.cloud.osconfig.v1.Inventory.ZypperPatch.severity
	Severity *string `json:"severity,omitempty"`

	// Any summary information provided about this patch.
	// +kcc:proto:field=google.cloud.osconfig.v1.Inventory.ZypperPatch.summary
	Summary *string `json:"summary,omitempty"`
}

// +kcc:proto=google.type.Date
type Date struct {
	// Year of the date. Must be from 1 to 9999, or 0 to specify a date without
	//  a year.
	// +kcc:proto:field=google.type.Date.year
	Year *int32 `json:"year,omitempty"`

	// Month of a year. Must be from 1 to 12, or 0 to specify a year without a
	//  month and day.
	// +kcc:proto:field=google.type.Date.month
	Month *int32 `json:"month,omitempty"`

	// Day of a month. Must be from 1 to 31 and valid for the year and month, or 0
	//  to specify a year by itself or a year and month where the day isn't
	//  significant.
	// +kcc:proto:field=google.type.Date.day
	Day *int32 `json:"day,omitempty"`
}

// +kcc:proto=google.cloud.osconfig.v1.Inventory
type InventoryObservedState struct {
	// Output only. The `Inventory` API resource name.
	//
	//  Format:
	//  `projects/{project_number}/locations/{location}/instances/{instance_id}/inventory`
	// +kcc:proto:field=google.cloud.osconfig.v1.Inventory.name
	Name *string `json:"name,omitempty"`

	// Output only. Timestamp of the last reported inventory for the VM.
	// +kcc:proto:field=google.cloud.osconfig.v1.Inventory.update_time
	UpdateTime *string `json:"updateTime,omitempty"`
}
