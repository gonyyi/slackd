package slackd

import (
	"errors"
	"fmt"
)

func wErr(prev error, newErrMsg string) error {
	if prev != nil {
		return fmt.Errorf("%s: %w", newErrMsg, prev)
	}
	return errors.New(newErrMsg)
}

func wErrIf(condition bool, prev error, newErrMsg string) error {
	if condition {
		return wErr(prev, newErrMsg)
	}
	return prev
}

func wErrs(prev error, curr error) error {
	if prev == nil && curr == nil {
		return nil
	}
	if prev == nil || curr == nil {
		if prev != nil {
			return prev
		}
		if curr != nil {
			return curr
		}
	}
	return fmt.Errorf("%v: %w", prev, curr)
}

func joinErrs(e ...error) error {
	var err error
	for i:=0; i<len(e); i++ {
		if e[i] != nil {
			if err != nil {
				err = fmt.Errorf("%v: %w", err, e[i])
			} else {
				err = e[i]
			}
		}
	}
	return err
}
