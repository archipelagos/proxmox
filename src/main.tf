terraform {
  required_version = ">= 1.12.2"
  required_providers {
    proxmox = {
      source = "telmate/proxmox"
      version = "2.7.4"
    }
  }
}

provider "proxmox" {
  pm_api_url          = "${var.api_url}"
  pm_api_token_id     = "${var.api_token_id}"
  pm_api_token_secret = "${var.api_token_secret}"
  pm_tls_insecure = true
}

resource "proxmox_vm_qemu" "my_vm" {
  count       = 1
  name        = "my-vm"
  target_node = "miner"
#  preprovision = true
  os_type      = "ubuntu"
  iso         = "home/Terraform/foo.iso"
#  clone       = "ubuntusrvgp"
  agent       = 1
  sockets     = 1
  cores       = 2
  memory      = 2048
  scsihw      = "virtio-scsi-pci"
  bootdisk    = "scsi0"

  disk {
    slot      = 0
    size      = "10G"
    type      = "scsi"
    storage   = "local-lvm"
    iothread  = 1
  }

  network {
    model     = "virtio"
    bridge    = "vmbr0"
  }

  lifecycle {
    ignore_changes = [
      network,
    ]
  }
}