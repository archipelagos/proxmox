# INFO: 
include mk/download.mk
include mk/dirs.mk
include mk/qemu.mk
include mk/disk.mk

# INFO: 
all:

# INFO: 
clean:
	rm \
		-rf \
		$(DIRS_IMG_DIR)

# INFO: 
run: \
		$(DOWNLOAD_ISO_FILE) \
		$(DISK_FILE)
	sudo $(QEMU_BIN) \
		$(QEMU_FLAGS)

# INFO: 
$(DOWNLOAD_ISO_FILE): \
		%: | \
		$(DIRS_ISO_DIR)
	curl \
		-o \
		$* \
		$(DOWNLOAD_URL)

# INFO: 
$(DISK_FILE): \
		$(DIRS_IMG_DIR)/%.$(QEMU_IMAGE_FORMAT): | \
		$(DIRS_IMG_DIR)
	$(QEMU_IMG_BIN) \
		create \
		-f \
		$(QEMU_IMAGE_FORMAT) \
		$@ \
		$(QEMU_IMAGE_SIZE)

# INFO:
$(DIRS_ISO_DIR): \
		%:
	mkdir \
		-p \
		$*

$(DIRS_IMG_DIR): \
		%:
	mkdir \
		-p \
		$*

