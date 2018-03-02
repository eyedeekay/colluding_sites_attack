
network:
	docker network create fingerprint
	@echo 'fingerprint' | tee network

log-network:
	docker network inspect fingerprint

clean-network: clean
	rm -f network
	docker network rm fingerprint; true

build: build-eepsite build-service

build-eepsite:
	docker build -f Dockerfiles/Dockerfile.eepSite -t eyedeekay/colluding_sites_attack_eepsite .

build-service:
	docker build -f Dockerfiles/Dockerfile.service -t eyedeekay/colluding_sites_attack_service .

clean: clean-eepsite clean-service

clean-eepsite:
	docker rm -f fingerprint-eepsite; true

clean-service:
	docker rm -f fingerprint-service; true

clobber: clobber-eepsite clobber-service

clobber-eepsite:
	docker rmi -f eyedeekay/colluding_sites_attack_eepsite; true

clobber-service:
	docker rmi -f eyedeekay/colluding_sites_attack_service; true

log-eepsite:
	docker logs fingerprint-eepsite

log-service:
	docker logs fingerprint-service
