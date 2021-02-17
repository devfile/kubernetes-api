package common

import (
	"fmt"
	"strconv"

	schema "github.com/devfile/api/v2/pkg/apis/workspaces/v1alpha2"
)

// componentAdded adds a new component to the test schema data and to the parser data
func (devfile *TestDevfile) componentAdded(component schema.Component) {
	LogInfoMessage(fmt.Sprintf("component added Name: %s", component.Name))
	devfile.SchemaDevFile.Components = append(devfile.SchemaDevFile.Components, component)
	if devfile.Follower != nil {
		devfile.Follower.AddComponent(component)
	}
}

// componetUpdated updates a component in the parser data
func (devfile *TestDevfile) componentUpdated(component schema.Component) {
	LogInfoMessage(fmt.Sprintf("component updated Name: %s", component.Name))
	if devfile.Follower != nil {
		devfile.Follower.UpdateComponent(component)
	}
}

// addVolume returns volumeMounts in a schema structure based on a specified number of volumes
func (devfile *TestDevfile) addVolume(numVols int) []schema.VolumeMount {
	commandVols := make([]schema.VolumeMount, numVols)
	for i := 0; i < numVols; i++ {
		volumeComponent := devfile.AddComponent(schema.VolumeComponentType)
		commandVols[i].Name = volumeComponent.Name
		commandVols[i].Path = "/Path_" + GetRandomString(5, false)
		LogInfoMessage(fmt.Sprintf("....... Add Volume: %s", commandVols[i]))
	}
	return commandVols
}

// AddComponent adds a component of the specified type, with random attributes, to the devfile schema
func (devfile *TestDevfile) AddComponent(componentType schema.ComponentType) schema.Component {

	var component schema.Component
	if componentType == schema.ContainerComponentType {
		component = devfile.createContainerComponent()
		devfile.SetContainerComponentValues(&component)
	} else if componentType == schema.VolumeComponentType {
		component = devfile.createVolumeComponent()
		devfile.SetVolumeComponentValues(&component)
	}
	return component
}

// createContainerComponent creates a container component, ready for attribute setting
func (devfile *TestDevfile) createContainerComponent() schema.Component {

	LogInfoMessage("Create a container component :")
	component := schema.Component{}
	component.Name = GetRandomUniqueString(8, true)
	LogInfoMessage(fmt.Sprintf("....... Name: %s", component.Name))
	component.Container = &schema.ContainerComponent{}
	devfile.componentAdded(component)
	return component

}

// createVolumeComponent creates a volume component , ready for attribute setting
func (devfile *TestDevfile) createVolumeComponent() schema.Component {

	LogInfoMessage("Create a volume component :")
	component := schema.Component{}
	component.Name = GetRandomUniqueString(8, true)
	LogInfoMessage(fmt.Sprintf("....... Name: %s", component.Name))
	component.Volume = &schema.VolumeComponent{}
	devfile.componentAdded(component)
	return component

}

// GetContainer returns the name of an existing, or newly created, container.
func (devfile *TestDevfile) GetContainerName() string {

	componentName := ""
	for _, currentComponent := range devfile.SchemaDevFile.Components {
		if currentComponent.Container != nil {
			componentName = currentComponent.Name
			LogInfoMessage(fmt.Sprintf("return existing container from GetContainerName  : %s", componentName))
			break
		}
	}

	if componentName == "" {
		component := devfile.createContainerComponent()
		component.Container.Image = GetRandomUniqueString(GetRandomNumber(8, 18), false)
		componentName = component.Name
		LogInfoMessage(fmt.Sprintf("retrun new container from GetContainerName : %s", componentName))
	}

	return componentName
}

// SetContainerComponentValues randomly sets container component attributes to random values
func (devfile *TestDevfile) SetContainerComponentValues(component *schema.Component) {

	containerComponent := component.Container

	containerComponent.Image = GetRandomUniqueString(GetRandomNumber(8, 18), false)

	if GetBinaryDecision() {
		numCommands := GetRandomNumber(1, 3)
		containerComponent.Command = make([]string, numCommands)
		for i := 0; i < numCommands; i++ {
			containerComponent.Command[i] = GetRandomString(GetRandomNumber(4, 16), false)
			LogInfoMessage(fmt.Sprintf("....... command %d of %d : %s", i, numCommands, containerComponent.Command[i]))
		}
	}

	if GetBinaryDecision() {
		numArgs := GetRandomNumber(1, 3)
		containerComponent.Args = make([]string, numArgs)
		for i := 0; i < numArgs; i++ {
			containerComponent.Args[i] = GetRandomString(GetRandomNumber(8, 18), false)
			LogInfoMessage(fmt.Sprintf("....... arg %d of %d : %s", i, numArgs, containerComponent.Args[i]))
		}
	}

	containerComponent.DedicatedPod = GetBinaryDecision()
	LogInfoMessage(fmt.Sprintf("....... DedicatedPod: %t", containerComponent.DedicatedPod))

	if GetBinaryDecision() {
		containerComponent.MemoryLimit = strconv.Itoa(GetRandomNumber(4, 124)) + "M"
		LogInfoMessage(fmt.Sprintf("....... MemoryLimit: %s", containerComponent.MemoryLimit))
	}

	if GetBinaryDecision() {
		setMountSources := GetBinaryDecision()
		containerComponent.MountSources = &setMountSources
		LogInfoMessage(fmt.Sprintf("....... MountSources: %t", *containerComponent.MountSources))

		if setMountSources {
			containerComponent.SourceMapping = "/" + GetRandomString(8, false)
			LogInfoMessage(fmt.Sprintf("....... SourceMapping: %s", containerComponent.SourceMapping))
		}
	}

	if GetBinaryDecision() {
		containerComponent.Env = addEnv(GetRandomNumber(1, 4))
	} else {
		containerComponent.Env = nil
	}

	if len(containerComponent.VolumeMounts) == 0 {
		if GetBinaryDecision() {
			containerComponent.VolumeMounts = devfile.addVolume(GetRandomNumber(1, 4))
		}
	}

	if GetBinaryDecision() {
		containerComponent.Endpoints = devfile.CreateEndpoints()
	}

	devfile.componentUpdated(*component)

}

// SetVolumeComponentValues randomly sets volume component attributes to random values
func (devfile *TestDevfile) SetVolumeComponentValues(component *schema.Component) {

	component.Volume.Size = strconv.Itoa(4+GetRandomNumber(64, 256)) + "G"
	LogInfoMessage(fmt.Sprintf("....... volumeComponent.Size: %s", component.Volume.Size))
	devfile.componentUpdated(*component)

}