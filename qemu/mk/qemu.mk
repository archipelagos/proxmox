# INFO: QEMU configuration file.
ifndef MK_QEMU_MK
MK_QEMU_MK=					true
#--------------------------------------------------------------------------------------------------

# INFO: Includes from other build script modules.
include mk/config.mk
include mk/download.mk
include mk/disk.mk

# INFO: General architecture configuration and binary associated to it.
QEMU_BIN=					qemu-system-x86_64
QEMU_IMG_BIN=				qemu-img

# INFO: Disk image configuration.
QEMU_IMAGE_FORMAT=			qcow2
QEMU_IMAGE_SIZE=			32G

# INFO: Default flags.
QEMU_FLAGS=

# INFO: Ports related flags.
QEMU_FLAGS+=				-serial \
							none \
							-parallel \
							none

# INFO: Name of the window displayed in Window Manager (title bar).
QEMU_FLAGS+=				-name \
							$(CONFIG_NAME)

# INFO: Time configuration.
QEMU_FLAGS+=				-rtc \
							clock=host,base=localtime

# INFO: Basic CPU settings.
#QEMU_FLAGS+=				-cpu \
#							host,kvm=off,hv_relaxed,hv_spinlocks=0x1fff,hv_vapic,hv_time,hv_vendor_id=Nvidia43FIX

# INFO: Remove hv_time to see if anything happens
#QEMU_FLAGS+=				-cpu \
#							host,kvm=off,hv_relaxed,hv_spinlocks=0x1fff,hv_vapic,hv_vendor_id=Nvidia43FIX

# INFO: CPU configuration.
QEMU_FLAGS+=				-cpu \
							host,kvm=on

# INFO: Symmetric multiprocessing configuration.
QEMU_FLAGS+=				-smp \
							8,sockets=1,cores=4,threads=2

# INFO: Enable KVM full virtualization support.
QEMU_FLAGS+=				-enable-kvm

# INFO: Assign memory to the VM. Hugepages requires additional configuration.
QEMU_FLAGS+=				-m \
							4G

#QEMU_FLAGS+=				-mem-path \
#							/dev/hugepages
#QEMU_FLAGS+=				-mem-prealloc

# INFO: VFIO GPU and GPU sound passthrough.
#QEMU_FLAGS+=				-device \
#							vfio-pci,host=02:00.0,multifunction=on

#QEMU_FLAGS+=				-device \
#							vfio-pci,host=02:00.1

# INFO: Supply OVMF (general UEFI bios, needed for EFI boot support with GPT disks).
QEMU_FLAGS+=				-drive \
							if=pflash,format=raw,readonly=on,file=/usr/share/OVMF/OVMF_CODE_4M.fd

QEMU_FLAGS+=				-drive \
							if=pflash,format=raw,readonly=on,file=/usr/share/OVMF/OVMF_VARS_4M.fd

# INFO: Load our created VM image as a harddrive.
#
# NOTE: Giving the VM *raw* disk access can lead to unintended things
# happening. Do this only if you're giving the VM dedicated and
# exclusive access to the HDD.
#
QEMU_FLAGS+=				-device \
							virtio-scsi-pci,id=scsi

QEMU_FLAGS+=				-drive \
							file=$(DISK_FILE),cache=none,if=virtio,format=qcow2

# INFO: Load virtio drivers
QEMU_FLAGS+=				-cdrom \
							$(DOWNLOAD_ISO_FILE)

# INFO: Use the following emulated video device (use none for disabled).
QEMU_FLAGS+=				-vga \
							std

# INFO: Running from the shell
#QEMU_FLAGS+=				-nographic

# INFO: Redirect QEMU's console input and output.
#QEMU_FLAGS+=				-monitor \
#							stdio

# INFO: Passthrough USB devices.
#QEMU_FLAGS+=				-usb

# INFO: USB mouse
#QEMU_FLAGS+=				-usbdevice \
#							host:dead:beef

# INFO: USB keyboard
#QEMU_FLAGS+=				-usbdevice \
#							host:dead:beef

# INFO: Network configuration
# 
# See qemu documentation for more details, but '-net user' will put the
# VM in its own subnet that only the host can access. The Windows VM
# will be able to initiate communications on the LAN and WAN.
#
# Passing the 'smb=/path/' option will start QEMU's Samba server on the
# host that the guest can access. It does have read/write access to the
# shared folder (depending on the permissions or mount options you've
# set).  If you need a more complex configuration, then you'll probably
# want to set up your own Samba server.
#
QEMU_FLAGS+=				-net \
							nic,model=virtio,macaddr=e6:c8:ff:09:76:99 \
							-net \
							bridge,br=br0

#QEMU_FLAGS+=				-net \
#							nic \
#							-net \
#							user,smb=/shared/

# first two options are related to audio configuration
# 'taskset' pins qemu to certain CPUs for more consistent performance
# (otherwise qemu seems switch to whichever cores it feels like).
#QEMU_AUDIO_DRV=pa QEMU_PA_SAMPLES=128 taskset -c 21-31 qemu-system-x86_64 $QEMU_FLAGS

#--------------------------------------------------------------------------------------------------
endif
