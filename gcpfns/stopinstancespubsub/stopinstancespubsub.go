// Package p contains a Pub/Sub Cloud Function.
package p

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
	"time"

	"cloud.google.com/go/pubsub"
	"google.golang.org/api/compute/v1"
	"google.golang.org/api/option"
)

// {"project": "xyz", "label": "shutdown-nightly"}
type message struct {
	Project string `json:"project"`
	Label   string `json:"label"`
}

func newMessage(psm *pubsub.Message) (*message, error) {
	m := message{}
	if err := json.Unmarshal(psm.Data, &m); err != nil {
		return nil, fmt.Errorf("unmarshaling %q: %w", string(psm.Data), err)
	}

	return &m, nil
}

// StopInstancesPubSub consumes a Pub/Sub message.
func StopInstancesPubSub(ctx context.Context, psm pubsub.Message) error {
	m, err := newMessage(&psm)
	if err != nil {
		return fmt.Errorf("create new message: %w", err)
	}

	var (
		log              = newInfoLog(os.Stdout)
		projectName      = m.Project
		labelName        = m.Label
		instancesFilters = strings.Join([]string{
			"status = RUNNING",
		}, " ")
	)

	log.Infof("processing label %q in %s", labelName, projectName)

	csvc, err := compute.NewService(ctx, option.WithScopes(compute.ComputeScope))
	if err != nil {
		return fmt.Errorf("create compute service: %w", err)
	}
	issvc := compute.NewInstancesService(csvc)

	zListCall := csvc.Zones.List(projectName)
	zList, err := zListCall.Do()
	if err != nil {
		return fmt.Errorf("get zones list: %w", err)
	}

	log.Infof("found %d zones", len(zList.Items))

	inZonesInstancesListCallErrs := "Instance List Call Errors"
	instanceStopCallErrs := "Instance Stop Call Errors"
	groupedErrs := groupedErrors{
		inZonesInstancesListCallErrs: nil,
		instanceStopCallErrs:         nil,
	}

	for _, z := range zList.Items {
		iListCall := csvc.Instances.List(projectName, z.Name)
		iListCall.Filter(instancesFilters)
		iList, err := iListCall.Do()
		if err != nil {
			groupedErrs.add(inZonesInstancesListCallErrs, err)
			continue
		}

		log.Infof("found %d instances in zone %s", len(iList.Items), z.Name)

		for _, inst := range iList.Items {
			log.Infof("found instance %s", inst.Name)

			if shutdownText, ok := inst.Labels[labelName]; ok {
				log.Infof("shutdown label found: %s", shutdownText)

				shouldShutdown, err := strconv.ParseBool(shutdownText)
				if !shouldShutdown || err != nil {
					continue
				}

				iStopCall := issvc.Stop(projectName, z.Name, inst.Name)
				_, err = iStopCall.Do()
				if err != nil {
					groupedErrs.add(instanceStopCallErrs, err)
					continue
				}

				log.Infof("stop called successfully for %s+%s", z.Name, inst.Name)
			}
		}
	}

	return groupedErrs.Err()
}

type groupedErrors map[string][]error

func (em groupedErrors) add(key string, err error) {
	em[key] = append(em[key], err)
}

func (em groupedErrors) Err() error {
	var msgs []string
	for k, v := range em {
		if v == nil {
			continue
		}

		var errsMsgs []string
		for _, err := range v {
			if v == nil {
				continue
			}

			errsMsgs = append(errsMsgs, err.Error())
		}

		if errsMsgs == nil {
			continue
		}

		errsTxt := strings.Join(errsMsgs, "//")
		msg := fmt.Sprintf("%s: %s", k, errsTxt)
		msgs = append(msgs, msg)
	}

	if len(msgs) > 0 {
		return errors.New(strings.Join(msgs, " | "))
	}

	return nil
}

type infoLog struct {
	start time.Time
	out   io.Writer
}

func newInfoLog(out io.Writer) *infoLog {
	return &infoLog{
		start: time.Now(),
		out:   out,
	}
}

func (l *infoLog) Infof(format string, args ...interface{}) {
	fmt.Fprintf(l.out, "%#07.3f ", time.Since(l.start).Seconds())
	fmt.Fprintf(l.out, format, args...)
	fmt.Fprintln(l.out)
}
