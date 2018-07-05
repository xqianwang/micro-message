resource "aws_instance" "openShiftConfig" {
  depends_on = ["null_resource.openShiftDeploy"]

  connection {
    type        = "ssh"
    user        = "${var.adminUsername}"
    host        = "${aws_instance.bastion.ip_address.public_dns}"
    private_key = "${file("${var.privateKeyFilename}")}"
  }

  provisioner "file" {
    source      = "../scripts/openshiftConfig.sh"
    destination = "~/openshiftConfig.sh"
  }

  provisioner "remote-exec" {
    inline = [
      "chmod 700 ~/openshiftConfig.sh",
      "~/openshiftConfig.sh ${aws_instance.master.private_dns} messages ${var.opUser} ${var.opPass} ${aws_instance.node1.private_dns} ${aws_instance.node2.private_dns} ${var.adminUsername}",
      "ansible-playbook ~/openShiftConfig.yml",
    ]
  }
}
