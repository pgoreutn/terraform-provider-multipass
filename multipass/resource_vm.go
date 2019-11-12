package multipass

import (
	"fmt"
	"regexp"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/terraform-providers/terraform-provider-multipass/api/client"
)

func resourceItem() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"name": {
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: validateName,
			},
			"cpu": {
				Type:     schema.TypeInt,
				Required: true,
				ForceNew: true,
			},
			"memory": {
				Type:     schema.TypeInt,
				Required: true,
				ForceNew: true,
			},
			"disk": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
		},
		Create: create,
		Delete: delete,
		Read:   read,
	}
}

func validateName(v interface{}, k string) (ws []string, es []error) {
	var errs []error
	var warns []string
	value, ok := v.(string)
	if !ok {
		errs = append(errs, fmt.Errorf("Expected name to be string"))
		return warns, errs
	}
	whiteSpace := regexp.MustCompile(`\s+`)
	if whiteSpace.Match([]byte(value)) {
		errs = append(errs, fmt.Errorf("name cannot contain whitespace. Got %s", value))
		return warns, errs
	}
	return warns, errs
}

func create(d *schema.ResourceData, meta interface{}) error {

	cli := getClient(d)

	if err := cli.Exists(); err == nil {
		cli.Delete()
	}

	d.SetId(d.Get("name").(string))
	if err := cli.AddVm(); err != nil {
		return err
	}
	return nil

}

func delete(d *schema.ResourceData, m interface{}) error {
	cli := getClient(d)

	if err := cli.Exists(); err == nil {
		cli.Delete()
	}
	return nil
}

func read(d *schema.ResourceData, m interface{}) error {
	fmt.Println("READ")

	return nil
}

func getClient(d *schema.ResourceData) *client.Client {
	name := d.Get("name").(string)
	cpu := d.Get("cpu").(int)
	memory := d.Get("memory").(int)
	disk := d.Get("disk").(string)

	return client.NewClient(name, cpu, memory, disk)
}
