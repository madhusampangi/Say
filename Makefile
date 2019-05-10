# build:
# 	GOOS=linux go build -o app
# 	docker build -t gcr.io/madhu-text2speech/say .
# 	rm -f app

# push:
# 	gcloud docker -- push gcr.io/madhu-text2speech/say

build:
	docker build -t madhusampangi/say .