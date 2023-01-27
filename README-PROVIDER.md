# Pulumi Megaport Resource Provider

The pulumi-megaport Resource Provider lets you manage Megaport resources.

## Building

1. docker buildx build --tag pulumi-megaport:<tag> .
2. CONTAINER_ID=$(docker create --name pulumi-megaport-extract pulumi-megaport:<tag> fake-command) (<-- The command is arbitrary - we are creating but not actually running a container - just extract the data from it)
3. docker export --out /tmp/out.tar ${CONTAINER_ID}
4. cd /tmp; # or whereever
5. tar xvf /tmp/out.tar pulumi-megaport/
6. (Optional) docker rm ${CONTAINER_ID}
7. pulumi plugin install resource pulumi-megaport 0.0.10 --file ./tmp/pulumi-megaport/pulumi-resource-megaport-v0.0.10-linux-amd64.tar.gz

## Installing

This package is (TODO: will be) available for several languages/platforms:

### Node.js (JavaScript/TypeScript)

To use from JavaScript or TypeScript in Node.js, install using either `npm`:

#### From packaged release

```bash
npm install @pulumi/foo
```

or `yarn`:

```bash
yarn add @pulumi/foo
```

#### In local development environment

```bash
yarn link -r <path>/to/extracted/pulumi-megaport/dist/nodejs
```

Note: 
1. See building steps above for extracting the nodejs.
2. you may need to manually add `"@pulumi/megaport": "*",` to dependencies in package.json.
3. yarn link adds a relative path to your local dev environment. Do not commit to git.

# TODO - complete the remaining readme boilerplate


### Python

To use from Python, install using `pip`:

```bash
pip install pulumi_foo
```

### Go

To use from Go, use `go get` to grab the latest version of the library:

```bash
go get github.com/pulumi/pulumi-foo/sdk/go/...
```

### .NET

To use from .NET, install using `dotnet add package`:

```bash
dotnet add package Pulumi.Foo
```

## Configuration

The following configuration points are available for the `foo` provider:

- `foo:apiKey` (environment: `FOO_API_KEY`) - the API key for `foo`
- `foo:region` (environment: `FOO_REGION`) - the region in which to deploy resources

## Reference

For detailed reference documentation, please visit [the Pulumi registry](https://www.pulumi.com/registry/packages/foo/api-docs/).
