---
to: src/router/router.admin.go
inject: true
after: // CodeGenPlace
skip_if: var <%=name%> ctrl.<%=h.changeCase.lowerCase(name)%>
---

	// <%=name%>
	var <%=h.changeCase.lowerCase(name)%> ctrl.<%=name%>
	r.GET("/admin/<%=h.changeCase.lowerCase(name)%>s/:id", ctrl.ShouldBeAdmin, <%=h.changeCase.lowerCase(name)%>.GetByID)
	r.POST("/admin/<%=h.changeCase.lowerCase(name)%>s", ctrl.ShouldBeAdmin, <%=h.changeCase.lowerCase(name)%>.Create)
	r.PUT("/admin/<%=h.changeCase.lowerCase(name)%>s/:id", ctrl.ShouldBeAdmin, <%=h.changeCase.lowerCase(name)%>.Update)
	r.DELETE("/admin/<%=h.changeCase.lowerCase(name)%>s/:id", ctrl.ShouldBeAdmin, <%=h.changeCase.lowerCase(name)%>.Delete)
	r.POST("/admin/<%=h.changeCase.lowerCase(name)%>s/bypage", ctrl.ShouldBeAdmin, <%=h.changeCase.lowerCase(name)%>.GetAllByPage)
