package teamcity

import (
	// "errors"
	"fmt"

	"github.com/hashicorp/terraform/helper/hashcode"
	"github.com/hashicorp/terraform/helper/schema"

	"github.com/bestseller-ecom/teamcity-sdk-go/teamcity"
	"github.com/bestseller-ecom/teamcity-sdk-go/types"

	"log"
	"reflect"
)

func resourceBuildStep() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"type": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"properties": &schema.Schema{
				Type:             schema.TypeMap,
				Optional:         true,
				DiffSuppressFunc: optionalSuffixReturnSuppress,
			},
		},
	}
}

func resourceBuildSetting() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"value": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func resourceBuildSharedObject() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"type": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"properties": &schema.Schema{
				Type:     schema.TypeMap,
				Optional: true,
			},
		},
	}
}

func resourceBuildSharedDependencyObject() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"type": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"dependent": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"properties": &schema.Schema{
				Type:     schema.TypeMap,
				Optional: true,
			},
		},
	}
}

func resourceAttachedVcsRoot() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"vcs_root": &schema.Schema{
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: teamcity.ValidateVcsRootID,
			},
			"checkout_rules": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func resourceBuildConfiguration() *schema.Resource {
	return &schema.Resource{
		Create: resourceBuildConfigurationCreate,
		Read:   resourceBuildConfigurationRead,
		Update: resourceBuildConfigurationUpdate,
		Delete: resourceBuildConfigurationDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"project": &schema.Schema{
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: teamcity.ValidateID,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"description": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"setting": &schema.Schema{
				Type:     schema.TypeList,
				Elem:     resourceBuildSetting(),
				Optional: true,
			},
			"template": &schema.Schema{
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: teamcity.ValidateID,
			},
			"attached_vcs_root": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     resourceAttachedVcsRoot(),
				Set:      attachedVcsRootValueHash,
				Optional: true,
			},
			"parameter": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     resourceParameter(),
				Set:      parameterValueHash,
				Optional: true,
			},
			"parameter_values": &schema.Schema{
				Type:     schema.TypeMap,
				Optional: true,
			},
			"step": &schema.Schema{
				Type:     schema.TypeList,
				Elem:     resourceBuildStep(),
				Optional: true,
			},
			"feature": &schema.Schema{
				Type:     schema.TypeList,
				Elem:     resourceBuildSharedObject(),
				Optional: true,
			},
			"trigger": &schema.Schema{
				Type:     schema.TypeList,
				Elem:     resourceBuildSharedObject(),
				Optional: true,
			},
			"snapshot_dependency": &schema.Schema{
				Type:     schema.TypeList,
				Elem:     resourceBuildSharedDependencyObject(),
				Optional: true,
			},
			"artifact_dependency": &schema.Schema{
				Type:     schema.TypeList,
				Elem:     resourceBuildSharedDependencyObject(),
				Optional: true,
			},
			"agent_requirement": &schema.Schema{
				Type:     schema.TypeList,
				Elem:     resourceBuildSharedObject(),
				Optional: true,
			},
		},
	}
}

func resourceBuildTemplate() *schema.Resource {
	return &schema.Resource{
		Create: resourceBuildTemplateCreate,
		Read:   resourceBuildTemplateRead,
		Update: resourceBuildTemplateUpdate,
		Delete: resourceBuildTemplateDelete,

		Schema: map[string]*schema.Schema{
			"project": &schema.Schema{
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: teamcity.ValidateID,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"description": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"setting": &schema.Schema{
				Type:     schema.TypeList,
				Elem:     resourceBuildSetting(),
				Optional: true,
			},
			"attached_vcs_root": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     resourceAttachedVcsRoot(),
				Set:      attachedVcsRootValueHash,
				Optional: true,
			},
			"parameter": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     resourceParameter(),
				Set:      parameterValueHash,
				Optional: true,
			},
			"parameter_values": &schema.Schema{
				Type:     schema.TypeMap,
				Optional: true,
			},
			"step": &schema.Schema{
				Type:     schema.TypeList,
				Elem:     resourceBuildStep(),
				Optional: true,
			},
			"feature": &schema.Schema{
				Type:     schema.TypeList,
				Elem:     resourceBuildSharedObject(),
				Optional: true,
			},
			"trigger": &schema.Schema{
				Type:     schema.TypeList,
				Elem:     resourceBuildSharedObject(),
				Optional: true,
			},
			"snapshot_dependency": &schema.Schema{
				Type:     schema.TypeList,
				Elem:     resourceBuildSharedDependencyObject(),
				Optional: true,
			},
			"artifact_dependency": &schema.Schema{
				Type:     schema.TypeList,
				Elem:     resourceBuildSharedDependencyObject(),
				Optional: true,
			},
			"agent_requirement": &schema.Schema{
				Type:     schema.TypeList,
				Elem:     resourceBuildSharedObject(),
				Optional: true,
			},
		},
	}
}

