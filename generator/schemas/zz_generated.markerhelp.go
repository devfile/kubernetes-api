// +build !ignore_autogenerated

// Generated for the devfile generator

// Code generated by helpgen. DO NOT EDIT.

package schemas

import (
	"sigs.k8s.io/controller-tools/pkg/markers"
)

func (GenerateJSONSchema) Help() *markers.DefinitionHelp {
	return &markers.DefinitionHelp{
		Category: "Devfile",
		DetailedHelp: markers.DetailedHelp{
			Summary: "drives whether a Json schema should be generated from this GO Struct type",
			Details: "",
		},
		FieldHelp: map[string]markers.DetailedHelp{
			"OmitCustomUnionMembers": markers.DetailedHelp{
				Summary: "indicates that the Json schema gnerated from this type should omit Custom component union members.",
				Details: "",
			},
			"OmitPluginUnionMembers": markers.DetailedHelp{
				Summary: "indicates that the Json schema gnerated from this type should omit Plugin component union members.",
				Details: "",
			},
			"Title": markers.DetailedHelp{
				Summary: "indicates the content ot the Json Schema `title` attribute",
				Details: "",
			},
		},
	}
}

func (Generator) Help() *markers.DefinitionHelp {
	return &markers.DefinitionHelp{
		Category: "",
		DetailedHelp: markers.DetailedHelp{
			Summary: "generates JSON schemas from the GO source code of the Kubernetes API ",
			Details: "A JSON Schema is generated for each GO structure that had the `devfile:jsonschema:generate` annotation. The semver-compatible version of JSON Schemas is defined by the `devfile:jsonschema:version` annotation on the package. Typically in the `doc.go` file.",
		},
		FieldHelp: map[string]markers.DetailedHelp{},
	}
}
