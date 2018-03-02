
usage:
	@echo 'usage:'
	@echo '======'
	@echo
	@echo ' install: make install'
	@echo ' reinstall without purging some settings: make reinstall'
	@echo ' re-generate all settings: make clobber install'
	@echo
	@echo 'Configured Directories'
	@echo '----------------------'
	@echo
	@echo "  * Host I2PD Data Directory $(i2pd_dat)"
	@echo
	@echo "Page Configuration"
	@echo "------------------"
	@echo
	@echo "$(CONFIG_PAGE)"
	@echo

include config.mk
include include/setup.mk

run: network run-eepsite run-website

run-eepsite: network
	docker run -d --name fingerprint-eepsite \
		--network fingerprint \
		--network-alias fingerprint-eepsite \
		--hostname fingerprint-eepsite \
		--expose 4567 \
		--link fingerprint-service \
		-p :4567 \
		-p 127.0.0.1:7072:7072 \
		--volume $(i2pd_dat):/var/lib/i2pd:rw \
		--restart always \
		eyedeekay/colluding_sites_attack_eepsite

run-website: network
	docker run -d --name fingerprint-service \
		--network fingerprint \
		--network-alias fingerprint-service \
		--hostname fingerprint-service \
		--restart always \
		eyedeekay/colluding_sites_attack_service

list:
	./tunlist