func resourceBuildConfigurationCreate(d *schema.ResourceData, meta interface{}) error {
	return resourceBuildConfigurationCreateInternal(d, meta, false)
}

func resourceBuildConfigurationRead(d *schema.ResourceData, meta interface{}) error {
	return resourceBuildConfigurationReadInternal(d, meta, false)
}

func resourceBuildConfigurationUpdate(d *schema.ResourceData, meta interface{}) error {
	return resourceBuildConfigurationUpdateInternal(d, meta, false)
}

func resourceBuildConfigurationDelete(d *schema.ResourceData, meta interface{}) error {
	return resourceBuildConfigurationDeleteInternal(d, meta, false)
}

func resourceBuildTemplateCreate(d *schema.ResourceData, meta interface{}) error {
	return resourceBuildConfigurationCreateInternal(d, meta, true)
}

func resourceBuildTemplateRead(d *schema.ResourceData, meta interface{}) error {
	return resourceBuildConfigurationReadInternal(d, meta, true)
}

func resourceBuildTemplateUpdate(d *schema.ResourceData, meta interface{}) error {
	return resourceBuildConfigurationUpdateInternal(d, meta, true)
}

func resourceBuildTemplateDelete(d *schema.ResourceData, meta interface{}) error {
	return resourceBuildConfigurationDeleteInternal(d, meta, true)
}

/*
	ID                   string                    `json:"id,omitempty"`
	ProjectID            string                    `json:"projectId"`
	TemplateFlag         bool                      `json:"templateFlag"`
	TemplateID           TemplateId                `json:"template,omitempty"`
	Name                 string                    `json:"name"`
	Description          string                    `json:"description,omitempty"`
	VcsRootEntries       VcsRootEntries            `json:"vcs-root-entries,omitempty"`
	Settings             BuildSettings             `json:"settings,omitempty"`
	Parameters           Parameters                `json:"parameters,omitempty"`
	Steps                BuildSteps                `json:"steps,omitempty"`
	Features             BuildFeatures             `json:"features,omitempty"`
	Triggers             BuildTriggers             `json:"triggers,omitempty"`
	SnapshotDependencies BuildSnapshotDependencies `json:"snapshot-dependencies,omitempty"`
	ArtifactDependencies BuildArtifactDependencies `json:"artifact-dependencies,omitempty"`
	AgentRequirements    BuildAgentRequirements    `json:"agent-requirements,omitempty"`
*/

