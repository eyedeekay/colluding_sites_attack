
network:
	docker network create fingerprint; true

log-network:
	docker network inspect fingerprint

clean-network:
	rm -f network
	docker network rm fingerprint; true

redo-network: clean-network network

build: build-eepsite build-service build-website

build-eepsite:
	docker build --force-rm -f Dockerfiles/Dockerfile.eepSite -t eyedeekay/colluding_sites_attack_eepsite .

build-service:
	docker build --force-rm -f Dockerfiles/Dockerfile.service -t eyedeekay/colluding_sites_attack_service .

build-website:
	docker build --force-rm -f Dockerfiles/Dockerfile.website -t eyedeekay/colluding_sites_attack_website .

clean: clean-eepsite clean-service clean-website

clean-eepsite:
	docker rm -f sam-host; true

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
	docker logs sam-host

log-service:
	docker logs fingerprint-service

log-website:
	docker logs fingerprint-website

update: clean build run

update-service: clean-service build-service run-service

update-eepsite: clean-eepsite build-eepsite run-eepsite

update-website: clean-website build-website run-website
