
GOPATH=$(PWD)/.go

usage:
	@echo 'usage:'
	@echo '======'
	@echo
	@echo "$(attacker)"
	@echo
	@echo ' install: make update'
	@echo ' reinstall without purging some settings: make update'
	@echo ' re-generate all settings: make docker-clobber update'
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

run: network run-service run-website

run-volume: network
	docker run -i -t -d \
		--name "collude-volume" \
		--restart always \
		-v reflect-volume:/home/reflect/ \
		eyedeekay/colluding_sites_attack_service; true

run-service: network run-volume
	docker run -i -t \
		-d \
		--name "collude-$(attacker)" \
		--network si \
		--restart always \
		--volumes-from collude-volume \
		eyedeekay/colluding_sites_attack_service
	docker logs "collude-$(attacker)" | tee -a colluders.md

run-website: network
	docker run -d --name fingerprint-website \
		--network si \
		--network-alias fingerprint-website \
		--hostname fingerprint-website \
		--restart always \
		-p 127.0.0.1:8081:8081 \
		-v fingerprint-website:/home/eephttpd \
		eyedeekay/colluding_sites_attack_website

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

