(function(t){function e(e){for(var o,c,a=e[0],u=e[1],l=e[2],d=0,p=[];d<a.length;d++)c=a[d],Object.prototype.hasOwnProperty.call(i,c)&&i[c]&&p.push(i[c][0]),i[c]=0;for(o in u)Object.prototype.hasOwnProperty.call(u,o)&&(t[o]=u[o]);s&&s(e);while(p.length)p.shift()();return r.push.apply(r,l||[]),n()}function n(){for(var t,e=0;e<r.length;e++){for(var n=r[e],o=!0,a=1;a<n.length;a++){var u=n[a];0!==i[u]&&(o=!1)}o&&(r.splice(e--,1),t=c(c.s=n[0]))}return t}var o={},i={app:0},r=[];function c(e){if(o[e])return o[e].exports;var n=o[e]={i:e,l:!1,exports:{}};return t[e].call(n.exports,n,n.exports,c),n.l=!0,n.exports}c.m=t,c.c=o,c.d=function(t,e,n){c.o(t,e)||Object.defineProperty(t,e,{enumerable:!0,get:n})},c.r=function(t){"undefined"!==typeof Symbol&&Symbol.toStringTag&&Object.defineProperty(t,Symbol.toStringTag,{value:"Module"}),Object.defineProperty(t,"__esModule",{value:!0})},c.t=function(t,e){if(1&e&&(t=c(t)),8&e)return t;if(4&e&&"object"===typeof t&&t&&t.__esModule)return t;var n=Object.create(null);if(c.r(n),Object.defineProperty(n,"default",{enumerable:!0,value:t}),2&e&&"string"!=typeof t)for(var o in t)c.d(n,o,function(e){return t[e]}.bind(null,o));return n},c.n=function(t){var e=t&&t.__esModule?function(){return t["default"]}:function(){return t};return c.d(e,"a",e),e},c.o=function(t,e){return Object.prototype.hasOwnProperty.call(t,e)},c.p="/";var a=window["webpackJsonp"]=window["webpackJsonp"]||[],u=a.push.bind(a);a.push=e,a=a.slice();for(var l=0;l<a.length;l++)e(a[l]);var s=u;r.push([0,"chunk-vendors"]),n()})({0:function(t,e,n){t.exports=n("56d7")},"339b":function(t,e,n){"use strict";var o=n("3867"),i=n.n(o);i.a},3867:function(t,e,n){},"56d7":function(t,e,n){"use strict";n.r(e);n("e260"),n("e6cf"),n("cca6"),n("a79d");var o=n("2b0e"),i=function(){var t=this,e=t.$createElement,n=t._self._c||e;return n("div",{attrs:{id:"app"}},[n("h1",[t._v("Things to do")]),n("TodoList")],1)},r=[],c=function(){var t=this,e=t.$createElement,n=t._self._c||e;return n("div",[n("ol",[n("draggable",{attrs:{draggable:".item",move:t.onMove},model:{value:t.items,callback:function(e){t.items=e},expression:"items"}},t._l(t.items,(function(e){return n("li",{key:e.id,staticClass:"item"},[n("input",{directives:[{name:"model",rawName:"v-model",value:e.complete,expression:"item.complete"}],attrs:{type:"checkbox"},domProps:{checked:Array.isArray(e.complete)?t._i(e.complete,null)>-1:e.complete},on:{change:[function(n){var o=e.complete,i=n.target,r=!!i.checked;if(Array.isArray(o)){var c=null,a=t._i(o,c);i.checked?a<0&&t.$set(e,"complete",o.concat([c])):a>-1&&t.$set(e,"complete",o.slice(0,a).concat(o.slice(a+1)))}else t.$set(e,"complete",r)},function(n){return t.updateTodo(e)}]}}),e.editing?t._e():n("span",{staticClass:"itemTitle"},[e.complete?t._e():n("span",[t._v(t._s(e.title))]),e.complete?n("strike",[t._v(t._s(e.title))]):t._e()],1),e.editing?t._e():n("button",{on:{click:function(n){return t.startEditTodo(e)}}},[t._v(" Edit ")]),e.editing?t._e():n("button",{on:{click:function(n){return t.deleteTodo(e)}}},[t._v(" Delete ")]),e.editing?n("div",{staticClass:"editTools"},[n("input",{directives:[{name:"model",rawName:"v-model",value:e.editedTitle,expression:"item.editedTitle"}],staticClass:"editedTitle",attrs:{type:"text"},domProps:{value:e.editedTitle},on:{input:function(n){n.target.composing||t.$set(e,"editedTitle",n.target.value)}}}),n("button",{on:{click:function(n){return t.saveEditTodo(e)}}},[t._v("Save")]),n("button",{on:{click:function(n){return t.cancelEditTodo(e)}}},[t._v("Cancel")])]):t._e()])})),0)],1),n("div",{attrs:{slot:"footer"},slot:"footer"},[n("label",[t._v(" Add a new item: "),n("input",{directives:[{name:"model",rawName:"v-model",value:t.newItemTitle,expression:"newItemTitle"}],attrs:{type:"text"},domProps:{value:t.newItemTitle},on:{input:function(e){e.target.composing||(t.newItemTitle=e.target.value)}}})]),n("button",{on:{click:function(e){return t.createNewTodo()}}},[t._v("Create")])])])},a=[],u=(n("a4d3"),n("99af"),n("4de4"),n("4160"),n("d81d"),n("e439"),n("dbb4"),n("b64b"),n("159b"),n("2909")),l=n("ade3"),s=n("310e"),d=n.n(s),p=n("bc3a"),f=n.n(p),m=n("2ef0"),v=n.n(m),b=void 0;function h(t,e){var n=Object.keys(t);if(Object.getOwnPropertySymbols){var o=Object.getOwnPropertySymbols(t);e&&(o=o.filter((function(e){return Object.getOwnPropertyDescriptor(t,e).enumerable}))),n.push.apply(n,o)}return n}function g(t){for(var e=1;e<arguments.length;e++){var n=null!=arguments[e]?arguments[e]:{};e%2?h(Object(n),!0).forEach((function(e){Object(l["a"])(t,e,n[e])})):Object.getOwnPropertyDescriptors?Object.defineProperties(t,Object.getOwnPropertyDescriptors(n)):h(Object(n)).forEach((function(e){Object.defineProperty(t,e,Object.getOwnPropertyDescriptor(n,e))}))}return t}function y(t){return f.a.put("/api/v1/todos/".concat(t.id),t)}var T={name:"TodoList",data:function(){return{items:[],error:null,newItemTitle:""}},components:{draggable:d.a},mounted:function(){var t=this;f.a.get("/api/v1/todos").then((function(e){t.items=v.a.orderBy(e.data,(function(t){return t.position})).map((function(t){return g({},t,{editing:!1,editedTitle:""})}))})).catch((function(e){t.error=e}))},methods:{updateTodo:function(t){y(t).catch((function(t){b.error=t}))},startEditTodo:function(t){t.editing=!0,t.editedTitle=t.title},saveEditTodo:function(t){t.editing=!1,t.title=t.editedTitle,y(t).catch((function(t){b.error=t}))},cancelEditTodo:function(t){t.editing=!1},createNewTodo:function(){var t=this,e={id:Math.max.apply(Math,Object(u["a"])(this.items.map((function(t){return t.id}))).concat([-1]))+1,title:this.newItemTitle,position:Math.max.apply(Math,Object(u["a"])(this.items.map((function(t){return t.position}))).concat([-1]))+1,complete:!1};this.items.push(e),this.newItemTitle="",f.a.post("/api/v1/todos",e).catch((function(e){t.error=e}))},deleteTodo:function(t){var e=this;this.items=this.items.filter((function(e){return e.id!==t.id})),f.a.delete("/api/v1/todos/".concat(t.id)).catch((function(t){e.error=t}))},onMove:function(t){var e=this,n=t.draggedContext.index,o=t.draggedContext.futureIndex,i=this.items[n],r=this.items[o],c=r.position;r.position=i.position,i.position=c,y(i).catch((function(t){e.error=t})),y(r).catch((function(t){e.error=t}))}}},O=T,w=(n("339b"),n("2877")),_=Object(w["a"])(O,c,a,!1,null,"961b237e",null),j=_.exports,x={name:"app",components:{TodoList:j},data:function(){return{}}},P=x,k=Object(w["a"])(P,i,r,!1,null,null,null),E=k.exports;o["a"].config.productionTip=!1,new o["a"]({render:function(t){return t(E)}}).$mount("#app")}});
//# sourceMappingURL=app.3f0a62e4.js.map