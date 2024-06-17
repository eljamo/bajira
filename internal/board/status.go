package board

import (
	"github.com/eljamo/bajira/internal/errorconc"
	bajiraStrings "github.com/eljamo/bajira/internal/strings"
)

type BajiraStatusString string

const (
	BajiraStatusStringBacklog         BajiraStatusString = "backlog"
	BajiraStatusStringToDo            BajiraStatusString = "todo"
	BajiraStatusStringInProgress      BajiraStatusString = "in_progress"
	BajiraStatusStringInDevelopment   BajiraStatusString = "in_development"
	BajiraStatusStringInCodeReview    BajiraStatusString = "in_code_review"
	BajiraStatusStringReadyForTesting BajiraStatusString = "ready_for_testing"
	BajiraStatusStringQATesting       BajiraStatusString = "qa_testing"
	BajiraStatusStringUATesting       BajiraStatusString = "ua_testing"
	BajiraStatusStringReadyForRelease BajiraStatusString = "ready_for_release"
	BajiraStatusStringReleased        BajiraStatusString = "released"
	BajiraStatusStringDone            BajiraStatusString = "done"
	BajiraStatusStringClosed          BajiraStatusString = "closed"
	BajiraStatusStringInvalid         BajiraStatusString = "invalid"
)

func GetStatus(status string) (BajiraStatusString, error) {
	switch status {
	case bajiraStrings.Backlog:
		return BajiraStatusStringBacklog, nil
	case bajiraStrings.ToDo:
		return BajiraStatusStringToDo, nil
	case bajiraStrings.InProgress:
		return BajiraStatusStringInProgress, nil
	case bajiraStrings.InDevelopment:
		return BajiraStatusStringInDevelopment, nil
	case bajiraStrings.InCodeReview:
		return BajiraStatusStringInCodeReview, nil
	case bajiraStrings.ReadyForTesting:
		return BajiraStatusStringReadyForTesting, nil
	case bajiraStrings.QATesting:
		return BajiraStatusStringQATesting, nil
	case bajiraStrings.UATesting:
		return BajiraStatusStringUATesting, nil
	case bajiraStrings.ReadyForRelease:
		return BajiraStatusStringReadyForRelease, nil
	case bajiraStrings.Released:
		return BajiraStatusStringReleased, nil
	case bajiraStrings.Done:
		return BajiraStatusStringDone, nil
	case bajiraStrings.Closed:
		return BajiraStatusStringClosed, nil
	}

	return BajiraStatusStringInvalid, errorconc.LocalizedError(nil, "invalid board status", status)
}

func GetStatusString(status string) (string, error) {
	switch BajiraStatusString(status) {
	case BajiraStatusStringBacklog:
		return bajiraStrings.Backlog, nil
	case BajiraStatusStringToDo:
		return bajiraStrings.ToDo, nil
	case BajiraStatusStringInProgress:
		return bajiraStrings.InProgress, nil
	case BajiraStatusStringInDevelopment:
		return bajiraStrings.InDevelopment, nil
	case BajiraStatusStringInCodeReview:
		return bajiraStrings.InCodeReview, nil
	case BajiraStatusStringReadyForTesting:
		return bajiraStrings.ReadyForTesting, nil
	case BajiraStatusStringQATesting:
		return bajiraStrings.QATesting, nil
	case BajiraStatusStringUATesting:
		return bajiraStrings.UATesting, nil
	case BajiraStatusStringReadyForRelease:
		return bajiraStrings.ReadyForRelease, nil
	case BajiraStatusStringReleased:
		return bajiraStrings.Released, nil
	case BajiraStatusStringDone:
		return bajiraStrings.Done, nil
	case BajiraStatusStringClosed:
		return bajiraStrings.Closed, nil
	}

	return bajiraStrings.InvalidUpper, errorconc.LocalizedError(nil, "invalid board status", status)
}

var StatusList = []BajiraStatusString{
	BajiraStatusStringBacklog,
	BajiraStatusStringToDo,
	BajiraStatusStringInProgress,
	BajiraStatusStringInDevelopment,
	BajiraStatusStringInCodeReview,
	BajiraStatusStringReadyForTesting,
	BajiraStatusStringQATesting,
	BajiraStatusStringUATesting,
	BajiraStatusStringReadyForRelease,
	BajiraStatusStringReleased,
	BajiraStatusStringDone,
	BajiraStatusStringClosed,
}
