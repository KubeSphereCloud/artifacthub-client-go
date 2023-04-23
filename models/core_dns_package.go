// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// CoreDNSPackage CoreDNSPackage
//
// swagger:model CoreDNSPackage
type CoreDNSPackage struct {

	// all containers images whitelisted
	// Example: false
	AllContainersImagesWhitelisted bool `json:"all_containers_images_whitelisted,omitempty"`

	// app version
	// Example: 0.1.0
	AppVersion string `json:"app_version,omitempty"`

	// available versions
	// Required: true
	AvailableVersions []*AvailableVersion `json:"available_versions"`

	// capabilities
	// Example: seamless upgrades
	Capabilities string `json:"capabilities,omitempty"`

	// category
	Category PackageCategoryID `json:"category,omitempty"`

	// channels
	Channels []*Channel `json:"channels"`

	// cncf
	Cncf bool `json:"cncf,omitempty"`

	// containers images
	ContainersImages *ContainersImages `json:"containers_images,omitempty"`

	// contains security updates
	// Required: true
	ContainsSecurityUpdates *bool `json:"contains_security_updates"`

	// content url
	ContentURL string `json:"content_url,omitempty"`

	// default channel
	// Example: stable
	DefaultChannel string `json:"default_channel,omitempty"`

	// deprecated
	// Example: false
	Deprecated bool `json:"deprecated,omitempty"`

	// description
	// Example: This is a package sample
	Description string `json:"description,omitempty"`

	// display name
	// Example: Package 1
	DisplayName string `json:"display_name,omitempty"`

	// has changelog
	// Required: true
	HasChangelog *bool `json:"has_changelog"`

	// has values schema
	// Required: true
	HasValuesSchema *bool `json:"has_values_schema"`

	// home url
	// Example: http://home
	HomeURL string `json:"home_url,omitempty"`

	// is operator
	// Example: false
	IsOperator bool `json:"is_operator,omitempty"`

	// keywords
	// Example: ["key1","key2"]
	Keywords []string `json:"keywords"`

	// latest version
	// Example: 2.0.0
	// Required: true
	LatestVersion *string `json:"latest_version"`

	// license
	// Example: MIT
	License string `json:"license,omitempty"`

	// links
	Links []*Link `json:"links"`

	// logo image id
	// Example: 12345abcde
	LogoImageID string `json:"logo_image_id,omitempty"`

	// logo url
	// Example: http://logo
	LogoURL string `json:"logo_url,omitempty"`

	// maintainers
	Maintainers []*Maintainer `json:"maintainers"`

	// name
	// Example: pkg1
	// Required: true
	Name *string `json:"name"`

	// normalized name
	// Example: pkg1
	// Required: true
	NormalizedName *string `json:"normalized_name"`

	// official
	Official bool `json:"official,omitempty"`

	// package id
	// Required: true
	// Format: uuid
	PackageID *strfmt.UUID `json:"package_id"`

	// prerelease
	// Required: true
	Prerelease *bool `json:"prerelease"`

	// production organizations count
	ProductionOrganizationsCount float64 `json:"production_organizations_count,omitempty"`

	// provider
	// Example: Provider 1
	Provider string `json:"provider,omitempty"`

	// readme
	// Example: ###Readme
	Readme string `json:"readme,omitempty"`

	// recommendations
	Recommendations []*Recommendation `json:"recommendations"`

	// repository
	// Required: true
	Repository *RepositorySummary `json:"repository"`

	// security report created at
	SecurityReportCreatedAt int32 `json:"security_report_created_at,omitempty"`

	// security report summary
	SecurityReportSummary *SecurityReportSummary `json:"security_report_summary,omitempty"`

	// signatures
	Signatures []Signature `json:"signatures"`

	// signed
	// Example: false
	Signed bool `json:"signed,omitempty"`

	// stats
	Stats *Stats `json:"stats,omitempty"`

	// ts
	// Example: 1552082346
	// Required: true
	Ts *int32 `json:"ts"`

	// version
	// Example: 1.0.0
	// Required: true
	Version *string `json:"version"`
}

