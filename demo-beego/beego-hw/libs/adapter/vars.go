package adapter

import (
	"encoding/json"
	"os"
)

var Stat = os.Stat
var Marshal = json.Marshal
var UnMarshal = json.Unmarshal