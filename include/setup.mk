
network:
	docker network create fingerprint
	@echo 'fingerprint' | tee network

log-network:
	docker network inspect fingerprint

clean-network:
	rm -f network
	docker network rm fingerprint; true

redo-network: clean-network network

build: build-eepsite build-service build-website

build-eepsite:
	docker build -f Dockerfiles/Dockerfile.eepSite -t eyedeekay/colluding_sites_attack_eepsite .

build-service:
	docker build -f Dockerfiles/Dockerfile.service -t eyedeekay/colluding_sites_attack_service .

build-website:
	docker build -f Dockerfiles/Dockerfile.website -t eyedeekay/colluding_sites_attack_website .

clean: clean-eepsite clean-service clean-website

clean-eepsite:
	docker rm -f fingerprint-eepsite; true

clean-service:
	docker rm -f fingerprint-service; true

clean-website:
	docker rm -f fingerprint-website; true

clobber: clobber-eepsite clobber-service clobber-website

clobber-eepsite:
	docker rmi -f eyedeekay/colluding_sites_attack_eepsite; true

clobber-service:
	docker rmi -f eyedeekay/colluding_sites_attack_service; true

clobber-website:
	docker rmi -f eyedeekay/colluding_sites_attack_website; true

log-eepsite:
	docker logs fingerprint-eepsite

log-service:
	docker logs fingerprint-service

log-website:
	docker logs fingerprint-website

update: clean build run

update-service: clean-service build-service run-service

update-eepsite: clean-eepsite build-eepsite run-eepsite

update-eepsite: clean-website build-website run-website
