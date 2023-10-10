# dumpSoftwareVersions

Tiny simple command line utility used to create a software versions YAML file from the various tools used throughout a [Nextflow](https://www.nextflow.io/) pipeline, in order to generate the Software Versions HTML table input used with [MultiQC](https://multiqc.info/) reports.

# Usage

Example usage:

```bash
$ ./dumpSoftwareVersions -manifestName dump-software-version-demo -manifestVersion 1.0 -nxfVersion 23.04.1 -processLabel CUSTOM_DUMPSOFTWAREVERSIONS example/collated_versions.yml
```

Output YAML is printed to stdout by default. You will want to pipe it to a file named `software_versions_mqc.yml` for use with MultiQC.

- example input file: `example/collated_versions.yml`
- example output: `example/software_versions_mqc.yml`

# Download

Get it from Docker Hub here: https://hub.docker.com/repository/docker/stevekm/dump-software-versions/general

```bash
docker pull stevekm/dump-software-versions:0.1
```

- try it out with the included Makefile recipe; `make docker-test-run`

Or download binaries from a release version here: https://github.com/stevekm/dumpSoftwareVersions/releases

Or build from source using Go version 1.20+;

```bash
go build -o ./dumpSoftwareVersions ./main.go
```

# Notes

See full usage in a Nextflow pipeline here:

- https://github.com/stevekm/nextflow-demos/tree/master/dumpsoftwareversions

Inspired by:

- https://github.com/nf-core/rnaseq/blob/master/modules/nf-core/custom/dumpsoftwareversions/templates/dumpsoftwareversions.py

For use with:

- https://multiqc.info/docs/reports/customisation/#listing-software-versions