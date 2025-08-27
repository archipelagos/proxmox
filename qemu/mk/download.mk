# INFO: Downloads configuration file.
ifndef MK_DOWNLOAD_MK
MK_ARCHISO_MK=				true
#--------------------------------------------------------------------------------------------------

# INFO: Includes from other build script modules.
include mk/dirs.mk

# INFO: Extension of downloaded image file.
DOWNLOAD_ISO_EXT=			iso

# INFO: File name of downloaded image file.
DOWNLOAD_ISO_FILE_NAME=		proxmox-ve_8.4-1.$(DOWNLOAD_ISO_EXT)

# INFO: Path of downloaded image file.
DOWNLOAD_ISO_FILE=			$(DIRS_ISO_DIR)/$(DOWNLOAD_ISO_FILE_NAME)

# INFO: URL where ISO image is loacted.
DOWNLOAD_URL=				https://enterprise.proxmox.com/iso/$(DOWNLOAD_ISO_FILE_NAME)

#--------------------------------------------------------------------------------------------------
endif
