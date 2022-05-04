#ARG base_image=statelesstestregistry.azurecr.io/stateless/base:24.0
ARG base_image=golang:1.18.1-alpine3.15
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

RUN \
git clone --depth=1 https://github.com/BeStateless/terraform-provider-megaport /root/terraform-provider-megaport; \
cd /root/terraform-provider-megaport; \
go mod tidy; \
make; \
go build; \
:;

COPY ./assets/root/.terraformrc /root/.terraformrc

COPY ./ /app/

RUN \
cd /app/tftest; \
terraform init; \
:;

RUN \
wget \
  --output-document=- \
  https://github.com/pulumi/pulumictl/releases/download/v0.0.31/pulumictl-v0.0.31-linux-amd64.tar.gz | \
    tar --extract --gzip --file - --directory=/usr/local/bin \
  ; \
:;

RUN \
cd /app; \
make build; \
:;

#make build