func resourceBuildConfigurationCreateInternal(d *schema.ResourceData, meta interface{}, template bool) error {
	client := meta.(*teamcity.Client)

	projectID := d.Get("project").(string)
	name := d.Get("name").(string)
	steps := resourceBuildSteps(d.Get("step").([]interface{}))
	settings := resourceBuildSettings(d.Get("setting").([]interface{}))
	features := resourceBuildFeatures(d.Get("feature").([]interface{}))
	triggers := resourceBuildTriggers(d.Get("trigger").([]interface{}))
	snapshotDependencies := resourceBuildSnapshotDependencies(d.Get("snapshot_dependency").([]interface{}))
	artifactDependencies := resourceBuildArtifactDependencies(d.Get("artifact_dependency").([]interface{}))
	agentRequirements := resourceBuildAgentRequirements(d.Get("agent_requirement").([]interface{}))
	d.Partial(true)
	templateID := ""
	if !template {
		templateID = d.Get("template").(string)
	}
	template_parameters := make(types.Parameters)
	if templateID != "" {
		if template_config, err := client.GetBuildConfiguration(templateID); err != nil {
			return err
		} else {
			template_parameters = template_config.Parameters
			if len(template_config.Steps) > 0 && len(steps) > 0 {
				return fmt.Errorf("Can't combine build config steps and template build steps %s", name)
			}
		}
	}

	var project_parameters types.Parameters
	if project, err := client.GetProject(projectID); err != nil {
		return err
	} else {
		project_parameters = project.Parameters
	}

	parameters := definitionToParameters(*d.Get("parameter").(*schema.Set))
	for name, _ := range parameters {
		if project_parameter, ok := project_parameters[name]; ok && project_parameter.Spec != nil {
			return fmt.Errorf("Can't redefine project parameter %s", name)
		}
		if template_parameter, ok := template_parameters[name]; ok && template_parameter.Spec != nil {
			return fmt.Errorf("Can't redefine template parameter %s", name)
		}
	}

	config := types.BuildConfiguration{
		ProjectID:            projectID,
		TemplateFlag:         template,
		TemplateID:           types.TemplateId(templateID),
		Name:                 name,
		Description:          d.Get("description").(string),
		Settings:             settings,
		Steps:                steps,
		Features:             features,
		Triggers:             triggers,
		SnapshotDependencies: snapshotDependencies,
		ArtifactDependencies: artifactDependencies,
		AgentRequirements:    agentRequirements,
	}

	if err := client.CreateBuildConfiguration(&config); err != nil {
		return err
	}
	id := config.ID
	d.SetId(id)
	d.SetPartial("project")
	d.SetPartial("name")
	d.SetPartial("description")
	if !template {
		d.SetPartial("template")
	}
	d.SetPartial("setting")
	d.SetPartial("step")
	d.SetPartial("feature")
	d.SetPartial("trigger")
	d.SetPartial("snapshot_dependency")
	d.SetPartial("artifact_dependency")
	d.SetPartial("agent_requirement")

	for name, v := range d.Get("parameter_values").(map[string]interface{}) {
		value := v.(string)
		parameter, ok := parameters[name]
		if !ok {
			if parameter, ok = project_parameters[name]; !ok {
				if parameter, ok = template_parameters[name]; !ok {
					parameter = types.Parameter{
						Value: value,
					}
				}
			}
		}
		parameter.Value = value
		parameters[name] = parameter
		log.Printf("Parameter value %s => %s", name, parameter.Value)
	}
	log.Printf("Replace Parameters value %q", parameters)
	if err := client.ReplaceAllBuildConfigurationParameters(id, &parameters); err != nil {
		return err
	}
	d.SetPartial("parameter_values")
	d.SetPartial("parameter")

	for _, root := range resourceAttachedVcsRoots(*d.Get("attached_vcs_root").(*schema.Set)) {
		err := client.AttachBuildConfigurationVcsRoot(id, &root)
		if err != nil {
			return err
		}
	}
	d.SetPartial("attached_vcs_root")

	d.Partial(false)
	return nil
}

