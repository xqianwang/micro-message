resource "null_resource" "openShiftConfig" {
  connection {
    type        = "ssh"
    user        = "${var.adminUsername}"
    host        = "${module.openshift-cluster.bastion-public_ip}"
    private_key = "${file("${var.privateKeyFilename}")}"
  }

  provisioner "file" {
    source      = "${var.privateKeyFilename}"
    destination = "~/.ssh/id_rsa"
  }

  provisioner "file" {
    source      = "./scripts/openshiftConfig.sh"
    destination = "~/openshiftConfig.sh"
  }

  provisioner "remote-exec" {
    inline = [
      "chmod 600 ~/.ssh/id_rsa",
      "chmod 700 ~/openshiftConfig.sh",
      "~/openshiftConfig.sh ${module.openshift-cluster.master-private_dns} messages ${var.opUser} ${var.opPass} ${module.openshift-cluster.node1-private_dns} ${module.openshift-cluster.node2-private_dns} ${var.adminUsername}",
      "ANSIBLE_HOST_KEY_CHECKING=False /usr/local/bin/ansible-playbook -i ~/inventory.cfg ~/openShiftConfig.yml",
    ]
  }
}
