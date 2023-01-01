package setup

// Negative status(es)
const Failed = "FAILED"

// Nefative message(s)
const Empty = "EMPTY"
const NotModified = "NOT_MODIFIED"

// Positive status(es)
const Success = "SUCCESS"

// Positive message(s)
const Ok = "OK"

func ResponseForQuerying(status bool, message ...string) (string, string) {
	if status {
		if len(message) > 0 {
			return Success, message[0]
		}
		return Success, Ok
	}
	return Failed, Empty
}

func ResponseForModifying(status bool) (string, string) {
	if status {
		return Success, Ok
	}
	return Failed, NotModified
}