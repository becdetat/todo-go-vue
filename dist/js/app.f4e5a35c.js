(function(t){function e(e){for(var o,i,u=e[0],c=e[1],l=e[2],p=0,d=[];p<u.length;p++)i=u[p],Object.prototype.hasOwnProperty.call(r,i)&&r[i]&&d.push(r[i][0]),r[i]=0;for(o in c)Object.prototype.hasOwnProperty.call(c,o)&&(t[o]=c[o]);s&&s(e);while(d.length)d.shift()();return a.push.apply(a,l||[]),n()}function n(){for(var t,e=0;e<a.length;e++){for(var n=a[e],o=!0,u=1;u<n.length;u++){var c=n[u];0!==r[c]&&(o=!1)}o&&(a.splice(e--,1),t=i(i.s=n[0]))}return t}var o={},r={app:0},a=[];function i(e){if(o[e])return o[e].exports;var n=o[e]={i:e,l:!1,exports:{}};return t[e].call(n.exports,n,n.exports,i),n.l=!0,n.exports}i.m=t,i.c=o,i.d=function(t,e,n){i.o(t,e)||Object.defineProperty(t,e,{enumerable:!0,get:n})},i.r=function(t){"undefined"!==typeof Symbol&&Symbol.toStringTag&&Object.defineProperty(t,Symbol.toStringTag,{value:"Module"}),Object.defineProperty(t,"__esModule",{value:!0})},i.t=function(t,e){if(1&e&&(t=i(t)),8&e)return t;if(4&e&&"object"===typeof t&&t&&t.__esModule)return t;var n=Object.create(null);if(i.r(n),Object.defineProperty(n,"default",{enumerable:!0,value:t}),2&e&&"string"!=typeof t)for(var o in t)i.d(n,o,function(e){return t[e]}.bind(null,o));return n},i.n=function(t){var e=t&&t.__esModule?function(){return t["default"]}:function(){return t};return i.d(e,"a",e),e},i.o=function(t,e){return Object.prototype.hasOwnProperty.call(t,e)},i.p="/";var u=window["webpackJsonp"]=window["webpackJsonp"]||[],c=u.push.bind(u);u.push=e,u=u.slice();for(var l=0;l<u.length;l++)e(u[l]);var s=c;a.push([0,"chunk-vendors"]),n()})({0:function(t,e,n){t.exports=n("56d7")},"034f":function(t,e,n){"use strict";var o=n("85ec"),r=n.n(o);r.a},"56d7":function(t,e,n){"use strict";n.r(e);n("e260"),n("e6cf"),n("cca6"),n("a79d");var o=n("2b0e"),r=function(){var t=this,e=t.$createElement,n=t._self._c||e;return n("div",{attrs:{id:"app"}},[n("md-app",[n("md-app-toolbar",{staticClass:"md-primary md-large"},[n("div",{staticClass:"md-toolbar-row"},[n("span",{staticClass:"md-title"},[t._v("Things to do")])])]),n("md-app-content",[n("TodoList")],1)],1)],1)},a=[],i=function(){var t=this,e=t.$createElement,n=t._self._c||e;return n("div",[n("md-list",t._l(t.items,(function(e){return n("md-list-item",{key:e.id},[n("md-checkbox",{on:{change:function(n){return t.updateTodo(e)}},model:{value:e.complete,callback:function(n){t.$set(e,"complete",n)},expression:"item.complete"}}),n("span",{staticClass:"md-list-item-text"},[e.complete?t._e():n("span",[t._v(t._s(e.title))]),e.complete?n("strike",[t._v(t._s(e.title))]):t._e()],1)],1)})),1)],1)},u=[],c=n("bc3a"),l=n.n(c),s=n("2ef0"),p=n.n(s),d=void 0,f={name:"TodoList",data:function(){return{items:[],error:null}},mounted:function(){var t=this;l.a.get("https://thawing-bayou-17829.herokuapp.com/todos").then((function(e){t.items=p.a.orderBy(e.data,(function(t){return t.position}))})).catch((function(e){t.error=e}))},methods:{updateTodo:function(t){l.a.put("https://thawing-bayou-17829.herokuapp.com/todos/".concat(t.id),t).catch((function(t){d.error=t}))}}},m=f,h=n("2877"),b=Object(h["a"])(m,i,u,!1,null,"16db57a9",null),v=b.exports,y={name:"app",components:{TodoList:v},data:function(){return{}}},g=y,_=(n("034f"),Object(h["a"])(g,r,a,!1,null,null,null)),w=_.exports,O=n("43f9"),j=n.n(O);n("51de"),n("e094");o["default"].config.productionTip=!1,o["default"].use(j.a),new o["default"]({render:function(t){return t(w)}}).$mount("#app")},"85ec":function(t,e,n){}});
//# sourceMappingURL=app.f4e5a35c.js.map