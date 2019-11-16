package multipass

import (
	"fmt"
	uuid "github.com/nu7hatch/gouuid"
	"log"
	"regexp"

	"github.com/hashicorp/terraform/helper/schema"
	cli "github.com/terraform-providers/terraform-provider-multipass/api/client"
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

	vms := meta.(*cli.VMClient).NewClient(d)

	if err:=vms.AddVm();err!=nil{
		return err
	}

	id,err:=uuid.NewV4()
	if err!=nil{
		log.Fatalf("uuid.NewV4() failed with %s\n", err)
	}

	d.SetId(id.String())

	return nil
}

func delete(d *schema.ResourceData, meta interface{}) error {
	vms := meta.(*cli.VMClient).NewClient(d)

	if err:=vms.Delete();err!=nil{
		return err
	}

	return nil
}

func read(d *schema.ResourceData, m interface{}) error {
	return nil
}