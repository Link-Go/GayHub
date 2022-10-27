## alpine/dfimage
逆向工程通过一个镜像获取其 dockerfile

* hub: https://hub.docker.com/r/alpine/dfimage
* git: https://github.com/P3GLEG/Whaler


get it
```bash
    docker pull alpine/dfimage
```


run it
```bash
    $ alias dfimage="docker run -v /var/run/docker.sock:/var/run/docker.sock --rm alpine/dfimage"
    $ dfimage -sV=1.36 nginx:latest

    Analyzing nginx:latest
    Docker Version: 18.09.7
    GraphDriver: overlay2
    Environment Variables
    |PATH=/usr/local/sbin:/usr/local/bin:/usr/sbin:/usr/bin:/sbin:/bin
    |NGINX_VERSION=1.19.0
    |NJS_VERSION=0.4.1
    |PKG_RELEASE=1~buster

    Open Ports
    |80

    Image user
    |User is root

    Potential secrets:
    Dockerfile:
    CMD ["bash"]
    LABEL maintainer=NGINX Docker Maintainers <docker-maint@nginx.com>
    ENV NGINX_VERSION=1.19.0
    ENV NJS_VERSION=0.4.1
    ENV PKG_RELEASE=1~buster
    RUN set -x  \
        && addgroup --system --gid 101 nginx  \
        && adduser --system --disabled-login --ingroup nginx --no-create-home --home /nonexistent --gecos "nginx user" --shell /bin/false --uid 101 nginx  \
        && apt-get update  \
        && apt-get install --no-install-recommends --no-install-suggests -y gnupg1 ca-certificates  \
        && NGINX_GPGKEY=573BFD6B3D8FBC641079A6ABABF5BD827BD9BF62; found=''; for server in ha.pool.sks-keyservers.net hkp://keyserver.ubuntu.com:80 hkp://p80.pool.sks-keyservers.net:80 pgp.mit.edu ; do echo "Fetching GPG key $NGINX_GPGKEY from $server"; apt-key adv --keyserver "$server" --keyserver-options timeout=10 --recv-keys "$NGINX_GPGKEY"  \
        && found=yes  \
        && break; done; test -z "$found"  \
        && echo >&2 "error: failed to fetch GPG key $NGINX_GPGKEY"  \
        && exit 1; apt-get remove --purge --auto-remove -y gnupg1  \
        && rm -rf /var/lib/apt/lists/*  \
        && dpkgArch="$(dpkg --print-architecture)"  \
        && nginxPackages=" nginx=${NGINX_VERSION}-${PKG_RELEASE} nginx-module-xslt=${NGINX_VERSION}-${PKG_RELEASE} nginx-module-geoip=${NGINX_VERSION}-${PKG_RELEASE} nginx-module-image-filter=${NGINX_VERSION}-${PKG_RELEASE} nginx-module-njs=${NGINX_VERSION}.${NJS_VERSION}-${PKG_RELEASE} "  \
        && case "$dpkgArch" in amd64|i386) echo "deb https://nginx.org/packages/mainline/debian/ buster nginx" >> /etc/apt/sources.list.d/nginx.list  \
        && apt-get update ;; *) echo "deb-src https://nginx.org/packages/mainline/debian/ buster nginx" >> /etc/apt/sources.list.d/nginx.list  \
        && tempDir="$(mktemp -d)"  \
        && chmod 777 "$tempDir"  \
        && savedAptMark="$(apt-mark showmanual)"  \
        && apt-get update  \
        && apt-get build-dep -y $nginxPackages  \
        && ( cd "$tempDir"  \
        && DEB_BUILD_OPTIONS="nocheck parallel=$(nproc)" apt-get source --compile $nginxPackages )  \
        && apt-mark showmanual | xargs apt-mark auto > /dev/null  \
        && { [ -z "$savedAptMark" ] || apt-mark manual $savedAptMark; }  \
        && ls -lAFh "$tempDir"  \
        && ( cd "$tempDir"  \
        && dpkg-scanpackages . > Packages )  \
        && grep '^Package: ' "$tempDir/Packages"  \
        && echo "deb [ trusted=yes ] file://$tempDir ./" > /etc/apt/sources.list.d/temp.list  \
        && apt-get -o Acquire::GzipIndexes=false update ;; esac  \
        && apt-get install --no-install-recommends --no-install-suggests -y $nginxPackages gettext-base curl  \
        && apt-get remove --purge --auto-remove -y  \
        && rm -rf /var/lib/apt/lists/* /etc/apt/sources.list.d/nginx.list  \
        && if [ -n "$tempDir" ]; then apt-get purge -y --auto-remove  \
        && rm -rf "$tempDir" /etc/apt/sources.list.d/temp.list; fi  \
        && ln -sf /dev/stdout /var/log/nginx/access.log  \
        && ln -sf /dev/stderr /var/log/nginx/error.log  \
        && mkdir /docker-entrypoint.d
    COPY file:d68fadb480cbc781c3424ce3e42e1b5be80133bdcce2569655e90411a4045da2 in /
        docker-entrypoint.sh

    COPY file:b96f664d94ca7bbe69241468d85ee421e9d310ffa36f3b04c762dcce9a42c7f1 in /docker-entrypoint.d
        docker-entrypoint.d/
        docker-entrypoint.d/10-listen-on-ipv6-by-default.sh

    COPY file:cc7d4f1d03426ebd11e960d6a487961e0540059dcfad14b33762f008eed03788 in /docker-entrypoint.d
        docker-entrypoint.d/
        docker-entrypoint.d/20-envsubst-on-templates.sh

    ENTRYPOINT ["/docker-entrypoint.sh"]
    EXPOSE 80
    STOPSIGNAL SIGTERM
    CMD ["nginx" "-g" "daemon off;"]
```