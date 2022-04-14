FROM registry.access.redhat.com/ubi8/ubi

MAINTAINER Red Hat Nvidia Team

# Install dependencies
RUN yum install -y \
		glibc-langpack-en \
		go git make jq vim wget rsync time && \
	yum clean all && \
	rm -rf $HOME/.cache && \
	rm -rf /var/cache/yum


# Install dependencies: `oc`
ARG OCP_CLI_VERSION=latest
ARG OCP_CLI_URL=https://mirror.openshift.com/pub/openshift-v4/x86_64/clients/ocp/${OCP_CLI_VERSION}/openshift-client-linux.tar.gz
RUN curl ${OCP_CLI_URL} | tar xfz - -C /usr/local/bin oc

# Install dependencies: `operator-sdk`
ARG OPERATOR_SDK_VERSION=v1.6.2
ARG OPERATOR_SDK_URL=https://github.com/operator-framework/operator-sdk/releases/download/${OPERATOR_SDK_VERSION}
RUN cd /usr/local/bin \
 && curl -LO ${OPERATOR_SDK_URL}/operator-sdk_linux_amd64 \
 && mv operator-sdk_linux_amd64 operator-sdk \
 && chmod +x operator-sdk

# Get the source code in there
RUN git clone https://github.com/rh-ecosystem-edge/nvidia-ci.git /root/nvidia-ci
WORKDIR /root/nvidia-ci

ENTRYPOINT ["bash"]
