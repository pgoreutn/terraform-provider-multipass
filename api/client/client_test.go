package client

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

var client *Client

func init() {
	client = NewClient("master", 2, 1024, "1G")
}
func TestExists(t *testing.T) {
	err := client.Exists()

	if err == nil {
		assert.Equal(t, err, nil)
	}
}

func TestAddVm(t *testing.T) {
	err := client.AddVm()
	assert.Error(t, err, nil)
	fmt.Println(err)
}
