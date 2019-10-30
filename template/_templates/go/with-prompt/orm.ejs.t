---
to: src/models/orm/orm.<%=h.changeCase.lowerCase(name)%>.go
---
package orm

// <%=name%> is an entity
type <%=name%> struct {
	OrmModel
}

func init() {
	// Migrate the schema
	db.AutoMigrate(&<%=name%>{})
}
