data "external" "master" {
  program = ["python3", "${path.module}/script/multipass.py" ]
  query = {
    name = "${var.master}"
    cpu  = "${var.cpu}"
    mem  = "${var.mem}"
    disk = "${var.disk}"
    init = "${data.template_file.cloud_init_master.rendered}"
  }
}