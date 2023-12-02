ARG DEVCONTAINER_BUILD_CERT_STATUS

FROM 658387908230.dkr.ecr.us-east-1.amazonaws.com/dkw-base-image:latest as base_no_cert
ONBUILD LABEL noop=noop

FROM 658387908230.dkr.ecr.us-east-1.amazonaws.com/dkw-base-image:latest as base_with_cert
ONBUILD USER root
ONBUILD COPY ./node_extra_cert.crt /usr/local/share/ca-certificates
ONBUILD RUN update-ca-certificates
ONBUILD ENV NODE_EXTRA_CA_CERTS=/usr/local/share/ca-certificates/node_extra_cert.crt

FROM base_${DEVCONTAINER_BUILD_CERT_STATUS}
USER vscode