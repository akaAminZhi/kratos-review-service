package snowflake

import (
	"errors"
	"time"

	sf "github.com/bwmarrin/snowflake"
)

var node *sf.Node

var (
	InavlidInitParamErr      = errors.New("snowflake init fail, invalid startTime or machineID")
	InavlidInitTimeFormatErr = errors.New("snowflake init fail, invalid startTime format")
)

func Init(startTime string, machineID int64) error {
	if len(startTime) == 0 || machineID <= 0 {
		return InavlidInitParamErr
	}
	var st time.Time
	st, err := time.Parse("2006-01-02", startTime)
	if err != nil {
		return InavlidInitTimeFormatErr
	}

	sf.Epoch = st.UnixNano() / 1000000
	node, err = sf.NewNode(machineID)
	return nil
}

func GenID() int64 {
	return node.Generate().Int64()
}
