# ytcli - Small script to search yt
# See LICENSE file for copyright and license details.

PREFIX ?= /usr/local

# Clean based on PREFIX
install:
	@mkdir -p ${DESTDIR}${PREFIX}/bin
	@cp -f ytcli ${DESTDIR}${PREFIX}/bin/ytcli
	@chmod 755 ${DESTDIR}${PREFIX}/bin/ytcli

.PHONY: install
