
network:
	docker network create si; true

log-network:
	docker network inspect si

clean-network:
	docker network rm si; true

docker-build: build-service build-website

build-service:
	docker build --force-rm -f Dockerfiles/Dockerfile.service -t eyedeekay/colluding_sites_attack_service .

build-website:
	docker build --force-rm -f Dockerfiles/Dockerfile.website -t eyedeekay/colluding_sites_attack_website .

docker-clean: clean-service clean-website

clean-service:
	docker rm -f fingerprint-service; true

clean-website:
	docker rm -f fingerprint-website; true

docker-clobber: clobber-service clobber-website

clobber-service:
	docker rmi -f eyedeekay/colluding_sites_attack_service; true

clobber-website:
	docker rmi -f eyedeekay/colluding_sites_attack_website; true

log-service:
	docker logs fingerprint-service

log-website:
	docker logs fingerprint-website

update: docker-clean docker-build run

update-service: clean-service build-service run-service

update-website: build-website run-website
