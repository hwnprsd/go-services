.PHONY: 

AWS_ECR_URL = 312136753954.dkr.ecr.ap-south-1.amazonaws.com
AWS_REGION = ap-south-1

build:
	$(MAKE) -C ./api modsync
	aws ecr get-login-password --region ${AWS_REGION} | docker login --username AWS --password-stdin ${AWS_ECR_URL}
	docker context use default
	docker compose build
	docker compose push
	echo "DONE"

deploy:
	docker context use opl 
	docker compose up
	echo "DONE"

run:
	$(MAKE) -C ./api modsync
	docker-compose -f docker-compose.dev.yml --env-file .dev.env up --build --remove-orphans --force-recreate
	echo "DONE"
	