func resourceBuildConfigurationReadInternal(d *schema.ResourceData, meta interface{}, template bool) error {
	log.Printf("Reading resource %q", d.Id())
	client := meta.(*teamcity.Client)
	config, err := client.GetBuildConfiguration(d.Id())
	if err != nil {
		return err
	}
	if config == nil || template != config.TemplateFlag {
		d.SetId("")
		return nil
	}

	log.Printf("Reading resource %q\n%q", d.Id(), d.Get("parameters"))
	d.Set("project", config.ProjectID)
	d.Set("name", config.Name)
	d.Set("description", config.Description)
	if !template {
		d.Set("template", config.TemplateID)
	}

	templateID := string(config.TemplateID)
	template_parameters := make(types.Parameters)
	template_settings := make(types.BuildSettings, 0)
	template_steps := make(types.BuildSteps, 0)
	template_features := make(types.BuildFeatures, 0)
	template_triggers := make(types.BuildTriggers, 0)
	template_snapshot_dependencies := make(types.BuildSnapshotDependencies, 0)
	template_artifact_dependencies := make(types.BuildArtifactDependencies, 0)
	template_agent_requirements := make(types.BuildAgentRequirements, 0)

	template_vcs_roots := make(types.VcsRootEntries, 0)
	if templateID != "" {
		if template_config, err := client.GetBuildConfiguration(templateID); err != nil {
			return err
		} else {
			template_parameters = template_config.Parameters
			template_settings = template_config.Settings
			template_steps = template_config.Steps
			template_features = template_config.Features
			template_triggers = template_config.Triggers
			template_snapshot_dependencies = template_config.SnapshotDependencies
			template_artifact_dependencies = template_config.ArtifactDependencies
			template_agent_requirements = template_config.AgentRequirements

			template_vcs_roots = template_config.VcsRootEntries
		}
	}

	settings := make([]map[string]interface{}, 0)
	for _, setting := range config.Settings {
		inTemplate := false
		for _, template_setting := range template_settings {
			if setting.Name == template_setting.Name {
				inTemplate = true
				break
			}
		}
		if inTemplate {
			continue
		}

		if setting.Name == "buildNumberCounter" {
			continue
		}

		v := make(map[string]interface{})
		v["name"] = setting.Name
		v["value"] = setting.Value

		settings = append(settings, v)
	}
	log.Printf("[INFO] Settings %q\n", settings)
	d.Set("setting", settings)

	steps := make([]map[string]interface{}, 0)
	for _, step := range config.Steps {
		inTemplate := false
		for _, template_step := range template_steps {
			if step.ID == template_step.ID {
				inTemplate = true
				break
			}
		}
		if inTemplate {
			continue
		}
		v := make(map[string]interface{})
		v["type"] = step.Type
		if step.Name != "" {
			v["name"] = step.Name
		}
		properties := make(map[string]interface{})
		for name, prop := range step.Properties {
			if properties[name] == "teamcity.step.mode" {
				continue
			}

			properties[name] = prop
		}
		if len(properties) > 0 {
			v["properties"] = properties
		}
		steps = append(steps, v)
	}
	log.Printf("[INFO] Steps %q\n", steps)
	d.Set("step", steps)

	features := make([]map[string]interface{}, 0)
	for _, feature := range config.Features {
		inTemplate := false
		for _, template_feature := range template_features {
			if feature.ID == template_feature.ID {
				inTemplate = true
				break
			}
		}
		if inTemplate {
			continue
		}
		v := make(map[string]interface{})
		v["type"] = feature.Type
		properties := make(map[string]interface{})
		for name, prop := range feature.Properties {
			properties[name] = prop
		}
		if len(properties) > 0 {
			v["properties"] = properties
		}
		features = append(features, v)
	}
	log.Printf("[INFO] Features %q\n", features)
	d.Set("feature", features)

	triggers := make([]map[string]interface{}, 0)
	for _, trigger := range config.Triggers {
		inTemplate := false
		for _, template_trigger := range template_triggers {
			if trigger.ID == template_trigger.ID {
				inTemplate = true
				break
			}
		}
		if inTemplate {
			continue
		}
		v := make(map[string]interface{})
		v["type"] = trigger.Type
		properties := make(map[string]interface{})
		for name, prop := range trigger.Properties {
			properties[name] = prop
		}
		if len(properties) > 0 {
			v["properties"] = properties
		}
		triggers = append(triggers, v)
	}
	log.Printf("[INFO] Triggers %q\n", triggers)
	d.Set("trigger", triggers)

	snapshot_dependencies := make([]map[string]interface{}, 0)
	for _, snapshot_dependency := range config.SnapshotDependencies {
		inTemplate := false
		for _, template_snapshot_dependency := range template_snapshot_dependencies {
			if snapshot_dependency.ID == template_snapshot_dependency.ID {
				inTemplate = true
				break
			}
		}
		if inTemplate {
			continue
		}
		v := make(map[string]interface{})
		v["type"] = snapshot_dependency.Type
		if snapshot_dependency.SourceBuildType.ID != "" {
			v["dependent"] = snapshot_dependency.SourceBuildType.ID
		}
		properties := make(map[string]interface{})
		for name, prop := range snapshot_dependency.Properties {
			properties[name] = prop
		}
		if len(properties) > 0 {
			v["properties"] = properties
		}
		snapshot_dependencies = append(snapshot_dependencies, v)
	}
	log.Printf("[INFO] SnapshotDependencies %q\n", snapshot_dependencies)
	d.Set("snapshot_dependency", snapshot_dependencies)

	artifact_dependencies := make([]map[string]interface{}, 0)
	for _, artifact_dependency := range config.ArtifactDependencies {
		inTemplate := false
		for _, template_artifact_dependency := range template_artifact_dependencies {
			if artifact_dependency.ID == template_artifact_dependency.ID {
				inTemplate = true
				break
			}
		}
		if inTemplate {
			continue
		}
		v := make(map[string]interface{})
		v["type"] = artifact_dependency.Type
		if artifact_dependency.SourceBuildType.ID != "" {
			v["dependent"] = artifact_dependency.SourceBuildType.ID
		}
		properties := make(map[string]interface{})
		for name, prop := range artifact_dependency.Properties {
			properties[name] = prop
		}
		if len(properties) > 0 {
			v["properties"] = properties
		}
		artifact_dependencies = append(artifact_dependencies, v)
	}
	log.Printf("[INFO] ArtifactDependencies %q\n", artifact_dependencies)
	d.Set("artifact_dependency", artifact_dependencies)

	agent_requirements := make([]map[string]interface{}, 0)
	for _, agent_requirement := range config.AgentRequirements {
		inTemplate := false
		for _, template_agent_requirement := range template_agent_requirements {
			if agent_requirement.ID == template_agent_requirement.ID {
				inTemplate = true
				break
			}
		}
		if inTemplate {
			continue
		}
		v := make(map[string]interface{})
		v["type"] = agent_requirement.Type
		properties := make(map[string]interface{})
		for name, prop := range agent_requirement.Properties {
			properties[name] = prop
		}
		if len(properties) > 0 {
			v["properties"] = properties
		}
		agent_requirements = append(agent_requirements, v)
	}
	log.Printf("[INFO] AgentRequirements %q\n", agent_requirements)
	d.Set("agent_requirement", agent_requirements)

	var project_parameters types.Parameters
	if project, err := client.GetProject(string(config.ProjectID)); err != nil {
		return err
	} else {
		project_parameters = project.Parameters
	}
	parameters := config.Parameters
	values := make(map[string]interface{})
	current := d.Get("parameter_values").(map[string]interface{})
	for name, parameter := range config.Parameters {
		if project_parameter, ok := project_parameters[name]; ok {
			if project_parameter.Value != parameter.Value {
				values[name] = parameter.Value
			}
			if project_parameter.Spec != nil || parameter.Spec == nil {
				delete(parameters, name)
			}
		} else if template_parameter, ok := template_parameters[name]; ok {
			if template_parameter.Value != parameter.Value {
				values[name] = parameter.Value
			}
			if template_parameter.Spec != nil || parameter.Spec == nil {
				delete(parameters, name)
			}
		} else {
			if parameter.Spec == nil {
				delete(parameters, name)
			}
			pwt := types.PasswordType{}
			if parameter.Value != "" {
				values[name] = parameter.Value
			} else if parameter.Spec != nil && parameter.Spec.Type == pwt {
				if value, ok := current[name]; ok && value != "" {
					values[name] = value
				}
			}
		}
	}
	d.Set("parameter", parametersToDefinition(parameters))
	d.Set("parameter_values", values)

	vcs_roots := make([]interface{}, 0)
	for _, root := range config.VcsRootEntries {
		inTemplate := false
		for _, template_vcs_root := range template_vcs_roots {
			if root.ID == template_vcs_root.ID {
				inTemplate = true
				break
			}
		}
		if inTemplate {
			continue
		}

		v := make(map[string]interface{})
		v["vcs_root"] = string(root.VcsRootID)
		v["checkout_rules"] = root.CheckoutRules
		vcs_roots = append(vcs_roots, v)
	}
	d.Set("attached_vcs_root", schema.NewSet(attachedVcsRootValueHash, vcs_roots))

	return nil
}

