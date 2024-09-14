include leetty-gateway.make.env

start: compile
	./bin/leetty-gateway

compile: clean
	mkdir bin && cd bin && go build ../cmd/leetty-gateway/leetty-gateway.go

clean:
	rm -rf ./bin

image:
	docker run \
	-v /var/run/docker.sock:/var/run/docker.sock \
 	-v ${PWD}:/workspace \
 	-w /workspace \
	${LEETTY_GATEWAY_IMAGE_PACK} build \
	leetty-gateway:${LEETTY_GATEWAY_VERSION} \
	--builder ${LEETTY_GATEWAY_IMAGE_BUILDER} \
	--buildpack ${LEETTY_GATEWAY_IMAGE_BUILDPACK} \
	--run-image ${LEETTY_GATEWAY_PACK_RUN_IMAGE} \
	--env BP_GO_VERSION=${LEETTY_GATEWAY_GO_VERSION} \
	--env BP_GO_TARGETS=./cmd/leetty-gateway