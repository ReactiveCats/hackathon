// Code generated by entc, DO NOT EDIT.

package ent

import (
	"server/internal/ent/schema"
	"server/internal/ent/tag"
	"server/internal/ent/task"
	"time"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	tagFields := schema.Tag{}.Fields()
	_ = tagFields
	// tagDescMult is the schema descriptor for mult field.
	tagDescMult := tagFields[2].Descriptor()
	// tag.DefaultMult holds the default value on creation for the mult field.
	tag.DefaultMult = tagDescMult.Default.(float64)
	taskFields := schema.Task{}.Fields()
	_ = taskFields
	// taskDescCreatedAt is the schema descriptor for created_at field.
	taskDescCreatedAt := taskFields[0].Descriptor()
	// task.DefaultCreatedAt holds the default value on creation for the created_at field.
	task.DefaultCreatedAt = taskDescCreatedAt.Default.(func() time.Time)
	// taskDescIcon is the schema descriptor for icon field.
	taskDescIcon := taskFields[1].Descriptor()
	// task.DefaultIcon holds the default value on creation for the icon field.
	task.DefaultIcon = taskDescIcon.Default.(int)
	// task.IconValidator is a validator for the "icon" field. It is called by the builders before save.
	task.IconValidator = func() func(int) error {
		validators := taskDescIcon.Validators
		fns := [...]func(int) error{
			validators[0].(func(int) error),
			validators[1].(func(int) error),
		}
		return func(icon int) error {
			for _, fn := range fns {
				if err := fn(icon); err != nil {
					return err
				}
			}
			return nil
		}
	}()
	// taskDescTitle is the schema descriptor for title field.
	taskDescTitle := taskFields[2].Descriptor()
	// task.TitleValidator is a validator for the "title" field. It is called by the builders before save.
	task.TitleValidator = taskDescTitle.Validators[0].(func(string) error)
	// taskDescDescription is the schema descriptor for description field.
	taskDescDescription := taskFields[3].Descriptor()
	// task.DescriptionValidator is a validator for the "description" field. It is called by the builders before save.
	task.DescriptionValidator = taskDescDescription.Validators[0].(func(string) error)
	// taskDescEstimated is the schema descriptor for estimated field.
	taskDescEstimated := taskFields[5].Descriptor()
	// task.EstimatedValidator is a validator for the "estimated" field. It is called by the builders before save.
	task.EstimatedValidator = func() func(int) error {
		validators := taskDescEstimated.Validators
		fns := [...]func(int) error{
			validators[0].(func(int) error),
			validators[1].(func(int) error),
		}
		return func(estimated int) error {
			for _, fn := range fns {
				if err := fn(estimated); err != nil {
					return err
				}
			}
			return nil
		}
	}()
	// taskDescImportance is the schema descriptor for importance field.
	taskDescImportance := taskFields[6].Descriptor()
	// task.DefaultImportance holds the default value on creation for the importance field.
	task.DefaultImportance = taskDescImportance.Default.(int8)
	// task.ImportanceValidator is a validator for the "importance" field. It is called by the builders before save.
	task.ImportanceValidator = func() func(int8) error {
		validators := taskDescImportance.Validators
		fns := [...]func(int8) error{
			validators[0].(func(int8) error),
			validators[1].(func(int8) error),
		}
		return func(importance int8) error {
			for _, fn := range fns {
				if err := fn(importance); err != nil {
					return err
				}
			}
			return nil
		}
	}()
	// taskDescUrgency is the schema descriptor for urgency field.
	taskDescUrgency := taskFields[7].Descriptor()
	// task.DefaultUrgency holds the default value on creation for the urgency field.
	task.DefaultUrgency = taskDescUrgency.Default.(int8)
	// task.UrgencyValidator is a validator for the "urgency" field. It is called by the builders before save.
	task.UrgencyValidator = func() func(int8) error {
		validators := taskDescUrgency.Validators
		fns := [...]func(int8) error{
			validators[0].(func(int8) error),
			validators[1].(func(int8) error),
		}
		return func(urgency int8) error {
			for _, fn := range fns {
				if err := fn(urgency); err != nil {
					return err
				}
			}
			return nil
		}
	}()
	// taskDescCustomMult is the schema descriptor for custom_mult field.
	taskDescCustomMult := taskFields[8].Descriptor()
	// task.DefaultCustomMult holds the default value on creation for the custom_mult field.
	task.DefaultCustomMult = taskDescCustomMult.Default.(float64)
}
