import{h}from"./disclose-version.DjcaTDsi.js";import{E as L,F as w,G as C,H as N,I as $,J as j,K as p,L as I}from"./runtime.BZ58vcTB.js";import{d as B}from"./render.FyCEabsq.js";function D(s,i){if(i){const t=document.body;s.autofocus=!0,L(()=>{document.activeElement===t&&s.focus()})}}function T(s){h&&s.firstChild!==null&&(s.textContent="")}function U(s){h&&(n(s,"value",null),n(s,"checked",null))}function n(s,i,t){t=t==null?null:t+"";var f=s.__attributes??(s.__attributes={});h&&(f[i]=s.getAttribute(i),i==="src"||i==="href"||i==="srcset")||f[i]!==(f[i]=t)&&(t===null?s.removeAttribute(i):s.setAttribute(i,t))}function F(s,i,t){if(i in s){var f=s[i],a=typeof f=="boolean"&&t===""?!0:t;(typeof f!="object"||f!==a)&&(s[i]=a)}else n(s,i,t)}function G(s,i,t,f,a){var c=w({},...t),g=a.length!==0;for(var r in i)r in c||(c[r]=null);g&&!c.class&&(c.class="");var l=C(E,s.nodeName);l||N(E,s.nodeName,l=J(s));for(r in c){var o=c[r];if(o!==(i==null?void 0:i[r])){var b=r[0]+r[1];if(b!=="$$")if(b==="on"){var d={},_=r.slice(2),A=p.includes(_);_.endsWith("capture")&&_!=="ongotpointercapture"&&_!=="onlostpointercapture"&&(_=_.slice(0,-7),d.capture=!0),!A&&(i!=null&&i[r])&&s.removeEventListener(_,i[r],d),o!=null&&(A?(s[`__${_}`]=o,B([_])):s.addEventListener(_,o,d))}else if(o==null)s.removeAttribute(r);else if(r==="style")s.style.cssText=o+"";else if(r==="autofocus")D(s,!!o);else if(r==="__value"||r==="value")s.value=s[r]=s.__value=o;else{var u=r;f&&(u=u.toLowerCase(),u=$[u]||u),l.includes(u)?h&&(u==="src"||u==="href"||u==="srcset")||(s[u]=o):typeof o!="function"&&(g&&u==="class"&&(o&&(o+=" "),o+=a),n(s,u,o))}}}return c}function W(s,i,t,f){if(s.tagName.includes("-")){var a=w({},...t);for(var c in i)c in a||(a[c]=null);for(c in a)F(s,c,a[c]);return a}return G(s,i,t,s.namespaceURI!==I,f)}var H=["width","height"],E=new Map;function J(s){var i=[],t=j(s.__proto__);for(var f in t)t[f].set&&!H.includes(f)&&i.push(f);return i}export{T as a,W as b,U as r,G as s};