func resourceBuildSettings(settings []interface{}) types.BuildSettings {
	tcSettings := make(types.BuildSettings, 0)
	for _, s := range settings {
		property := s.(map[string]interface{})
		name := property["name"].(string)
		value := property["value"].(string)

		tcSettings = append(tcSettings, types.BuildSetting{
			Name:  name,
			Value: value,
		})
	}

	return tcSettings
}

func resourceBuildSteps(steps []interface{}) types.BuildSteps {
	tcSteps := make(types.BuildSteps, 0)
	for _, s := range steps {
		step := s.(map[string]interface{})
		typeName := step["type"].(string)
		name := step["name"].(string)
		properties := step["properties"].(map[string]interface{})
		actualProps := make(types.Properties)
		for k, v := range properties {
			actualProps[k] = v.(string)
		}

		tcSteps = append(tcSteps, types.BuildStep{
			Type:       typeName,
			Name:       name,
			Properties: actualProps,
		})
	}

	return tcSteps
}

func resourceBuildFeatures(features []interface{}) types.BuildFeatures {
	tcFeatures := make(types.BuildFeatures, 0)
	for _, s := range features {
		feature := s.(map[string]interface{})
		typeName := feature["type"].(string)
		properties := feature["properties"].(map[string]interface{})
		actualProps := make(types.Properties)
		for k, v := range properties {
			actualProps[k] = v.(string)
		}

		tcFeatures = append(tcFeatures, types.BuildFeature{
			Type:       typeName,
			Properties: actualProps,
		})
	}

	return tcFeatures
}

