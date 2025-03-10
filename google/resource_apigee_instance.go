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

package google

import (
	"fmt"
	"log"
	"reflect"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceApigeeInstance() *schema.Resource {
	return &schema.Resource{
		Create: resourceApigeeInstanceCreate,
		Read:   resourceApigeeInstanceRead,
		Delete: resourceApigeeInstanceDelete,

		Importer: &schema.ResourceImporter{
			State: resourceApigeeInstanceImport,
		},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(60 * time.Minute),
			Delete: schema.DefaultTimeout(60 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"location": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
				Description: `Compute Engine location where the instance resides. For trial organization
subscriptions, the location must be a Compute Engine zone. For paid organization
subscriptions, it should correspond to a Compute Engine region.`,
			},
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: `Resource ID of the instance.`,
			},
			"org_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
				Description: `The Apigee Organization associated with the Apigee instance,
in the format 'organizations/{{org_name}}'.`,
			},
			"description": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: `Description of the instance.`,
			},
			"disk_encryption_key_name": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
				Description: `Customer Managed Encryption Key (CMEK) used for disk and volume encryption. Required for Apigee paid subscriptions only.
Use the following format: 'projects/([^/]+)/locations/([^/]+)/keyRings/([^/]+)/cryptoKeys/([^/]+)'`,
			},
			"display_name": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: `Display name of the instance.`,
			},
			"ip_range": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
				Description: `IP range represents the customer-provided CIDR block of length 22 that will be used for
the Apigee instance creation. This optional range, if provided, should be freely
available as part of larger named range the customer has allocated to the Service
Networking peering. If this is not provided, Apigee will automatically request for any
available /22 CIDR block from Service Networking. The customer should use this CIDR block
for configuring their firewall needs to allow traffic from Apigee.
Input format: "a.b.c.d/22"`,
			},
			"peering_cidr_range": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
				Description: `The size of the CIDR block range that will be reserved by the instance. For valid values, 
see [CidrRange](https://cloud.google.com/apigee/docs/reference/apis/apigee/rest/v1/organizations.instances#CidrRange) on the documentation.`,
			},
			"host": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `Output only. Hostname or IP address of the exposed Apigee endpoint used by clients to connect to the service.`,
			},
			"port": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `Output only. Port number of the exposed Apigee endpoint.`,
			},
		},
		UseJSONNumber: true,
	}
}

func resourceApigeeInstanceCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	userAgent, err := generateUserAgentString(d, config.userAgent)
	if err != nil {
		return err
	}

	obj := make(map[string]interface{})
	nameProp, err := expandApigeeInstanceName(d.Get("name"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("name"); !isEmptyValue(reflect.ValueOf(nameProp)) && (ok || !reflect.DeepEqual(v, nameProp)) {
		obj["name"] = nameProp
	}
	locationProp, err := expandApigeeInstanceLocation(d.Get("location"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("location"); !isEmptyValue(reflect.ValueOf(locationProp)) && (ok || !reflect.DeepEqual(v, locationProp)) {
		obj["location"] = locationProp
	}
	peeringCidrRangeProp, err := expandApigeeInstancePeeringCidrRange(d.Get("peering_cidr_range"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("peering_cidr_range"); !isEmptyValue(reflect.ValueOf(peeringCidrRangeProp)) && (ok || !reflect.DeepEqual(v, peeringCidrRangeProp)) {
		obj["peeringCidrRange"] = peeringCidrRangeProp
	}
	ipRangeProp, err := expandApigeeInstanceIpRange(d.Get("ip_range"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("ip_range"); !isEmptyValue(reflect.ValueOf(ipRangeProp)) && (ok || !reflect.DeepEqual(v, ipRangeProp)) {
		obj["ipRange"] = ipRangeProp
	}
	descriptionProp, err := expandApigeeInstanceDescription(d.Get("description"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("description"); !isEmptyValue(reflect.ValueOf(descriptionProp)) && (ok || !reflect.DeepEqual(v, descriptionProp)) {
		obj["description"] = descriptionProp
	}
	displayNameProp, err := expandApigeeInstanceDisplayName(d.Get("display_name"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("display_name"); !isEmptyValue(reflect.ValueOf(displayNameProp)) && (ok || !reflect.DeepEqual(v, displayNameProp)) {
		obj["displayName"] = displayNameProp
	}
	diskEncryptionKeyNameProp, err := expandApigeeInstanceDiskEncryptionKeyName(d.Get("disk_encryption_key_name"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("disk_encryption_key_name"); !isEmptyValue(reflect.ValueOf(diskEncryptionKeyNameProp)) && (ok || !reflect.DeepEqual(v, diskEncryptionKeyNameProp)) {
		obj["diskEncryptionKeyName"] = diskEncryptionKeyNameProp
	}

	url, err := replaceVars(d, config, "{{ApigeeBasePath}}{{org_id}}/instances")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Creating new Instance: %#v", obj)
	billingProject := ""

	// err == nil indicates that the billing_project value was found
	if bp, err := getBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := sendRequestWithTimeout(config, "POST", billingProject, url, userAgent, obj, d.Timeout(schema.TimeoutCreate))
	if err != nil {
		return fmt.Errorf("Error creating Instance: %s", err)
	}

	// Store the ID now
	id, err := replaceVars(d, config, "{{org_id}}/instances/{{name}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	// Use the resource in the operation response to populate
	// identity fields and d.Id() before read
	var opRes map[string]interface{}
	err = apigeeOperationWaitTimeWithResponse(
		config, res, &opRes, "Creating Instance", userAgent,
		d.Timeout(schema.TimeoutCreate))
	if err != nil {
		// The resource didn't actually create
		d.SetId("")
		return fmt.Errorf("Error waiting to create Instance: %s", err)
	}

	if err := d.Set("name", flattenApigeeInstanceName(opRes["name"], d, config)); err != nil {
		return err
	}

	// This may have caused the ID to update - update it if so.
	id, err = replaceVars(d, config, "{{org_id}}/instances/{{name}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	log.Printf("[DEBUG] Finished creating Instance %q: %#v", d.Id(), res)

	return resourceApigeeInstanceRead(d, meta)
}

func resourceApigeeInstanceRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	userAgent, err := generateUserAgentString(d, config.userAgent)
	if err != nil {
		return err
	}

	url, err := replaceVars(d, config, "{{ApigeeBasePath}}{{org_id}}/instances/{{name}}")
	if err != nil {
		return err
	}

	billingProject := ""

	// err == nil indicates that the billing_project value was found
	if bp, err := getBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := sendRequest(config, "GET", billingProject, url, userAgent, nil)
	if err != nil {
		return handleNotFoundError(err, d, fmt.Sprintf("ApigeeInstance %q", d.Id()))
	}

	if err := d.Set("name", flattenApigeeInstanceName(res["name"], d, config)); err != nil {
		return fmt.Errorf("Error reading Instance: %s", err)
	}
	if err := d.Set("location", flattenApigeeInstanceLocation(res["location"], d, config)); err != nil {
		return fmt.Errorf("Error reading Instance: %s", err)
	}
	if err := d.Set("peering_cidr_range", flattenApigeeInstancePeeringCidrRange(res["peeringCidrRange"], d, config)); err != nil {
		return fmt.Errorf("Error reading Instance: %s", err)
	}
	if err := d.Set("ip_range", flattenApigeeInstanceIpRange(res["ipRange"], d, config)); err != nil {
		return fmt.Errorf("Error reading Instance: %s", err)
	}
	if err := d.Set("description", flattenApigeeInstanceDescription(res["description"], d, config)); err != nil {
		return fmt.Errorf("Error reading Instance: %s", err)
	}
	if err := d.Set("display_name", flattenApigeeInstanceDisplayName(res["displayName"], d, config)); err != nil {
		return fmt.Errorf("Error reading Instance: %s", err)
	}
	if err := d.Set("disk_encryption_key_name", flattenApigeeInstanceDiskEncryptionKeyName(res["diskEncryptionKeyName"], d, config)); err != nil {
		return fmt.Errorf("Error reading Instance: %s", err)
	}
	if err := d.Set("host", flattenApigeeInstanceHost(res["host"], d, config)); err != nil {
		return fmt.Errorf("Error reading Instance: %s", err)
	}
	if err := d.Set("port", flattenApigeeInstancePort(res["port"], d, config)); err != nil {
		return fmt.Errorf("Error reading Instance: %s", err)
	}

	return nil
}

func resourceApigeeInstanceDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	userAgent, err := generateUserAgentString(d, config.userAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	url, err := replaceVars(d, config, "{{ApigeeBasePath}}{{org_id}}/instances/{{name}}")
	if err != nil {
		return err
	}

	var obj map[string]interface{}
	log.Printf("[DEBUG] Deleting Instance %q", d.Id())

	// err == nil indicates that the billing_project value was found
	if bp, err := getBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := sendRequestWithTimeout(config, "DELETE", billingProject, url, userAgent, obj, d.Timeout(schema.TimeoutDelete))
	if err != nil {
		return handleNotFoundError(err, d, "Instance")
	}

	err = apigeeOperationWaitTime(
		config, res, "Deleting Instance", userAgent,
		d.Timeout(schema.TimeoutDelete))

	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Finished deleting Instance %q: %#v", d.Id(), res)
	return nil
}

func resourceApigeeInstanceImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	config := meta.(*Config)

	// current import_formats cannot import fields with forward slashes in their value
	if err := parseImportId([]string{"(?P<name>.+)"}, d, config); err != nil {
		return nil, err
	}

	nameParts := strings.Split(d.Get("name").(string), "/")
	if len(nameParts) == 4 {
		// `organizations/{{org_name}}/instances/{{name}}`
		orgId := fmt.Sprintf("organizations/%s", nameParts[1])
		if err := d.Set("org_id", orgId); err != nil {
			return nil, fmt.Errorf("Error setting org_id: %s", err)
		}
		if err := d.Set("name", nameParts[3]); err != nil {
			return nil, fmt.Errorf("Error setting name: %s", err)
		}
	} else if len(nameParts) == 3 {
		// `organizations/{{org_name}}/{{name}}`
		orgId := fmt.Sprintf("organizations/%s", nameParts[1])
		if err := d.Set("org_id", orgId); err != nil {
			return nil, fmt.Errorf("Error setting org_id: %s", err)
		}
		if err := d.Set("name", nameParts[2]); err != nil {
			return nil, fmt.Errorf("Error setting name: %s", err)
		}
	} else {
		return nil, fmt.Errorf(
			"Saw %s when the name is expected to have shape %s or %s",
			d.Get("name"),
			"organizations/{{org_name}}/instances/{{name}}",
			"organizations/{{org_name}}/{{name}}")
	}

	// Replace import id for the resource id
	id, err := replaceVars(d, config, "{{org_id}}/instances/{{name}}")
	if err != nil {
		return nil, fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	return []*schema.ResourceData{d}, nil
}

func flattenApigeeInstanceName(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenApigeeInstanceLocation(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenApigeeInstancePeeringCidrRange(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenApigeeInstanceIpRange(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenApigeeInstanceDescription(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenApigeeInstanceDisplayName(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenApigeeInstanceDiskEncryptionKeyName(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenApigeeInstanceHost(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenApigeeInstancePort(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func expandApigeeInstanceName(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandApigeeInstanceLocation(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandApigeeInstancePeeringCidrRange(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandApigeeInstanceIpRange(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandApigeeInstanceDescription(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandApigeeInstanceDisplayName(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandApigeeInstanceDiskEncryptionKeyName(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}
