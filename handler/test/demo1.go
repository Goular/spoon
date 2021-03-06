package test

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"fmt"
)

const str = `{
	"tooltip": {
		"trigger": "axis"
	},
	"xAxis": {
		"type": "time"
	},
	"yAxis": {
		"type": "value",
		"axisLabel": {
			"formatter": "{value}℃"
		}
	},
	"series": [{
		"data": [{
			"name": "时间",
			"value": [1531299908000, 39]
		}, {
			"name": "时间",
			"value": [1531299878000, 39]
		}, {
			"name": "时间",
			"value": [1531299848000, 39.1]
		}, {
			"name": "时间",
			"value": [1531299818000, 39.2]
		}, {
			"name": "时间",
			"value": [1531299788000, 39.2]
		}, {
			"name": "时间",
			"value": [1531299758000, 39.3]
		}, {
			"name": "时间",
			"value": [1531299728000, 39.3]
		}, {
			"name": "时间",
			"value": [1531299698000, 39.4]
		}, {
			"name": "时间",
			"value": [1531299668000, 39.4]
		}, {
			"name": "时间",
			"value": [1531299638000, 39.5]
		}, {
			"name": "时间",
			"value": [1531299608000, 39.5]
		}, {
			"name": "时间",
			"value": [1531299578000, 39.6]
		}, {
			"name": "时间",
			"value": [1531299548000, 39.7]
		}, {
			"name": "时间",
			"value": [1531299518000, 39.7]
		}, {
			"name": "时间",
			"value": [1531299488000, 39.7]
		}, {
			"name": "时间",
			"value": [1531299458000, 39.5]
		}, {
			"name": "时间",
			"value": [1531299428000, 39.1]
		}, {
			"name": "时间",
			"value": [1531299398000, 39]
		}, {
			"name": "时间",
			"value": [1531299368000, 38.9]
		}, {
			"name": "时间",
			"value": [1531299338000, 39]
		}, {
			"name": "时间",
			"value": [1531299308000, 39]
		}, {
			"name": "时间",
			"value": [1531299278000, 39.1]
		}, {
			"name": "时间",
			"value": [1531299248000, 39.2]
		}, {
			"name": "时间",
			"value": [1531299218000, 39.2]
		}, {
			"name": "时间",
			"value": [1531299188000, 39.3]
		}, {
			"name": "时间",
			"value": [1531299158000, 39.3]
		}, {
			"name": "时间",
			"value": [1531299128000, 39.4]
		}, {
			"name": "时间",
			"value": [1531299098000, 39.4]
		}, {
			"name": "时间",
			"value": [1531299068000, 39.5]
		}, {
			"name": "时间",
			"value": [1531299038000, 39.5]
		}, {
			"name": "时间",
			"value": [1531299008000, 39.6]
		}, {
			"name": "时间",
			"value": [1531298978000, 39.6]
		}, {
			"name": "时间",
			"value": [1531298948000, 39.6]
		}, {
			"name": "时间",
			"value": [1531298918000, 39.6]
		}, {
			"name": "时间",
			"value": [1531298888000, 39.2]
		}, {
			"name": "时间",
			"value": [1531298858000, 39.3]
		}, {
			"name": "时间",
			"value": [1531298828000, 39]
		}, {
			"name": "时间",
			"value": [1531298798000, 38.9]
		}, {
			"name": "时间",
			"value": [1531298768000, 39]
		}, {
			"name": "时间",
			"value": [1531298738000, 39]
		}, {
			"name": "时间",
			"value": [1531298708000, 39.1]
		}, {
			"name": "时间",
			"value": [1531298678000, 39.1]
		}, {
			"name": "时间",
			"value": [1531298648000, 39.2]
		}]
	}]
}`

func Demo01(c *gin.Context)  {
	c.String(http.StatusOK,"%s",str)
	fmt.Println(str)
}