func resourceBuildTriggers(triggers []interface{}) types.BuildTriggers {
	tcTriggers := make(types.BuildTriggers, 0)
	for _, s := range triggers {
		trigger := s.(map[string]interface{})
		typeName := trigger["type"].(string)
		properties := trigger["properties"].(map[string]interface{})
		actualProps := make(types.Properties)
		for k, v := range properties {
			actualProps[k] = v.(string)
		}

		tcTriggers = append(tcTriggers, types.BuildTrigger{
			Type:       typeName,
			Properties: actualProps,
		})
	}

	return tcTriggers
}

func resourceBuildSnapshotDependencies(snapshotDependencies []interface{}) types.BuildSnapshotDependencies {
	tcSnapshotDependencies := make(types.BuildSnapshotDependencies, 0)
	for _, s := range snapshotDependencies {
		snapshotDependency := s.(map[string]interface{})
		typeName := snapshotDependency["type"].(string)
		dependent := snapshotDependency["dependent"].(string)
		properties := snapshotDependency["properties"].(map[string]interface{})
		actualProps := make(types.Properties)
		for k, v := range properties {
			actualProps[k] = v.(string)
		}

		tcSnapshotDependencies = append(tcSnapshotDependencies, types.BuildSnapshotDependency{
			Type:       typeName,
			Properties: actualProps,
			SourceBuildType: types.BuildType{
				ID: dependent,
			},
		})
	}

	return tcSnapshotDependencies
}

func resourceBuildArtifactDependencies(artifactDependencies []interface{}) types.BuildArtifactDependencies {
	tcArtifactDependencies := make(types.BuildArtifactDependencies, 0)
	for _, s := range artifactDependencies {
		artifactDependency := s.(map[string]interface{})
		typeName := artifactDependency["type"].(string)
		dependent := artifactDependency["dependent"].(string)
		properties := artifactDependency["properties"].(map[string]interface{})
		actualProps := make(types.Properties)
		for k, v := range properties {
			actualProps[k] = v.(string)
		}

		tcArtifactDependencies = append(tcArtifactDependencies, types.BuildArtifactDependency{
			Type:       typeName,
			Properties: actualProps,
			SourceBuildType: types.BuildType{
				ID: dependent,
			},
		})
	}

	return tcArtifactDependencies
}

func resourceBuildAgentRequirements(agentRequirements []interface{}) types.BuildAgentRequirements {
	tcAgentRequirements := make(types.BuildAgentRequirements, 0)
	for _, s := range agentRequirements {
		agentRequirement := s.(map[string]interface{})
		typeName := agentRequirement["type"].(string)
		properties := agentRequirement["properties"].(map[string]interface{})
		actualProps := make(types.Properties)
		for k, v := range properties {
			actualProps[k] = v.(string)
		}

		tcAgentRequirements = append(tcAgentRequirements, types.BuildAgentRequirement{
			Type:       typeName,
			Properties: actualProps,
		})
	}

	return tcAgentRequirements
}

func resourceAttachedVcsRoots(vcsRoots schema.Set) types.VcsRootEntries {
	keySet := schema.NewSet(attachedVcsRootKeyHash, vcsRoots.List())
	tcRoots := make(types.VcsRootEntries, 0)
	for _, s := range keySet.List() {
		entry := s.(map[string]interface{})
		vcsRoot := entry["vcs_root"].(string)
		rules := entry["checkout_rules"].(string)

		tcRoots = append(tcRoots, types.VcsRootEntry{
			VcsRootID:     types.VcsRootId(vcsRoot),
			CheckoutRules: rules,
		})
	}

	return tcRoots
}

