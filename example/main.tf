
resource "multipass_vm" "node_master" {
  name    ="master"
  cpu     = 2
  memory  = 1024
  disk    = "10G"
}
