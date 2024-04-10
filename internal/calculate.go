package ligma

import (
	"os"
	"time"

	"gopkg.in/yaml.v3"
)

type Conf struct {
  WorkHours float64 `json:"work_hours" yaml:"work_hours"`
  BreakTime float64 `json:"break_time" yaml:"break_time"`
}

func NewConf() (*Conf, error) {
  c := &Conf{}
  return c.readOrInitConfig()
}

func (c *Conf) readOrInitConfig() (*Conf, error) {
    // Read the config file if it exists, otherwise create a new one
    yamlFile, err := os.ReadFile("config.yaml")
    if err != nil {
        // If config file doesn't exist, create a new one
      defaultConf := Conf{WorkHours: 7.8, BreakTime: 0.5} // Default values
        data, err := yaml.Marshal(defaultConf)
        if err != nil {
            return nil, err
        }
        err = os.WriteFile("config.yaml", data, 0644)
        if err != nil {
            return nil, err
        }
        return &defaultConf, nil
    }

    err = yaml.Unmarshal(yamlFile, c)
    if err != nil {
        return nil, err
    }

    return c, nil
}

func (c *Conf) WriteToConfig(newConf Conf) error {
    data, err := yaml.Marshal(newConf)
    if err != nil {
        return err
    }
    err = os.WriteFile("config.yaml", data, 0644)
    if err != nil {
        return err
    }
    return nil
}

func FormatToTime(s string) (*time.Time, error) {
  // Parse the time from a string
  // we will only get a time in the format of HH:MM
  t, err := time.Parse("15:04", s)
  if err != nil {
    return nil, err
  }
  return &t, nil
}

func TimeToGo(start time.Time) (*time.Time, error) {
  // Calculate the next time after x hours in the form of float hours
  conf, err := NewConf()
  if err != nil {
    return nil, err
  }

  next := start.Add(time.Duration((conf.WorkHours + conf.BreakTime) * float64(time.Hour)))

  return &next, nil
}

func CalculateTotalHoursBetweenTwoTimes(start, end time.Time) float64 {
  return end.Sub(start).Hours()
}