func resourceBuildConfigurationUpdateInternal(d *schema.ResourceData, meta interface{}, template bool) error {
	client := meta.(*teamcity.Client)
	//var err error
	id := d.Id()
	d.Partial(true)

	steps := resourceBuildSteps(d.Get("step").([]interface{}))
	// settings := resourceBuildSettings(d.Get("setting").([]interface{}))
	features := resourceBuildFeatures(d.Get("feature").([]interface{}))
	triggers := resourceBuildTriggers(d.Get("trigger").([]interface{}))
	snapshotDependencies := resourceBuildSnapshotDependencies(d.Get("snapshot_dependency").([]interface{}))
	artifactDependencies := resourceBuildArtifactDependencies(d.Get("artifact_dependency").([]interface{}))
	agentRequirements := resourceBuildAgentRequirements(d.Get("agent_requirement").([]interface{}))

	templateID := ""
	if !template {
		templateID = d.Get("template").(string)
	}
	template_parameters := make(types.Parameters)
	if templateID != "" {
		if template_config, err := client.GetBuildConfiguration(templateID); err != nil {
			return err
		} else {
			template_parameters = template_config.Parameters
			if len(template_config.Steps) > 0 && len(steps) > 0 {
				return fmt.Errorf("Can't combine build config steps and template build steps")
			}
		}
	}

	if !template && d.HasChange("template") {
		if err := client.SetBuildConfigurationTemplate(id, ""); err != nil {
			return err
		}
		if templateID == "" {
			d.SetPartial("template")
		}
	}
	if d.HasChange("name") {
		o, n := d.GetChange("name")
		log.Printf("Updating Name '%s' => '%s'", o, n)
		if err := client.ReplaceBuildConfigurationField(id, "name", n.(string)); err != nil {
			return err
		}
		d.SetPartial("name")
	}
	if d.HasChange("setting") {
		o, n := d.GetChange("setting")
		for _, s := range o.([]interface{}) {
			property := s.(map[string]interface{})
			name := property["name"].(string)
			value := property["value"].(string)

			log.Printf("Removing Setting '%s' => '%s'", name, value)
			if err := client.DeleteBuildConfigurationSetting(id, name); err != nil {
				return err
			}

		}
		for _, s := range n.([]interface{}) {
			property := s.(map[string]interface{})
			name := property["name"].(string)
			value := property["value"].(string)

			log.Printf("Replacing Setting '%s' => '%s'", name, value)
			if err := client.DeleteBuildConfigurationSetting(id, name); err != nil {
				return err
			}
			if err := client.ReplaceBuildConfigurationSetting(id, name, value); err != nil {
				return err
			}
		}

		d.SetPartial("setting")
	}

	if d.HasChange("parameter") || (!template && d.HasChange("template")) {
		projectID := d.Get("project").(string)
		var project_parameters types.Parameters
		if project, err := client.GetProject(projectID); err != nil {
			return err
		} else {
			project_parameters = project.Parameters
		}

		o, n := d.GetChange("parameter")
		ops := o.(*schema.Set).Difference(n.(*schema.Set))
		nps := n.(*schema.Set).Difference(o.(*schema.Set))

		// log.Printf("Removing Parameters: %q", ops)
		// log.Printf("Replacing Parameters: %q", nps)

		parameters := definitionToParameters(*nps)
		old := definitionToParameters(*ops)
		replace_parameters := make(types.Parameters)
		delete_parameters := old
		for name, parameter := range parameters {

			if project_parameter, ok := project_parameters[name]; ok && project_parameter.Spec != nil {
				return fmt.Errorf("Can't redefine project parameter %s", name)
			}
			if template_parameter, ok := template_parameters[name]; ok && template_parameter.Spec != nil {
				return fmt.Errorf("Can't redefine template parameter %s", name)
			}
			if !reflect.DeepEqual(parameter, old[name]) {
				replace_parameters[name] = parameter
			}
			// This should never occur
			delete(delete_parameters, name)
		}
		for name, _ := range delete_parameters {
			if err := client.DeleteBuildConfigurationParameter(id, name); err != nil {
				return err
			}
		}
		for name, parameter := range replace_parameters {
			if err := client.DeleteBuildConfigurationParameter(id, name); err != nil {
				return err
			}
			if err := client.ReplaceBuildConfigurationParameter(id, name, &parameter); err != nil {
				return err
			}
		}
		d.SetPartial("parameter")
	}

	if d.HasChange("parameter_values") {
		o, n := d.GetChange("parameter_values")
		opv := o.(map[string]interface{})
		npv := n.(map[string]interface{})

		remove, replace := diffTypeMap(typeMapFromSchema(opv), typeMapFromSchema(npv))

		// log.Printf("[DEBUG] Removing Parameter_Values: %v", remove)
		// log.Printf("[DEBUG] Replacing Paramter_Values: %v", replace)

		for name, value := range remove {
			log.Printf("[DEBUG] Removing Parameter Value: '%s' => '%s'", name, value)
			if err := client.DeleteBuildConfigurationParameter(id, name); err != nil {
				return err
			}

		}

		for name, value := range replace {
			log.Printf("[DEBUG] Replacing Parameter Value: '%s' => '%s'", name, value)
			if err := client.ReplaceBuildConfigurationParameterValue(id, name, value); err != nil {
				return err
			}
		}

		d.SetPartial("parameter_values")
	}

	if d.HasChange("attached_vcs_root") {
		old, n := d.GetChange("attached_vcs_root")
		existing := make(map[types.VcsRootId]bool)

		for _, root := range resourceAttachedVcsRoots(*n.(*schema.Set)) {
			err := client.AttachBuildConfigurationVcsRoot(id, &root)
			if err != nil {
				return err
			}
			existing[root.VcsRootID] = true
		}
		for _, root := range resourceAttachedVcsRoots(*old.(*schema.Set)) {
			if !existing[root.VcsRootID] {
				err := client.DetachBuildConfigurationVcsRoot(id, string(root.VcsRootID))
				if err != nil {
					return err
				}
			}
		}

		d.SetPartial("attached_vcs_root")
	}

	if d.HasChange("description") {
		if err := client.SetBuildConfigurationDescription(d.Id(), d.Get("description").(string)); err != nil {
			return err
		}
		d.SetPartial("description")
	}

	if d.HasChange("step") {
		if err := client.ReplaceAllBuildConfigurationSteps(id, &steps); err != nil {
			return err
		}
		d.SetPartial("step")
	}
	if d.HasChange("feature") {
		if err := client.ReplaceAllBuildConfigurationFeatures(id, &features); err != nil {
			return err
		}
		d.SetPartial("feature")
	}
	if d.HasChange("trigger") {
		if err := client.ReplaceAllBuildConfigurationTriggers(id, &triggers); err != nil {
			return err
		}
		d.SetPartial("trigger")
	}
	if d.HasChange("snapshot_dependency") {
		if err := client.ReplaceAllBuildConfigurationSnapshotDependencies(id, &snapshotDependencies); err != nil {
			return err
		}
		d.SetPartial("snapshot_dependency")
	}
	if d.HasChange("artifact_dependency") {
		if err := client.ReplaceAllBuildConfigurationArtifactDependencies(id, &artifactDependencies); err != nil {
			return err
		}
		d.SetPartial("artifact_dependency")
	}
	if d.HasChange("agent_requirement") {
		if err := client.ReplaceAllBuildConfigurationAgentRequirements(id, &agentRequirements); err != nil {
			return err
		}
		d.SetPartial("agent_requirement")
	}
	if !template && d.HasChange("template") && templateID != "" {
		if err := client.SetBuildConfigurationTemplate(id, templateID); err != nil {
			return err
		}
	}

	d.Partial(false)
	return nil
}

