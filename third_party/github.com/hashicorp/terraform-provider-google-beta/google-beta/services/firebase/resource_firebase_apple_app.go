// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    Type: MMv1     ***
//
// ----------------------------------------------------------------------------
//
//     This file is automatically generated by Magic Modules and manual
//     changes will be clobbered when the file is regenerated.
//
//     Please read more about how to change this file in
//     .github/CONTRIBUTING.md.
//
// ----------------------------------------------------------------------------

package firebase

import (
	"fmt"
	"log"
	"reflect"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
)

func ResourceFirebaseAppleApp() *schema.Resource {
	return &schema.Resource{
		Create: resourceFirebaseAppleAppCreate,
		Read:   resourceFirebaseAppleAppRead,
		Update: resourceFirebaseAppleAppUpdate,
		Delete: resourceFirebaseAppleAppDelete,

		Importer: &schema.ResourceImporter{
			State: resourceFirebaseAppleAppImport,
		},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(20 * time.Minute),
			Update: schema.DefaultTimeout(20 * time.Minute),
			Delete: schema.DefaultTimeout(20 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"bundle_id": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: `The canonical bundle ID of the Apple app as it would appear in the Apple AppStore.`,
			},
			"display_name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: `The user-assigned display name of the App.`,
			},
			"api_key_id": {
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
				Description: `The globally unique, Google-assigned identifier (UID) for the Firebase API key associated with the AppleApp.
If apiKeyId is not set during creation, then Firebase automatically associates an apiKeyId with the AppleApp.
This auto-associated key may be an existing valid key or, if no valid key exists, a new one will be provisioned.`,
			},
			"app_store_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: `The automatically generated Apple ID assigned to the Apple app by Apple in the Apple App Store.`,
			},
			"team_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: `The Apple Developer Team ID associated with the App in the App Store.`,
			},
			"app_id": {
				Type:     schema.TypeString,
				Computed: true,
				Description: `The globally unique, Firebase-assigned identifier of the App.
This identifier should be treated as an opaque token, as the data format is not specified.`,
			},
			"name": {
				Type:     schema.TypeString,
				Computed: true,
				Description: `The fully qualified resource name of the App, for example:
projects/projectId/iosApps/appId`,
			},
			"deletion_policy": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "DELETE",
				Description: `(Optional) Set to 'ABANDON' to allow the Apple to be untracked from terraform state
rather than deleted upon 'terraform destroy'. This is useful because the Apple may be
serving traffic. Set to 'DELETE' to delete the Apple. Defaults to 'DELETE'.`,
			},
			"project": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
		},
		UseJSONNumber: true,
	}
}

func resourceFirebaseAppleAppCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	obj := make(map[string]interface{})
	displayNameProp, err := expandFirebaseAppleAppDisplayName(d.Get("display_name"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("display_name"); !tpgresource.IsEmptyValue(reflect.ValueOf(displayNameProp)) && (ok || !reflect.DeepEqual(v, displayNameProp)) {
		obj["displayName"] = displayNameProp
	}
	bundleIdProp, err := expandFirebaseAppleAppBundleId(d.Get("bundle_id"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("bundle_id"); !tpgresource.IsEmptyValue(reflect.ValueOf(bundleIdProp)) && (ok || !reflect.DeepEqual(v, bundleIdProp)) {
		obj["bundleId"] = bundleIdProp
	}
	appStoreIdProp, err := expandFirebaseAppleAppAppStoreId(d.Get("app_store_id"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("app_store_id"); !tpgresource.IsEmptyValue(reflect.ValueOf(appStoreIdProp)) && (ok || !reflect.DeepEqual(v, appStoreIdProp)) {
		obj["appStoreId"] = appStoreIdProp
	}
	teamIdProp, err := expandFirebaseAppleAppTeamId(d.Get("team_id"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("team_id"); !tpgresource.IsEmptyValue(reflect.ValueOf(teamIdProp)) && (ok || !reflect.DeepEqual(v, teamIdProp)) {
		obj["teamId"] = teamIdProp
	}
	apiKeyIdProp, err := expandFirebaseAppleAppApiKeyId(d.Get("api_key_id"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("api_key_id"); !tpgresource.IsEmptyValue(reflect.ValueOf(apiKeyIdProp)) && (ok || !reflect.DeepEqual(v, apiKeyIdProp)) {
		obj["apiKeyId"] = apiKeyIdProp
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{FirebaseBasePath}}projects/{{project}}/iosApps")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Creating new AppleApp: %#v", obj)
	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for AppleApp: %s", err)
	}
	billingProject = project

	// err == nil indicates that the billing_project value was found
	if bp, err := tpgresource.GetBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := transport_tpg.SendRequest(transport_tpg.SendRequestOptions{
		Config:    config,
		Method:    "POST",
		Project:   billingProject,
		RawURL:    url,
		UserAgent: userAgent,
		Body:      obj,
		Timeout:   d.Timeout(schema.TimeoutCreate),
	})
	if err != nil {
		return fmt.Errorf("Error creating AppleApp: %s", err)
	}

	// Store the ID now
	id, err := tpgresource.ReplaceVars(d, config, "projects/{{project}}/iosApps/{{app_id}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	// Use the resource in the operation response to populate
	// identity fields and d.Id() before read
	var opRes map[string]interface{}
	err = FirebaseOperationWaitTimeWithResponse(
		config, res, &opRes, project, "Creating AppleApp", userAgent,
		d.Timeout(schema.TimeoutCreate))
	if err != nil {
		// The resource didn't actually create
		d.SetId("")

		return fmt.Errorf("Error waiting to create AppleApp: %s", err)
	}

	if err := d.Set("name", flattenFirebaseAppleAppName(opRes["name"], d, config)); err != nil {
		return err
	}
	if err := d.Set("app_id", flattenFirebaseAppleAppAppId(opRes["appId"], d, config)); err != nil {
		return err
	}

	// This may have caused the ID to update - update it if so.
	id, err = tpgresource.ReplaceVars(d, config, "projects/{{project}}/iosApps/{{app_id}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	log.Printf("[DEBUG] Finished creating AppleApp %q: %#v", d.Id(), res)

	return resourceFirebaseAppleAppRead(d, meta)
}

func resourceFirebaseAppleAppRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{FirebaseBasePath}}projects/{{project}}/iosApps/{{app_id}}")
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for AppleApp: %s", err)
	}
	billingProject = project

	// err == nil indicates that the billing_project value was found
	if bp, err := tpgresource.GetBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := transport_tpg.SendRequest(transport_tpg.SendRequestOptions{
		Config:    config,
		Method:    "GET",
		Project:   billingProject,
		RawURL:    url,
		UserAgent: userAgent,
	})
	if err != nil {
		return transport_tpg.HandleNotFoundError(err, d, fmt.Sprintf("FirebaseAppleApp %q", d.Id()))
	}

	// Explicitly set virtual fields to default values if unset
	if _, ok := d.GetOkExists("deletion_policy"); !ok {
		if err := d.Set("deletion_policy", "DELETE"); err != nil {
			return fmt.Errorf("Error setting deletion_policy: %s", err)
		}
	}
	if err := d.Set("project", project); err != nil {
		return fmt.Errorf("Error reading AppleApp: %s", err)
	}

	if err := d.Set("name", flattenFirebaseAppleAppName(res["name"], d, config)); err != nil {
		return fmt.Errorf("Error reading AppleApp: %s", err)
	}
	if err := d.Set("display_name", flattenFirebaseAppleAppDisplayName(res["displayName"], d, config)); err != nil {
		return fmt.Errorf("Error reading AppleApp: %s", err)
	}
	if err := d.Set("app_id", flattenFirebaseAppleAppAppId(res["appId"], d, config)); err != nil {
		return fmt.Errorf("Error reading AppleApp: %s", err)
	}
	if err := d.Set("bundle_id", flattenFirebaseAppleAppBundleId(res["bundleId"], d, config)); err != nil {
		return fmt.Errorf("Error reading AppleApp: %s", err)
	}
	if err := d.Set("app_store_id", flattenFirebaseAppleAppAppStoreId(res["appStoreId"], d, config)); err != nil {
		return fmt.Errorf("Error reading AppleApp: %s", err)
	}
	if err := d.Set("team_id", flattenFirebaseAppleAppTeamId(res["teamId"], d, config)); err != nil {
		return fmt.Errorf("Error reading AppleApp: %s", err)
	}
	if err := d.Set("api_key_id", flattenFirebaseAppleAppApiKeyId(res["apiKeyId"], d, config)); err != nil {
		return fmt.Errorf("Error reading AppleApp: %s", err)
	}

	return nil
}

func resourceFirebaseAppleAppUpdate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for AppleApp: %s", err)
	}
	billingProject = project

	obj := make(map[string]interface{})
	displayNameProp, err := expandFirebaseAppleAppDisplayName(d.Get("display_name"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("display_name"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, displayNameProp)) {
		obj["displayName"] = displayNameProp
	}
	appStoreIdProp, err := expandFirebaseAppleAppAppStoreId(d.Get("app_store_id"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("app_store_id"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, appStoreIdProp)) {
		obj["appStoreId"] = appStoreIdProp
	}
	teamIdProp, err := expandFirebaseAppleAppTeamId(d.Get("team_id"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("team_id"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, teamIdProp)) {
		obj["teamId"] = teamIdProp
	}
	apiKeyIdProp, err := expandFirebaseAppleAppApiKeyId(d.Get("api_key_id"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("api_key_id"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, apiKeyIdProp)) {
		obj["apiKeyId"] = apiKeyIdProp
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{FirebaseBasePath}}projects/{{project}}/iosApps/{{app_id}}")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Updating AppleApp %q: %#v", d.Id(), obj)
	updateMask := []string{}

	if d.HasChange("display_name") {
		updateMask = append(updateMask, "displayName")
	}

	if d.HasChange("app_store_id") {
		updateMask = append(updateMask, "appStoreId")
	}

	if d.HasChange("team_id") {
		updateMask = append(updateMask, "teamId")
	}

	if d.HasChange("api_key_id") {
		updateMask = append(updateMask, "apiKeyId")
	}
	// updateMask is a URL parameter but not present in the schema, so ReplaceVars
	// won't set it
	url, err = transport_tpg.AddQueryParams(url, map[string]string{"updateMask": strings.Join(updateMask, ",")})
	if err != nil {
		return err
	}

	// err == nil indicates that the billing_project value was found
	if bp, err := tpgresource.GetBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := transport_tpg.SendRequest(transport_tpg.SendRequestOptions{
		Config:    config,
		Method:    "PATCH",
		Project:   billingProject,
		RawURL:    url,
		UserAgent: userAgent,
		Body:      obj,
		Timeout:   d.Timeout(schema.TimeoutUpdate),
	})

	if err != nil {
		return fmt.Errorf("Error updating AppleApp %q: %s", d.Id(), err)
	} else {
		log.Printf("[DEBUG] Finished updating AppleApp %q: %#v", d.Id(), res)
	}

	return resourceFirebaseAppleAppRead(d, meta)
}

func resourceFirebaseAppleAppDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	// Handwritten
	obj := make(map[string]interface{})
	if d.Get("deletion_policy") == "DELETE" {
		obj["immediate"] = true
	} else {
		fmt.Printf("Skip deleting App %q due to deletion_policy: %q\n", d.Id(), d.Get("deletion_policy"))
		return nil
	}
	// End of Handwritten
	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for App: %s", err)
	}
	billingProject = project

	url, err := tpgresource.ReplaceVars(d, config, "{{FirebaseBasePath}}{{name}}:remove")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Deleting App %q", d.Id())

	// err == nil indicates that the billing_project value was found
	if bp, err := tpgresource.GetBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := transport_tpg.SendRequest(transport_tpg.SendRequestOptions{
		Config:    config,
		Method:    "POST",
		Project:   billingProject,
		RawURL:    url,
		UserAgent: userAgent,
		Body:      obj,
		Timeout:   d.Timeout(schema.TimeoutDelete),
	})
	if err != nil {
		return transport_tpg.HandleNotFoundError(err, d, "App")
	}

	err = FirebaseOperationWaitTime(
		config, res, project, "Deleting App", userAgent,
		d.Timeout(schema.TimeoutDelete))

	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Finished deleting App %q: %#v", d.Id(), res)

	// This is useful if the Delete operation returns before the Get operation
	// during post-test destroy shows the completed state of the resource.
	time.Sleep(5 * time.Second)

	return nil
}

func resourceFirebaseAppleAppImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	config := meta.(*transport_tpg.Config)
	if err := tpgresource.ParseImportId([]string{
		"(?P<project>[^/]+) projects/(?P<project>[^/]+)/iosApps/(?P<app_id>[^/]+)",
		"projects/(?P<project>[^/]+)/iosApps/(?P<app_id>[^/]+)",
		"(?P<project>[^/]+)/(?P<project>[^/]+)/(?P<app_id>[^/]+)",
		"iosApps/(?P<app_id>[^/]+)",
		"(?P<app_id>[^/]+)",
	}, d, config); err != nil {
		return nil, err
	}

	// Replace import id for the resource id
	id, err := tpgresource.ReplaceVars(d, config, "projects/{{project}}/iosApps/{{app_id}}")
	if err != nil {
		return nil, fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	// Explicitly set virtual fields to default values on import
	if err := d.Set("deletion_policy", "DELETE"); err != nil {
		return nil, fmt.Errorf("Error setting deletion_policy: %s", err)
	}

	return []*schema.ResourceData{d}, nil
}

func flattenFirebaseAppleAppName(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenFirebaseAppleAppDisplayName(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenFirebaseAppleAppAppId(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenFirebaseAppleAppBundleId(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenFirebaseAppleAppAppStoreId(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenFirebaseAppleAppTeamId(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenFirebaseAppleAppApiKeyId(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func expandFirebaseAppleAppDisplayName(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandFirebaseAppleAppBundleId(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandFirebaseAppleAppAppStoreId(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandFirebaseAppleAppTeamId(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandFirebaseAppleAppApiKeyId(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}