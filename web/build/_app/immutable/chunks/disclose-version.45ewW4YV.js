import{w as l,h as T,a1 as E,a2 as u,a3 as w}from"./runtime.BZ58vcTB.js";let o=!1;function b(t){o=t}let a=null;function S(t){a=t}function _(t){if(t.nodeType!==8)return t;var e=t;if((e==null?void 0:e.data)!=="[")return t;for(var n=[],r=0;(e=e.nextSibling)!==null;){if(e.nodeType===8){var i=e.data;if(i==="[")r+=1;else if(i==="]"){if(r===0)return a=n,e;r-=1}}n.push(e)}throw new Error("Expected a closing hydration marker")}var s,f,h,m,p,v,y;function M(){s===void 0&&(s=Node.prototype,f=Element.prototype,h=Text.prototype,s.appendChild,m=s.cloneNode,f.__click=void 0,h.__nodeValue=" ",f.__className="",f.__attributes=null,p=l(s,"firstChild").get,v=l(s,"nextSibling").get,y=l(s,"textContent").set,l(f,"className").set)}function x(t,e){return m.call(t,e)}function c(){return document.createTextNode("")}function O(t){const e=p.call(t);return o?e===null?t.appendChild(c()):_(e):e}function P(t,e){if(!o)return p.call(t);const n=t[0];if(e&&(n==null?void 0:n.nodeType)!==3){const r=c();return a.unshift(r),n==null||n.before(r),r}return _(n)}function k(t,e=!1){const n=v.call(t);if(!o)return n;if(e&&(n==null?void 0:n.nodeType)!==3){const r=c();if(n){const i=a.indexOf(n);a.splice(i,0,r),n.before(r)}else a.push(r);return r}return _(n)}function A(t){y.call(t,"")}function g(t,e){var n=(e&u)!==0,r=(e&w)!==0,i;return()=>o?n?a:a[0]:(i||(i=E(t),n||(i=i.firstChild)),r?document.importNode(i,!0):x(i,!0))}function I(t,e){var n=(e&u)!==0,r=g(`<svg>${t}</svg>`,0),i;return()=>{if(o)return n?a:a[0];if(!i){var d=r();if(!(e&u))i=d.firstChild;else for(i=document.createDocumentFragment();d.firstChild;)i.appendChild(d.firstChild)}return x(i,!0)}}function L(t){if(!o)return c();var e=a[0];return e||t.before(e=c()),e}const R=g("<!>",u);function D(t,e){var n=e;if(!o){var r=e;r.nodeType===11&&(n=[...r.childNodes]),t.before(r)}T.dom=n}const C="5";typeof window<"u"&&(window.__svelte||(window.__svelte={v:new Set})).v.add(C);export{D as a,R as b,O as c,L as d,c as e,P as f,b as g,o as h,_ as i,A as j,M as k,S as l,a as m,I as n,k as s,g as t};