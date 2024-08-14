#для докера
d_connect:
#как обычный флаг d_connect container=1234
	docker exec -it ${container} /bin/bash

d_run:
	docker run -p 8080:8080 end1essrage/whats-distrib-backend

d_build: 
	docker build -t end1essrage/whats-distrib-backend .

#для подмена
p_connect:
#как обычный флаг p_connect container=1234
	podman exec -it ${container} /bin/bash
	
p_run:
	podman run -p 8080:8080 end1essrage/whats-distrib-backend

p_build: 
	podman build -t end1essrage/whats-distrib-backend .