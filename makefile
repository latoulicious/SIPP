.PHONY: server web

server:
	$(MAKE) -C server go

web:
	$(MAKE) -C web vue

run: server web