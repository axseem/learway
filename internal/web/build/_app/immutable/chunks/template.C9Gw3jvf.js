import{l as T,o as N,q as H,t as O,v as $,w as j,P as C,x as q,e as z,a as B,c as F,p as G,d as U,y as W,T as P,z as Y}from"./runtime.DYeCD_Dl.js";let l=!1;function x(t){l=t}let i=null;function J(t){i=t}function E(t){if(t.nodeType!==8)return t;var e=t;if((e==null?void 0:e.data)!=="[")return t;for(var n=[],r=0;(e=e.nextSibling)!==null;){if(e.nodeType===8){var o=e.data;if(o==="[")r+=1;else if(o==="]"){if(r===0)return i=n,e;r-=1}}n.push(e)}throw new Error("Expected a closing hydration marker")}var d,m,S,k,w,A,R;function K(){d===void 0&&(d=Node.prototype,m=Element.prototype,S=Text.prototype,d.appendChild,k=d.cloneNode,m.__click=void 0,S.__nodeValue=" ",m.__className="",m.__attributes=null,w=T(d,"firstChild").get,A=T(d,"nextSibling").get,R=T(d,"textContent").set,T(m,"className").set)}function Q(t,e){return k.call(t,e)}function v(){return document.createTextNode("")}function oe(t){const e=w.call(t);return l?e===null?t.appendChild(v()):E(e):e}function ae(t,e){if(!l)return w.call(t);const n=t[0];if(e&&(n==null?void 0:n.nodeType)!==3){const r=v();return i.unshift(r),n==null||n.before(r),r}return E(n)}function ie(t,e=!1){const n=A.call(t);if(!l)return n;if(e&&(n==null?void 0:n.nodeType)!==3){const r=v();if(n){const o=i.indexOf(n);i.splice(o,0,r),n.before(r)}else i.push(r);return r}return E(n)}function X(t){R.call(t,"")}function M(t,e){var u;var n=t.ownerDocument,r=e.type,o=((u=e.composedPath)==null?void 0:u.call(e))||[],a=o[0]||e.target;e.target!==a&&N(e,"target",{configurable:!0,value:a});var s=0,c=e.__root;if(c){var _=o.indexOf(c);if(_!==-1&&(t===document||t===window)){e.__root=t;return}var b=o.indexOf(t);if(b===-1)return;_<=b&&(s=_+1)}for(a=o[s]||e.target,N(e,"currentTarget",{configurable:!0,get(){return a||n}});a!==null;){var p=a.parentNode||a.host||null,y="__"+r,h=a[y];if(h!==void 0&&!a.disabled)if(H(h)){var[f,...g]=h;f.apply(a,[e,...g])}else h.call(a,e);if(e.cancelBubble||p===t||a===t)break;a=p}e.__root=t,a=t}const Z=new Set,V=new Set;let L=!0;function se(t,e){const n=t.__nodeValue,r=ee(e);l&&t.nodeValue===r?t.__nodeValue=r:n!==r&&(t.nodeValue=r,t.__nodeValue=r)}function ue(t,e,n,r){e===void 0?r!==null&&r(t):e(t,n)}function ee(t){return typeof t=="string"?t:t==null?"":t+""}function te(t,e){const n=e.target.appendChild(v());return O(()=>D(t,{...e,anchor:n}),!1)}function ce(t,e){const n=e.target,r=n.firstChild,o=i;let a=!1;try{return O(()=>{x(!0);const s=E(r),c=D(t,{...e,anchor:s});return x(!1),a=!0,c},!1)}catch(s){if(!a&&e.recover!==!1)return console.error("ERR_SVELTE_HYDRATION_MISMATCH",s),X(n),x(!1),te(t,e);throw s}finally{x(!!o),J(o)}}function D(t,{target:e,anchor:n,props:r={},events:o,context:a,intro:s=!1}){K();const c=new Set,_=M.bind(null,e),b=M.bind(null,document),p=f=>{for(let g=0;g<f.length;g++){const u=f[g];c.has(u)||(c.add(u),e.addEventListener(u,_,C.includes(u)?{passive:!0}:void 0),document.addEventListener(u,b,C.includes(u)?{passive:!0}:void 0))}};p($(Z)),V.add(p);let y;const h=j(()=>(q(()=>{z(()=>{if(a){B({});var f=F;f.c=a}o&&(r.$$events=o),L=s,y=t(n,r)||{},L=!0,a&&G()})}),()=>{for(const f of c)e.removeEventListener(f,_);V.delete(p)}));return I.set(y,h),y}let I=new WeakMap;function fe(t){const e=I.get(t);e==null||e()}function ne(t,e){var n=(e&P)!==0,r=(e&Y)!==0,o;return()=>l?n?i:i[0]:(o||(o=W(t),n||(o=o.firstChild)),r?document.importNode(o,!0):Q(o,!0))}function le(t){if(!l)return v();var e=i[0];return e||t.before(e=v()),e}const de=ne("<!>",P);function _e(t,e){var n=e;if(!l){var r=e;r.nodeType===11&&(n=[...r.childNodes]),t.before(r)}U.dom=n}export{_e as a,ie as b,oe as c,se as d,i as e,ae as f,x as g,l as h,ce as i,de as j,le as k,te as m,ue as s,ne as t,fe as u};
