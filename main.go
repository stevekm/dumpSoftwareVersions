package main

import (
	"flag"
	"fmt"
	"gopkg.in/yaml.v3" // NOTE: make sure this is the same one used in GetYAMLibVersion
	"io/ioutil"
	"log"
	"sort"
	"strings"

	"runtime"
	"runtime/debug"
)

func LoadYAMLFile(filepath string, processLabel string, nxfVersion string, manifestName string, manifestVersion string) map[string]map[string]string {
	// load versions from YAML file
	var versions map[string]map[string]string

	yamlData, err := ioutil.ReadFile(filepath)
	if err != nil {
		log.Fatal(err)
	}

	err = yaml.Unmarshal(yamlData, &versions)
	if err != nil {
		log.Fatal(err)
	}

	// add some extra keys
	yamlModule := GetYAMLibVersion()
	versions[processLabel] = map[string]string{
		"Go":            runtime.Version(),
		yamlModule.Path: yamlModule.Version,
	}
	versions["Workflow"] = map[string]string{
		"Nextflow":   nxfVersion,
		manifestName: manifestVersion,
	}

	return versions
}

func AggregateByModule(versions map[string]map[string]string) map[string]map[string]string {
	// aggregate per-process; strip out subworkflow-strings and condense the map
	versionsByModule := make(map[string]map[string]string)

	for key, value := range versions {
		// Split the string by the colon delimiter
		keyParts := strings.Split(key, ":")

		// Get the last section of the string
		keyTrimmed := keyParts[len(keyParts)-1]

		// add it to the map if its missing
		if _, ok := versionsByModule[keyTrimmed]; !ok {
			versionsByModule[keyTrimmed] = value
		}
	}
	return versionsByModule
}

func GetSortedKeys[V any](mapping map[string]V) []string {
	// get the string keys from a map and sort them
	// https://go.dev/doc/tutorial/generics
	// https://stackoverflow.com/questions/25772347/handle-maps-with-same-key-type-but-different-value-type
	sortedKeys := []string{}
	for key := range mapping {
		sortedKeys = append(sortedKeys, key)
	}
	sort.Strings(sortedKeys)
	return sortedKeys
}

func MakeVersionHTML(versions map[string]map[string]string) string {
	// make an HTML table out of the entries
	var html = []string{
		`
<style>
#nf-core-versions tbody:nth-child(even) {
background-color: #f2f2f2;
}
</style>
<table class="table" style="width:100%" id="nf-core-versions">
<thead>
<tr>
<th> Process Name </th>
<th> Software </th>
<th> Version  </th>
</tr>
</thead>
`}

	sortedKeys := GetSortedKeys(versions)

	// iterate over the versions
	for _, process := range sortedKeys {
		tmp_versions := versions[process]
		html = append(html, "<tbody>")

		sortedValuesKeys := GetSortedKeys(tmp_versions)

		//  iterate over the versions values items
		for i, tool := range sortedValuesKeys {
			version := tmp_versions[tool]

			// I am not entirely sure why this part functions like this
			processString := ""
			if i == 0 {
				processString = process
			}

			htmlToolString := fmt.Sprintf(
				`
<tr>
<td><samp>%v</samp></td>
<td><samp>%v</samp></td>
<td><samp>%v</samp></td>
</tr>
`, processString, tool, version)
			html = append(html, htmlToolString)
		}
		html = append(html, "</tbody>")
	}
	html = append(html, "</table>")
	htmlString := strings.Join(html, "\n")
	return htmlString
}

func GetYAMLibVersion() debug.Module {
	bi, ok := debug.ReadBuildInfo()
	if !ok {
		log.Fatal("Failed to read build info")
	}

	for _, dep := range bi.Deps {
		if dep.Path == "gopkg.in/yaml.v3" {
			return *dep
		}
	}
	return debug.Module{}
}

func main() {
	// command line args
	processLabel := flag.String("processLabel", "default-process-label", "the name of the currently running Nextflow process (task.process)")
	nxfVersion := flag.String("nxfVersion", "default-nxf-version", " (workflow.nextflow.version)")
	manifestName := flag.String("manifestName", "default-manifest-name", " (workflow.manifest.name)")
	manifestVersion := flag.String("manifestVersion", "default-manifest-version", " (workflow.manifest.version)")
	flag.Parse()

	// first positional arg passed
	inputYAMLFilepath := flag.Arg(0)
	if inputYAMLFilepath == "" {
		log.Fatal("ERROR: Input YAML file path not provided")
	}

	// load the data
	versions := LoadYAMLFile(inputYAMLFilepath, *processLabel, *nxfVersion, *manifestName, *manifestVersion)

	// get per-module / process data
	versionsByModule := AggregateByModule(versions)

	// generate an HTML table for the mappings
	htmlString := MakeVersionHTML(versionsByModule)

	// generate final mapping for MultiQC
	versionsMQC := map[string]string{
		"section_name": fmt.Sprintf("%v Software Versions", *manifestName),
		"section_href": fmt.Sprintf("https://github.com/%v", *manifestName),
		"plot_type":    "html",
		"description":  "are collected at run time from the software output.",
		"data":         htmlString,
	}

	// convert to YAML data
	yamlData, err := yaml.Marshal(&versionsMQC)
	if err != nil {
		fmt.Printf("Error while Marshaling YAML from data; %v", err)
		log.Fatal(err)
	}

	// print the YAML output to stdout so it can be redirected as desired
	fmt.Println(string(yamlData))
}
