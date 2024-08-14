#defaults
port ?= 8080
env = dev

#для докера
d_connect:
#как обычный флаг d_connect container=1234
	docker exec -it ${container} /bin/bash

d_run:
	docker ps -a --filter "ancestor=end1essrage/whats-distrib-backend" -q | xargs docker stop
	docker ps -a --filter "ancestor=end1essrage/whats-distrib-backend" -q | xargs docker rm

	docker run --rm -p ${port}:8080 -e ENV=$(env) end1essrage/whats-distrib-backend

d_build: 
	docker build -t end1essrage/whats-distrib-backend .

#для подмена
p_connect:
#как обычный флаг p_connect container=1234
	podman exec -it ${container} /bin/bash
	
p_run:
	podman ps -a --filter "ancestor=end1essrage/whats-distrib-backend" -q | xargs podman stop
	podman ps -a --filter "ancestor=end1essrage/whats-distrib-backend" -q | xargs podman rm

	podman run -p ${port}:8080 -e ENV=$(env) end1essrage/whats-distrib-backend

p_build: 
	podman build -t end1essrage/whats-distrib-backend .