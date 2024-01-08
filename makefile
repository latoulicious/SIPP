.PHONY: frontend backend

frontend:
	cd web && yarn dev

backend:
	cd server && air

run:
wt -d web yarn dev ; wt -d server air