package core

import (
	"fmt"
)

/*
 * @Functions
 */

// SetID ...
func SetID(typ string, suffix interface{}) string {
	var r string
	switch typ {
	case "definitions":
		r = fmt.Sprintf("Definitions_%v", suffix)
	case "id":
		r = fmt.Sprintf("%s", suffix)
	}
	return r
}

/*
func FindActivitiesBySourceRef(definitions *TDefinitions, sourceRef string) []interface{} {

	var activities []interface{}

	for _, proc := range definitions.Process {
		for _, seqFlow := range proc.SequenceFlow {
			if seqFlow.SourceTargetRef.SourceRef == sourceRef {
				activity := FindActivityByID(definitions, seqFlow.SourceTargetRef.TargetRef)
				if activity != nil {
					activities = append(activities, activity)
				}
			}
		}
	}

	// return the found activities
	return activities

}

// FindactivityByID searches for an activity by its ID in the given definitions
func FindActivityByID(definitions *TDefinitions, activityID string) interface{} {

	// search for the activity in each process
	for _, process := range definitions.TDefinitionsBaseElements.Process {

		// search in each script task of the process
		for _, scriptTask := range process.TTasks.ScriptTask {
			if scriptTask.BaseAttributes.ID == activityID {
				return scriptTask
			}
		}

		// search in each task of the process
		for _, task := range process.TTasks.Task {
			if task.BaseAttributes.ID == activityID {
				return task
			}
		}

		// search in each user task of the process
		for _, userTask := range process.TTasks.UserTask {
			if userTask.BaseAttributes.ID == activityID {
				return userTask
			}
		}

		// search in each event of the process
		for _, startEvent := range process.TProcessEvents.StartEvent {
			if startEvent.BaseAttributes.ID == activityID {
				return startEvent
			}
		}

		// search in each endevent of the process
		for _, endEvent := range process.TProcessEvents.EndEvent {
			if endEvent.BaseAttributes.ID == activityID {
				return endEvent
			}
		}

	}

	// if the activity is not found, return nil
	return nil

}
*/
