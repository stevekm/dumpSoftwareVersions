# dumpSoftwareVersions

Tiny simple command line utility used to create a software versions YAML file from the various tools used throughout a [Nextflow](https://www.nextflow.io/) pipeline, in order to generate the Software Versions HTML table input used with [MultiQC](https://multiqc.info/) reports.

# Usage

```bash
$ ./dumpSoftwareVersions -manifestName dump-software-version-demo -manifestVersion 1.0 -nxfVersion 23.04.1 -processLabel CUSTOM_DUMPSOFTWAREVERSIONS example/collated_versions.yml
```

See full usage in a Nextflow pipeline here:

- https://github.com/stevekm/nextflow-demos/tree/master/dumpsoftwareversions

Inspired by:

- https://github.com/nf-core/rnaseq/blob/master/modules/nf-core/custom/dumpsoftwareversions/templates/dumpsoftwareversions.py

For use with:

- https://multiqc.info/docs/reports/customisation/#listing-software-versions