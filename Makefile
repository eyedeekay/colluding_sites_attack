
GOPATH=$(PWD)/.go

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

include include/config.mk
include include/setup.mk

run: network run-eepsite run-service run-website

run-eepsite: network
	docker run -d --name sam-host \
		--network si \
		--network-alias sam-host \
		--hostname sam-host \
		--expose 4567 \
		--link fingerprint-service \
		--link fingerprint-website \
		-p :4567 \
		-p 127.0.0.1:7072:7072 \
		--volume $(i2pd_dat):/var/lib/i2pd:rw \
		--restart always \
		eyedeekay/colluding_sites_attack_eepsite

run-service: network
	docker run -i -t \
		--network si \
		--restart always \
		-v /home/reflect/ \
		eyedeekay/colluding_sites_attack_service

run-website: network
	docker run -d --name fingerprint-website \
		--network si \
		--network-alias fingerprint-website \
		--hostname fingerprint-website \
		--restart always \
		-p 127.0.0.1:8081:8081 \
		eyedeekay/colluding_sites_attack_website

list:
	./scripts/tunlist | tee artifacts/tunlist.log

test-classic:
	./scripts/test.sh | tee artifacts/test.oldproxy.log

test-newhotness:
	./scripts/test.sh n | tee artifacts/test.newproxy.log

diff:
	diff --width=210 --side-by-side --color=always artifacts/test.oldproxy.log artifacts/test.newproxy.log | tee artifacts/test.diff

deps:
	go get -u github.com/eyedeekay/sam-forwarder

compile:
	go build http-headers.go

del:
	rm http-headers
