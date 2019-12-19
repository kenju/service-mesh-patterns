.PHONY: ping-foo
ping-foo:
	curl --verbose --header 'Host: foo.bar.com' localhost:80

.PHONY: ping-bar
ping-bar:
	curl --verbose --header 'Host: bar.foo.com' localhost:80

.PHONY: ping-root
ping-root:
	curl --verbose localhost:80
