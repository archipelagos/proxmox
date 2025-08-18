# INFO: Disks configuration file.
ifndef MK_DISK_MK
MK_DISK_MK=					true
#--------------------------------------------------------------------------------------------------

# INFO: Includes from other build script modules.
include mk/qemu.mk
include mk/dirs.mk

# INFO: Disk file name.
DISK_FILE_NAME=				disk.$(QEMU_IMAGE_FORMAT)

# INFO: Disk file path.
DISK_FILE=					$(DIRS_IMG_DIR)/$(DISK_FILE_NAME)

#--------------------------------------------------------------------------------------------------
endif
