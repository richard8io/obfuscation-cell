FROM phusion/baseimage:0.9.18
MAINTAINER Richard Hunt <richard.io@protonmail.com>
LABEL app=obfuscation-cell

# Set correct environment variables.
ENV HOME /root

RUN apt-get update && DEBIAN_FRONTEND=noninteractive apt-get upgrade -y -o Dpkg::Options::="--force-confold"

# Clean up APT when done.
RUN apt-get clean && rm -rf /var/lib/apt/lists/* /tmp/* /var/tmp/*

ADD obfuscation-cell /home/app/obfuscation-cell
ADD sample.pdf /home/app/sample.pdf

# Use baseimage-docker's init system.
ADD cell_init.sh /etc/service/obfuscation-cell/run
CMD ["/sbin/my_init"]