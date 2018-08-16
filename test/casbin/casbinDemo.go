package casbin

import (
	"github.com/casbin/casbin"
	"github.com/casbin/gorm-adapter"
	_ "github.com/go-sql-driver/mysql"
	"fmt"
)

func main() {
	// Initialize a Gorm adapter and use it in a Casbin enforcer:
	// The adapter will use the MySQL database named "casbins".
	// If it doesn't exist, the adapter will create it automatically.
	a := gormadapter.NewAdapter("mysql", "root:123456@tcp(127.0.0.1:3306)/db_apiserver", true) // Your driver and data source.
	e := casbin.NewEnforcer("conf/rbac_model.conf", a)
	// e.EnableLog(false)
	// Or you can use an existing DB "abc" like this:
	// The adapter will use the table named "casbin_rule".
	// If it doesn't exist, the adapter will create it automatically.
	// a := gormadapter.NewAdapter("mysql", "mysql_username:mysql_password@tcp(127.0.0.1:3306)/abc", true)

	// Load the policy from DB.
	e.LoadPolicy()

	// Check the permission.
	e.Enforce("alice", "data1", "read")

	// Modify the policy.
	e.AddPolicy("data2_admin", "data2", "write")
	e.AddGroupingPolicy("goular","admin")
	e.Enforce("data2_admin", "data2", "write")
	flag := e.RemoveGroupingPolicy("goular","root")
	fmt.Println(flag)
	// e.RemovePolicy(...)

	// Save the policy back to DB.
	e.SavePolicy()
}