func resourceBuildConfigurationDeleteInternal(d *schema.ResourceData, meta interface{}, template bool) error {
	client := meta.(*teamcity.Client)
	return client.DeleteBuildConfiguration(d.Id())
}

func attachedVcsRootValueHash(v interface{}) int {
	rd := v.(map[string]interface{})
	vcsRoot := rd["vcs_root"].(string)
	checkoutRules := rd["checkout_rules"].(string)
	hk := fmt.Sprintf("%s:%s", vcsRoot, checkoutRules)
	log.Printf("[DEBUG] TeamCity attachedVcsRootValueHash(%#v): %s: hk=%s,hc=%d", v, vcsRoot, hk, hashcode.String(hk))
	return hashcode.String(hk)
}

func attachedVcsRootKeyHash(v interface{}) int {
	m := v.(map[string]interface{})
	hk := m["vcs_root"].(string)
	log.Printf("[DEBUG] TeamCity attachedVcsRootKeyHash(%#v): %s: hk=%s,hc=%d", v, hk, hk, hashcode.String(hk))
	return hashcode.String(hk)
}

func diffTypeMap(oldMap, newMap map[string]string) (map[string]string, map[string]string) {
	for k, old := range oldMap {
		new, ok := newMap[k]
		if ok && old == new {
			delete(oldMap, k)
			delete(newMap, k)
		}
	}
	return oldMap, newMap
}

func typeMapFromSchema(m map[string]interface{}) map[string]string {
	result := make(map[string]string, len(m))
	for k, v := range m {
		result[k] = v.(string)
	}
	return result
}
