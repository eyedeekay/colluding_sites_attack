
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
include attack

run: network run-service run-website

run-volume: network
	docker run -i -t -d \
		-e TAG=$(attacker) \
		--network si \
		--name "collude-volume" \
		-v reflect-volume:/home/reflect/ \
		eyedeekay/colluding_sites_attack_service; true
	docker run -i -t -d \
		--network si \
		--name "site-volume" \
		-v static-volume:/opt/eephttpd/ \
		eyedeekay/colluding_sites_attack_service; true

run-service: network run-volume
	docker run -i -t \
		-d \
		-e TAG=$(attacker) \
		--name "collude-$(attacker)" \
		--network si \
		--restart always \
		--volumes-from collude-volume \
		eyedeekay/colluding_sites_attack_service
	sleep 5
	@echo -n "  * " | tee -a colluders.md
	docker logs "collude-$(attacker)" | grep "b32.i2p" | head -n 1 | tee -a colluders.md

run-website: network
	docker run -d --name fingerprint-website \
		--network si \
		--network-alias fingerprint-website \
		--hostname fingerprint-website \
		--restart always \
		-p 127.0.0.1:8081:8081 \
		--volumes-from site-volume \
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

codemd:
	@echo -n "        " > temp.md
	cat http-headers.go >> temp.md
	awk 1 ORS='\n        ' temp.md > code.md
	rm -r temp.md

index: codemd
	cat include/index.top.html > index.html
	markdown colluders.md >> index.html
	cat include/index.middle.html >> index.html
	markdown code.md >> index.html
	cat include/index.bottom.html >> index.html

finger:
	wget -qO include/fingerprint2.js https://github.com/Valve/fingerprintjs2/raw/master/fingerprint2.js

readme:
	head -n $(SAVE_README_LINES) README.md > TEMPREADME.md
	@echo "" >> TEMPREADME.md
	cat TEMPREADME.md colluders.md > README.md
	rm -f TEMPREADME.md
