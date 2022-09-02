ARG base_image=golang:1.18.5-alpine3.16
FROM $base_image as base_image

SHELL ["/bin/sh", "-euxo", "pipefail", "-c"]

RUN \
--mount=type=cache,target=/var/cache/apk \
apk add --no-cache \
 bash \
 curl \
 git \
 jq \
 make \
 nodejs \
 terraform \
 yarn \
; \
:;

SHELL ["/bin/bash", "-euxo", "pipefail", "-c"]

COPY ./provider/go.* /app/provider/

WORKDIR /app

RUN \
curl -fsSL https://get.pulumi.com -o /tmp/pulumi-install.sh; \
chmod +x /tmp/pulumi-install.sh; \
cd /app/provider; \
/tmp/pulumi-install.sh \
  --version \
  "$(go list -m -u -json "github.com/pulumi/pulumi/sdk/v3" | jq --raw-output '.Version[1:]')" \
; \
:;

ENV PATH="/root/.pulumi/bin:${PATH}"

COPY ./assets/root/.terraformrc /root/.terraformrc

COPY ./ /app/

RUN \
cd /app/tftest; \
terraform init; \
:;

ENV PULUMICTL_VERSION="0.0.36"
RUN \
git clone --depth=1 --branch="v${PULUMICTL_VERSION}" https://github.com/pulumi/pulumictl/ /tmp/pulumictl; \
cd /tmp/pulumictl; \
make; \
make install; \
:;

RUN \
cd /app/provider; \
go mod tidy; \
go build; \
cd /app; \
#Disable CGO to get a static exe.  We end up with trouble on other platforms that don't use musl otherwise.
CGO_ENABLED=0 make build; \
:;

COPY ./bundle /root/bundle
RUN \
cp /app/bin/pulumi-resource-megaport /root/bundle; \
:;

RUN \
if ldd /root/bundle/pulumi-resource-megaport; then false; fi; \
:;

ARG PACKAGE_VERSION="0.0.9"
RUN tar czvf "/root/pulumi-resource-megaport-v${PACKAGE_VERSION}-linux-amd64.tar.gz" --directory=/root/bundle .
RUN cp -a /app/sdk/nodejs/scripts /app/sdk/nodejs/bin/scripts

FROM scratch as pulumi-megaport-npm
ARG PACKAGE_VERSION="0.0.9"

COPY --from=base_image /root/.terraformrc /pulumi-megaport/.terraformrc
COPY --from=base_image /root/pulumi-resource-megaport-v${PACKAGE_VERSION}-linux-amd64.tar.gz /pulumi-megaport/
COPY --from=base_image /app/sdk/nodejs /pulumi-megaport/nodejs
