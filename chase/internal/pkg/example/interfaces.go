package example

import "fmt"

// Robot generic interface for a robot
type Robot interface {
	getPosition() string
	getDetails() string
}

// Details is the information about a robot
type Details struct {
	Name        string
	PhoneNumber string
}

// T1000 is a type of robot
type T1000 struct {
	Position string
	Strength int
	Details
}

// Data is a type of robot
type Data struct {
	Position     string
	Intelligence int
	Details
}

func (t *T1000) getPosition() string {
	return t.Position
}

func (t *T1000) getDetails() string {
	return fmt.Sprintf("%s has %d strength", t.Name, t.Strength)
}

func (d *Data) getPosition() string {
	return d.Position
}

func (d *Data) getDetails() string {
	return fmt.Sprintf("%s has %d intelligence", d.Name, d.Intelligence)
}

func getDetails(r Robot) string {
	return r.getDetails()
}

func interfaceRun() {
	robots := []Robot{
		&T1000{
			Position: "Earth",
			Strength: 100,
			Details: Details{
				Name:        "Arnold",
				PhoneNumber: "010",
			},
		},
		&Data{
			Position:     "Enterprise",
			Intelligence: 120,
			Details: Details{
				Name:        "Data",
				PhoneNumber: "0203102",
			},
		},
	}

	for _, robot := range robots {
		// Note: there are two types of invocation here to investigate further in the future.
		// One calls a function that wraps around the concrete types method
		// Two calls a concrete types method directly
		// Both work and I can see some use-cases for each but I need to investigate further in the future for additional details
		fmt.Printf("%s and is located on: %s\n", getDetails(robot), robot.getPosition())
	}
}

func init() {
	examples := runs{
		{"Example Interfaces", interfaceRun},
	}
	GetMyExamples().Add("interfaces", examples.runExamples)
}