// Validate validates this core DNS package
func (m *CoreDNSPackage) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAvailableVersions(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCategory(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateChannels(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateContainersImages(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateContainsSecurityUpdates(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateHasChangelog(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateHasValuesSchema(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLatestVersion(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLinks(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMaintainers(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNormalizedName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePackageID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePrerelease(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRecommendations(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRepository(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSecurityReportSummary(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSignatures(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStats(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTs(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVersion(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CoreDNSPackage) validateAvailableVersions(formats strfmt.Registry) error {

	if err := validate.Required("available_versions", "body", m.AvailableVersions); err != nil {
		return err
	}

	for i := 0; i < len(m.AvailableVersions); i++ {
		if swag.IsZero(m.AvailableVersions[i]) { // not required
			continue
		}

		if m.AvailableVersions[i] != nil {
			if err := m.AvailableVersions[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("available_versions" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("available_versions" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *CoreDNSPackage) validateCategory(formats strfmt.Registry) error {
	if swag.IsZero(m.Category) { // not required
		return nil
	}

	if err := m.Category.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("category")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("category")
		}
		return err
	}

	return nil
}

func (m *CoreDNSPackage) validateChannels(formats strfmt.Registry) error {
	if swag.IsZero(m.Channels) { // not required
		return nil
	}

	for i := 0; i < len(m.Channels); i++ {
		if swag.IsZero(m.Channels[i]) { // not required
			continue
		}

		if m.Channels[i] != nil {
			if err := m.Channels[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("channels" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("channels" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *CoreDNSPackage) validateContainersImages(formats strfmt.Registry) error {
	if swag.IsZero(m.ContainersImages) { // not required
		return nil
	}

	if m.ContainersImages != nil {
		if err := m.ContainersImages.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("containers_images")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("containers_images")
			}
			return err
		}
	}

	return nil
}

func (m *CoreDNSPackage) validateContainsSecurityUpdates(formats strfmt.Registry) error {

	if err := validate.Required("contains_security_updates", "body", m.ContainsSecurityUpdates); err != nil {
		return err
	}

	return nil
}

func (m *CoreDNSPackage) validateHasChangelog(formats strfmt.Registry) error {

	if err := validate.Required("has_changelog", "body", m.HasChangelog); err != nil {
		return err
	}

	return nil
}

func (m *CoreDNSPackage) validateHasValuesSchema(formats strfmt.Registry) error {

	if err := validate.Required("has_values_schema", "body", m.HasValuesSchema); err != nil {
		return err
	}

	return nil
}

func (m *CoreDNSPackage) validateLatestVersion(formats strfmt.Registry) error {

	if err := validate.Required("latest_version", "body", m.LatestVersion); err != nil {
		return err
	}

	return nil
}

func (m *CoreDNSPackage) validateLinks(formats strfmt.Registry) error {
	if swag.IsZero(m.Links) { // not required
		return nil
	}

	for i := 0; i < len(m.Links); i++ {
		if swag.IsZero(m.Links[i]) { // not required
			continue
		}

		if m.Links[i] != nil {
			if err := m.Links[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("links" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("links" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *CoreDNSPackage) validateMaintainers(formats strfmt.Registry) error {
	if swag.IsZero(m.Maintainers) { // not required
		return nil
	}

	for i := 0; i < len(m.Maintainers); i++ {
		if swag.IsZero(m.Maintainers[i]) { // not required
			continue
		}

		if m.Maintainers[i] != nil {
			if err := m.Maintainers[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("maintainers" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("maintainers" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *CoreDNSPackage) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *CoreDNSPackage) validateNormalizedName(formats strfmt.Registry) error {

	if err := validate.Required("normalized_name", "body", m.NormalizedName); err != nil {
		return err
	}

	return nil
}

func (m *CoreDNSPackage) validatePackageID(formats strfmt.Registry) error {

	if err := validate.Required("package_id", "body", m.PackageID); err != nil {
		return err
	}

	if err := validate.FormatOf("package_id", "body", "uuid", m.PackageID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *CoreDNSPackage) validatePrerelease(formats strfmt.Registry) error {

	if err := validate.Required("prerelease", "body", m.Prerelease); err != nil {
		return err
	}

	return nil
}

func (m *CoreDNSPackage) validateRecommendations(formats strfmt.Registry) error {
	if swag.IsZero(m.Recommendations) { // not required
		return nil
	}

	for i := 0; i < len(m.Recommendations); i++ {
		if swag.IsZero(m.Recommendations[i]) { // not required
			continue
		}

		if m.Recommendations[i] != nil {
			if err := m.Recommendations[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("recommendations" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("recommendations" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *CoreDNSPackage) validateRepository(formats strfmt.Registry) error {

	if err := validate.Required("repository", "body", m.Repository); err != nil {
		return err
	}

	if m.Repository != nil {
		if err := m.Repository.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("repository")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("repository")
			}
			return err
		}
	}

	return nil
}

func (m *CoreDNSPackage) validateSecurityReportSummary(formats strfmt.Registry) error {
	if swag.IsZero(m.SecurityReportSummary) { // not required
		return nil
	}

	if m.SecurityReportSummary != nil {
		if err := m.SecurityReportSummary.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("security_report_summary")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("security_report_summary")
			}
			return err
		}
	}

	return nil
}

func (m *CoreDNSPackage) validateSignatures(formats strfmt.Registry) error {
	if swag.IsZero(m.Signatures) { // not required
		return nil
	}

	for i := 0; i < len(m.Signatures); i++ {

		if err := m.Signatures[i].Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("signatures" + "." + strconv.Itoa(i))
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("signatures" + "." + strconv.Itoa(i))
			}
			return err
		}

	}

	return nil
}

func (m *CoreDNSPackage) validateStats(formats strfmt.Registry) error {
	if swag.IsZero(m.Stats) { // not required
		return nil
	}

	if m.Stats != nil {
		if err := m.Stats.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("stats")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("stats")
			}
			return err
		}
	}

	return nil
}

func (m *CoreDNSPackage) validateTs(formats strfmt.Registry) error {

	if err := validate.Required("ts", "body", m.Ts); err != nil {
		return err
	}

	return nil
}

func (m *CoreDNSPackage) validateVersion(formats strfmt.Registry) error {

	if err := validate.Required("version", "body", m.Version); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this core DNS package based on the context it is used
func (m *CoreDNSPackage) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAvailableVersions(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateCategory(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateChannels(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateContainersImages(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateLinks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateMaintainers(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateRecommendations(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateRepository(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSecurityReportSummary(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSignatures(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateStats(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CoreDNSPackage) contextValidateAvailableVersions(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.AvailableVersions); i++ {

		if m.AvailableVersions[i] != nil {
			if err := m.AvailableVersions[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("available_versions" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("available_versions" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *CoreDNSPackage) contextValidateCategory(ctx context.Context, formats strfmt.Registry) error {

	if err := m.Category.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("category")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("category")
		}
		return err
	}

	return nil
}

func (m *CoreDNSPackage) contextValidateChannels(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Channels); i++ {

		if m.Channels[i] != nil {
			if err := m.Channels[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("channels" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("channels" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *CoreDNSPackage) contextValidateContainersImages(ctx context.Context, formats strfmt.Registry) error {

	if m.ContainersImages != nil {
		if err := m.ContainersImages.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("containers_images")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("containers_images")
			}
			return err
		}
	}

	return nil
}

func (m *CoreDNSPackage) contextValidateLinks(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Links); i++ {

		if m.Links[i] != nil {
			if err := m.Links[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("links" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("links" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *CoreDNSPackage) contextValidateMaintainers(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Maintainers); i++ {

		if m.Maintainers[i] != nil {
			if err := m.Maintainers[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("maintainers" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("maintainers" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *CoreDNSPackage) contextValidateRecommendations(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Recommendations); i++ {

		if m.Recommendations[i] != nil {
			if err := m.Recommendations[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("recommendations" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("recommendations" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *CoreDNSPackage) contextValidateRepository(ctx context.Context, formats strfmt.Registry) error {

	if m.Repository != nil {
		if err := m.Repository.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("repository")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("repository")
			}
			return err
		}
	}

	return nil
}

func (m *CoreDNSPackage) contextValidateSecurityReportSummary(ctx context.Context, formats strfmt.Registry) error {

	if m.SecurityReportSummary != nil {
		if err := m.SecurityReportSummary.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("security_report_summary")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("security_report_summary")
			}
			return err
		}
	}

	return nil
}

func (m *CoreDNSPackage) contextValidateSignatures(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Signatures); i++ {

		if err := m.Signatures[i].ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("signatures" + "." + strconv.Itoa(i))
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("signatures" + "." + strconv.Itoa(i))
			}
			return err
		}

	}

	return nil
}

func (m *CoreDNSPackage) contextValidateStats(ctx context.Context, formats strfmt.Registry) error {

	if m.Stats != nil {
		if err := m.Stats.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("stats")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("stats")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *CoreDNSPackage) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CoreDNSPackage) UnmarshalBinary(b []byte) error {
	var res CoreDNSPackage
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
