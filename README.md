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

## With Nextflow + MultiQC

You can see an example of how to use this utility inside your Nextflow pipeline [here](https://github.com/stevekm/nextflow-demos/blob/master/dumpsoftwareversions/modules/local/steve-dumpSoftwareVersions.nf).

A simple example [Nextflow process](https://www.nextflow.io/docs/latest/process.html) might look like this

```groovy
process DUMPSOFTWAREVERSIONS {
    container "stevekm/dump-software-versions:0.1"

    input:
    path(versionsYAMLFile) // single file with collated version YAML from all previous tools in the pipeline

    output:
    path(output_filename), emit: mqc_yml

    script:
    output_filename = "software_versions_mqc.yml"
    """
    dumpSoftwareVersions \
    -manifestName "${workflow.manifest.name}" \
    -manifestVersion "${workflow.manifest.version}" \
    -nxfVersion "${workflow.nextflow.version}" \
    -processLabel "${task.process}" \
    "${versionsYAMLFile}" > "${output_filename}"
    """
}
```

The output file here `software_versions_mqc.yml` can then be passed as an input item to your [MultiQC Nextflow process](https://github.com/stevekm/nextflow-demos/blob/master/dumpsoftwareversions/modules/nf-core/multiqc/main.nf).

Once incorporated into MultiQC, you should get a nice table that looks something like this;

<img width="1024" alt="Screenshot" src="https://github.com/stevekm/dumpSoftwareVersions/assets/10505524/569b0115-c9ed-4552-801d-33de68110d79">

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