# ---------------------- continerd build -------------------
FROM golang:1.12 as containerd

RUN apt update && \
	apt install -y \
	rsync \
	autoconf \
	automake \
	g++ \
	libtool \
	unzip \
	btrfs-tools \
	gcc \
	git \
	libseccomp-dev \
	make \
	xfsprogs

ENV CONTAINERD_VERSION 3d577f172af7798e85a51d8706a6b043414aeaf6
RUN git clone https://github.com/crosbymichael/containerd /go/src/github.com/containerd/containerd
RUN git clone https://github.com/opencontainers/runc /go/src/github.com/opencontainers/runc

WORKDIR /go/src/github.com/containerd/containerd
RUN git checkout ${CONTAINERD_VERSION}

RUN rsync -au /go/src/github.com/containerd/containerd/vendor/ /go/src/ && \
	rm -rf /go/src/github.com/containerd/containerd/vendor/

RUN ./script/setup/install-protobuf
RUN ./script/setup/install-runc
RUN make

# ------------------ orbit build -------------------------
FROM containerd as orbit

RUN git clone https://github.com/stellarproject/orbit /go/src/github.com/stellarproject/orbit

RUN	rm -rf /go/src/github.com/stellarproject/orbit/vendor/github.com/containerd/ && \
	rsync -au --ignore-existing /go/src/github.com/stellarproject/orbit/vendor/ /go/src/ && \
	rm -rf /go/src/github.com/stellarproject/orbit/vendor/

WORKDIR /go/src/github.com/stellarproject/orbit

RUN make && make plugin

# ------------------- systemd build -------------------------
FROM ubuntu:18.10 as systemd

RUN apt update && \
        apt upgrade -y

RUN apt install -y \
	build-essential \
	ca-certificates \
	bison \
	git \
	meson \
	gperf \
	make \
	libssl-dev \
	libcap-dev \
	pkg-config \
	mount \
	libmount-dev \
	gettext \
	libapparmor-dev \
	libpam-dev \
	libaudit-dev \
	libcryptsetup-dev \
	libgcrypt-dev \
	libkmod-dev \
	libacl1-dev \
	libseccomp-dev \
	liblz-dev \
	libidn2-dev \
	libcurl4-openssl-dev

ENV SYSTEMD_VERSION v240
RUN git clone https://github.com/systemd/systemd /systemd

WORKDIR /systemd
RUN git checkout ${SYSTEMD_VERSION}

RUN mkdir /install && \
	./configure && \
	make && \
	make DESTDIR=/install install

# -------------------- production image ------------------------
FROM ubuntu:18.10

RUN apt update && \
	apt upgrade -y

RUN	apt install -y \
	openssh-server \
	iproute2 \
	kmod

# systemd install
COPY --from=systemd /install /

# containerd binaries
COPY --from=containerd /go/src/github.com/containerd/containerd/containerd.service /lib/systemd/system/
COPY --from=containerd /go/src/github.com/containerd/containerd/bin/* /bin/
COPY --from=containerd /usr/local/sbin/runc /sbin/

# orbit binaries
RUN mkdir -p /var/lib/containerd/plugins
COPY --from=orbit /go/src/github.com/stellarproject/orbit/bin/* /usr/local/bin/
COPY --from=orbit /go/src/github.com/stellarproject/orbit/plugins/* /var/lib/containerd/plugins/

RUN systemctl enable ssh containerd
RUN systemctl disable systemd-resolved

CMD ["/sbin/init"]
