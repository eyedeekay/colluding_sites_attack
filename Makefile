
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

run: network run-eepsite run-service run-website

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

run-service: network
	docker run -d --name fingerprint-service \
		--network fingerprint \
		--network-alias fingerprint-service \
		--hostname fingerprint-service \
		-p 8080:8080 \
		--restart always \
		eyedeekay/colluding_sites_attack_service

run-website: network
	docker run -d --name fingerprint-website \
		--network fingerprint \
		--network-alias fingerprint-website \
		--hostname fingerprint-website \
		-p 8080:8080 \
		--restart always \
		eyedeekay/colluding_sites_attack_website

list:
	./tunlist

test-classic:
	./test.sh | tee test.oldproxy.log

test-newhotness:
	./test.sh n | tee test.newproxy.log

diff:
	diff --width=210 --side-by-side --color=always test.oldproxy.log test.newproxy.log | tee test.diff

easysurf:
	http_proxy=http://127.0.0.1:4443 surf http://i2p-projekt.i2p

surf:
	http_proxy=http://127.0.0.1:4443 surf http://lqnwvwsgio6k53zq6d7r5bpaxuslc45vgsiqo6i3ebshkqpgrnma.b32.i2p
	http_proxy=http://127.0.0.1:4443 surf http://zcofypupen75rdv5zihviweyw5emk2l34idq423kbhj7n3owoe5a.b32.i2p
	http_proxy=http://127.0.0.1:4443 surf http://zjjjd756aucwz3pa2fl4mb3po2wtf752aefpod4gvedwreeox52q.b32.i2p
