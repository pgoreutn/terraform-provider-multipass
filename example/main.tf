
resource "multipass_vm" "node_master" {
  name    ="node-master-${count.index}"
  cpu     = 2
  memory  = 1024
  disk    = "1G"
  count = 3
}
