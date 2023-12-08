package db

import (
	"strings"

	"github.com/flanksource/canary-checker/pkg"
	"github.com/flanksource/canary-checker/pkg/db/types"
	"github.com/flanksource/duty/context"
	"github.com/flanksource/duty/models"
	"github.com/google/uuid"
)

func FindConfigIDsByNameNamespaceType(ctx context.Context, namespace, name, configType string) ([]uuid.UUID, error) {
	if name == "" && namespace == "" && configType == "" {
		return nil, nil
	}

	var ids []uuid.UUID
	query := Gorm.Model(&models.ConfigItem{}).Select("id")
	if name != "" {
		query = query.Where("name = ?", name)
	}
	if namespace != "" {
		query = query.Where("namespace = ?", namespace)
	}
	if configType != "" {
		query = query.Where("type = ?", configType)
	}
	if err := query.Find(&ids).Error; err != nil {
		return nil, err
	}

	return ids, nil
}

func FindComponentIDsByNameNamespaceType(ctx context.Context, namespace, name, componentType string) ([]uuid.UUID, error) {
	if name == "" && namespace == "" && componentType == "" {
		return nil, nil
	}

	var ids []uuid.UUID
	query := Gorm.Model(&models.Component{}).Select("id")
	if name != "" {
		query = query.Where("name = ?", name)
	}
	if namespace != "" {
		query = query.Where("namespace = ?", namespace)
	}
	if componentType != "" {
		query = query.Where("type = ?", componentType)
	}
	if err := query.Find(&ids).Error; err != nil {
		return nil, err
	}

	return ids, nil
}

func GetLabelsFromSelector(selector string) (matchLabels map[string]string) {
	matchLabels = make(types.JSONStringMap)
	labels := strings.Split(selector, ",")
	for _, label := range labels {
		if strings.Contains(label, "=") {
			kv := strings.Split(label, "=")
			if len(kv) == 2 {
				matchLabels[kv[0]] = kv[1]
			} else {
				matchLabels[kv[0]] = ""
			}
		}
	}
	return
}

func GetComponentsWithLabelSelector(labelSelector string) (components pkg.Components, err error) {
	if labelSelector == "" {
		return nil, nil
	}
	var uninqueComponents = make(map[string]*pkg.Component)
	matchLabels := GetLabelsFromSelector(labelSelector)
	var labels = make(map[string]string)
	var onlyKeys []string
	for k, v := range matchLabels {
		if v != "" {
			labels[k] = v
		} else {
			onlyKeys = append(onlyKeys, k)
		}
	}
	var comps pkg.Components
	if err := Gorm.Table("components").
		Where("labels @> ?", types.JSONStringMap(labels)).
		Where("agent_id = '00000000-0000-0000-0000-000000000000'").
		Where("deleted_at IS NULL").
		Find(&comps).Error; err != nil {
		return nil, err
	}
	for _, c := range comps {
		uninqueComponents[c.ID.String()] = c
	}
	for _, k := range onlyKeys {
		var comps pkg.Components
		if err := Gorm.Table("components").
			Where("labels ?? ?", k).
			Where("agent_id = '00000000-0000-0000-0000-000000000000'").
			Where("deleted_at IS NULL").
			Find(&comps).Error; err != nil {
			continue
		}
		for _, c := range comps {
			uninqueComponents[c.ID.String()] = c
		}
	}
	for _, c := range uninqueComponents {
		components = append(components, c)
	}
	return components, nil
}

func GetComponentsWithFieldSelector(fieldSelector string) (components pkg.Components, err error) {
	if fieldSelector == "" {
		return nil, nil
	}
	var uninqueComponents = make(map[string]*pkg.Component)
	matchLabels := GetLabelsFromSelector(fieldSelector)
	for k, v := range matchLabels {
		var comp pkg.Components
		Gorm.Raw("select * from lookup_component_by_property(?, ?)", k, v).Scan(&comp)
		for _, c := range comp {
			uninqueComponents[c.ID.String()] = c
		}
	}
	for _, c := range uninqueComponents {
		components = append(components, c)
	}
	return
}
