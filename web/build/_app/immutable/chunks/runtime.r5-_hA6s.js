var Ie=Array.isArray,Ke=Array.from,ze=Object.assign,Ge=Object.isFrozen,Xe=Object.defineProperty,Ze=Object.getOwnPropertyDescriptor,we=Object.getOwnPropertyDescriptors,$e=Object.prototype,Qe=Array.prototype,We=Object.getPrototypeOf,te=Map.prototype,Ne=te.set,Ce=te.get;function Je(e,n,t){Ne.call(e,n,t)}function en(e,n){return Ce.call(e,n)}function nn(e){return typeof e=="function"}const P=2,q=4,R=8,re=16,E=32,G=64,I=128,m=256,T=512,N=1024,C=2048,k=4096,oe=8192,ke=16384,J=Symbol("$state");function se(e){return e===this.v}function be(e,n){return e!=e?n==n:e!==n||e!==null&&typeof e=="object"||typeof e=="function"}function ue(e){return!be(e,this.v)}const tn=1,rn=2,on=4,sn=8,un=64,fn=1,an=2,ln=4,cn=8,_n=1,pn=2,vn=4,dn=1,hn=2,xe=Symbol(),En=["beforeinput","click","dblclick","contextmenu","focusin","focusout","keydown","keyup","mousedown","mousemove","mouseout","mouseover","mouseup","pointerdown","pointermove","pointerout","pointerover","pointerup","touchend","touchmove","touchstart"],yn=["touchstart","touchmove","touchend"],mn={formnovalidate:"formNoValidate",ismap:"isMap",nomodule:"noModule",playsinline:"playsInline",readonly:"readOnly"},Tn="http://www.w3.org/2000/svg";function fe(e){return{f:0,reactions:null,equals:se,v:e,version:0}}function An(e){const n=fe(e);return n.equals=ue,l&&(l.d??(l.d=[])).push(n),n}function Le(e,n){var t=e.v!==xe;if(!O&&t&&v!==null&&K()&&v.f&P)throw new Error("ERR_SVELTE_UNSAFE_MUTATION");return e.equals(n)||(e.v=n,e.version++,K()&&t&&a!==null&&a.f&m&&!(a.f&E)&&(p!==null&&p.includes(e)?(y(a,T),V(a)):A===null?qe([e]):A.push(e)),Q(e,T,!0)),n}const ae=()=>{};function On(e){return e()}function Sn(e){for(var n=0;n<e.length;n++)e[n]()}let ie=!1;function gn(e){ie=e}let X=null;function Rn(e){X=e}function In(e){if(e.nodeType!==8)return e;var n=e;if((n==null?void 0:n.data)!=="[")return e;for(var t=[],r=0;(n=n.nextSibling)!==null;){if(n.nodeType===8){var o=n.data;if(o==="[")r+=1;else if(o==="]"){if(r===0)return X=t,n;r-=1}}t.push(n)}throw new Error("Expected a closing hydration marker")}function Pe(e){var n=document.createElement("template");return n.innerHTML=e,n.content}function Fe(e){if(Ie(e))for(var n=0;n<e.length;n++){var t=e[n];t.isConnected&&t.remove()}else e.isConnected&&e.remove()}function wn(e,n,t){if(ie)return X;var r=n+"",o;t&&(r=`<svg>${r}</svg>`);var s=Pe(r);t&&(s=s.firstChild);var i=s.cloneNode(!0);return o=[...i.childNodes],o.forEach(f=>{e.before(f)}),o}function F(e,n,t){var r=(e&G)!==0,o={parent:r?null:a,dom:null,deps:null,f:e|T,fn:n,effects:null,teardown:null,ctx:l,transitions:null};if(v!==null&&!r&&(v.effects===null?v.effects=[o]:v.effects.push(o)),t){var s=w;try{ee(!0),M(o),o.f|=ke}finally{ee(s)}}else V(o);return o}function Nn(){return a?(a.f&(E|G))===0:!1}function Cn(e){if(a===null)throw new Error("ERR_SVELTE_ORPHAN_EFFECT");if(a.f&R&&l!==null&&!l.m){const t=l;(t.e??(t.e=[])).push(e)}else le(e)}function kn(e){if(a===null)throw new Error("ERR_SVELTE_ORPHAN_EFFECT");return Z(e)}function bn(e){const n=F(G,()=>ge(e),!0);return()=>{D(n)}}function le(e){return F(q,e,!1)}function xn(e,n){var t=l,r={effect:null,ran:!1};t.l1.push(r),r.effect=Z(()=>{e(),!r.ran&&(r.ran=!0,Le(t.l2,!0),ge(n))})}function Ln(){var e=l;Z(()=>{if(Se(e.l2)){for(var n of e.l1){var t=n.effect;b(t)&&M(t),n.ran=!1}e.l2.v=!1}})}function Z(e){return F(R,e,!0)}function Pn(e){return F(R|re,e,!0)}function Fn(e){return F(R|E,e,!0)}function D(e){var o;if(ye(e),H(e,0),y(e,k),e.transitions)for(const s of e.transitions)s.stop();(o=e.teardown)==null||o.call(e),e.dom!==null&&Fe(e.dom);var n=e.parent;if(n!==null&&e.f&E){var t=n.effects;if(t!==null){var r=t.indexOf(e);t.splice(r,1)}}e.effects=e.teardown=e.ctx=e.dom=e.deps=e.parent=e.fn=null}function Dn(e,n=ae){var t=[];$(e,t,!0),ce(t,()=>{D(e),n()})}function Mn(e,n=ae){var t=[];for(var r of e)$(r,t,!0);ce(t,()=>{for(var o of e)D(o);n()})}function ce(e,n){var t=e.length;if(t>0){var r=()=>--t||n();for(var o of e)o.out(r)}else n()}function $(e,n,t){if(!(e.f&C)){if(e.f^=C,e.transitions!==null)for(const o of e.transitions)(o.is_global||t)&&n.push(o);if(e.effects!==null)for(const o of e.effects){var r=(o.f&oe)!==0||(o.f&E)!==0;$(o,n,r?t:!1)}}}function jn(e){_e(e,!0)}function _e(e,n){if(e.f&C){if(e.f^=C,b(e)&&M(e),e.effects!==null)for(const r of e.effects){var t=(r.f&oe)!==0||(r.f&E)!==0;_e(r,t?n:!1)}if(e.transitions!==null)for(const r of e.transitions)(r.is_global||n)&&r.in()}}let B=!1;function De(e){let n=P|T;a===null&&(n|=I);const t={reactions:null,deps:null,equals:se,f:n,fn:e,effects:null,deriveds:null,v:null,version:0};if(v!==null&&v.f&P){var r=v;r.deriveds===null?r.deriveds=[t]:r.deriveds.push(t)}return t}function qn(e){const n=De(e);return n.equals=ue,n}function pe(e){var n=e.effects;if(n!==null){e.effects=null;for(var t=0;t<n.length;t+=1)D(n[t])}var r=e.deriveds;if(r!==null)for(e.deriveds=null,t=0;t<r.length;t+=1)Me(r[t])}function ve(e,n){var t=B;B=!0,pe(e);var r=he(e);B=t;var o=(x||e.f&I)&&e.deps!==null?N:m;y(e,o),e.equals(r)||(e.v=r,Q(e,T,n))}function Me(e){pe(e),H(e,0),y(e,k),e.effects=e.deps=e.reactions=e.fn=null}const de=0,je=1;let j=de,U=!1,w=!1;function ee(e){w=e}let S=[],g=0,v=null,a=null,p=null,c=0,A=null;function qe(e){A=e}let O=!1,x=!1,l=null;function K(){return l!==null&&l.r}function b(e){var n=e.f;if(n&T)return!0;if(n&N){var t=e.deps;if(t!==null)for(var r=t.length,o=0;o<r;o++){var s=t[o];if(b(s)&&(ve(s,!0),e.f&T))return!0;var i=(n&I)!==0,f=s.version;if(i&&f>e.version)return e.version=f,!0}y(e,m)}return!1}function he(e){const n=p,t=c,r=A,o=v,s=x,i=O;p=null,c=0,A=null,v=e,x=!w&&(e.f&I)!==0,O=!1;try{let f=e.fn(),u=e.deps;if(p!==null){let _;if(u!==null){const d=u.length,h=c===0?p:u.slice(0,c).concat(p),W=h.length>16&&d-c>1?new Set(h):null;for(_=c;_<d;_++){const Y=u[_];(W!==null?!W.has(Y):!h.includes(Y))&&Ee(e,Y)}}if(u!==null&&c>0)for(u.length=c+p.length,_=0;_<p.length;_++)u[c+_]=p[_];else e.deps=u=p;if(!x)for(_=c;_<u.length;_++){const d=u[_],h=d.reactions;h===null?d.reactions=[e]:h[h.length-1]!==e&&h.push(e)}}else u!==null&&c<u.length&&(H(e,c),u.length=c);return f}finally{p=n,c=t,A=r,v=o,x=s,O=i}}function Ee(e,n){const t=n.reactions;let r=0;if(t!==null){r=t.length-1;const o=t.indexOf(e);o!==-1&&(r===0?n.reactions=null:(t[o]=t[r],t.pop()))}r===0&&n.f&I&&(y(n,T),H(n,0))}function H(e,n){const t=e.deps;if(t!==null){const r=n===0?null:t.slice(0,n);let o;for(o=n;o<t.length;o++){const s=t[o];(r===null||!r.includes(s))&&Ee(e,s)}}}function ye(e){var n=e.effects;if(n!==null){e.effects=null;for(var t=0;t<n.length;t+=1)D(n[t])}}function M(e){var i;var n=e.f;if(!(n&k)){y(e,m);var t=e.ctx,r=a,o=l;a=e,l=t;try{n&re||ye(e),(i=e.teardown)==null||i.call(e);var s=he(e);e.teardown=typeof s=="function"?s:null}finally{a=r,l=o}}}function me(){if(g>1e3)throw g=0,new Error("ERR_SVELTE_TOO_MANY_UPDATES");g++}function Te(e){for(var n=0;n<e.length;n++){var t=e[n];Ae(t,R|q)}}function ne(e){var n=e.length;if(n!==0){me();for(var t=0;t<n;t++){var r=e[t];!(r.f&(k|C))&&b(r)&&M(r)}}}function Ue(){if(U=!1,g>101)return;const e=S;S=[],Te(e),U||(g=0)}function V(e){j===de&&(U||(U=!0,queueMicrotask(Ue)));for(var n=e;n.parent!==null;){n=n.parent;var t=n.f;if(t&E){if(!(t&m))return;y(n,N)}}S.push(n)}function L(e,n,t,r){var o=e.effects;if(o!==null){for(var s=[],i=0;i<o.length;i++){var f=o[i],u=f.f,_=(u&(k|C))!==0;if(!_){var d=u&E,h=(u&m)!==0;if(d){if(h)continue;y(f,m)}if(u&R)if(d){if(t)continue;L(f,n,!1,r)}else b(f)&&M(f),L(f,n,!1,r);else if(u&q)if(d||h){if(t)continue;L(f,n,!1,r)}else s.push(f)}}if(s.length>0&&(n&q&&r.push(...s),!t))for(i=0;i<s.length;i++)L(s[i],n,!1,r)}}function Ae(e,n,t=!1){var r=[],o=w;w=!0;try{e.effects===null&&!(e.f&E)?ne([e]):(L(e,n,t,r),ne(r))}finally{w=o}}function Un(e){g=0,Ae(e,R,!0)}function Oe(e,n=!0){var t=j,r=S;try{me();const s=[];j=je,S=s,n&&Te(r);var o=e==null?void 0:e();return(S.length>0||s.length>0)&&Oe(),g=0,o}finally{j=t,S=r}}async function Hn(){await Promise.resolve(),Oe()}function Se(e){const n=e.f;if(n&k)return e.v;if(v!==null&&!(v.f&E)&&!O){const t=(v.f&I)!==0,r=v.deps;p===null&&r!==null&&r[c]===e&&!(t&&a!==null)?c++:(r===null||c===0||r[c-1]!==e)&&(p===null?p=[e]:p.push(e)),A!==null&&a!==null&&a.f&m&&!(a.f&E)&&A.includes(e)&&(y(a,T),V(a))}return n&P&&b(e)&&ve(e,!1),e.v}function Q(e,n,t){var r=e.reactions;if(r!==null)for(var o=K(),s=r.length,i=0;i<s;i++){var f=r[i];if(!((!t||!o)&&f===a)){var u=f.f;y(f,n);var _=(u&N)!==0,d=(u&I)!==0;(u&m||_&&d)&&(f.f&P?Q(f,N,t):V(f))}}}function ge(e){const n=O;try{return O=!0,e()}finally{O=n}}const He=~(T|N|m);function y(e,n){e.f=e.f&He|n}function Ve(e){return typeof e=="object"&&e!==null&&typeof e.f=="number"}function Vn(e){return Re().get(e)}function Yn(e,n){return Re().set(e,n),n}function Re(){const e=l;if(e===null)throw new Error("ERR_SVELTE_ORPHAN_CONTEXT");return e.c??(e.c=new Map(Ye(e)||void 0))}function Ye(e){let n=e.p;for(;n!==null;){const t=n.c;if(t!==null)return t;n=n.p}return null}function Bn(e,n=!1,t){l={x:null,c:null,e:null,m:!1,p:l,d:null,s:e,r:n,l1:[],l2:fe(!1),u:null}}function Kn(e){const n=l;if(n!==null){e!==void 0&&(n.x=e);const t=n.e;if(t!==null){n.e=null;for(let r=0;r<t.length;r++)le(t[r])}l=n.p,n.m=!0}return e||{}}function zn(e){if(!(typeof e!="object"||!e||e instanceof EventTarget)){if(J in e)z(e);else if(!Array.isArray(e))for(let n in e){const t=e[n];typeof t=="object"&&t&&J in t&&z(t)}}}function z(e,n=new Set){if(typeof e=="object"&&e!==null&&!(e instanceof EventTarget)&&!n.has(e)){n.add(e);for(let r in e)try{z(e[r],n)}catch{}const t=Object.getPrototypeOf(e);if(t!==Object.prototype&&t!==Array.prototype&&t!==Map.prototype&&t!==Set.prototype&&t!==Date.prototype){const r=we(t);for(let o in r){const s=r[o].get;if(s)try{s.call(e)}catch{}}}}}function Gn(e){return Ve(e)?Se(e):e}export{Fe as $,Nn as A,We as B,m as C,Gn as D,De as E,ie as F,ze as G,en as H,Je as I,mn as J,we as K,En as L,Tn as M,Oe as N,gn as O,In as P,Ke as Q,bn as R,J as S,Rn as T,xe as U,yn as V,X as W,Pn as X,jn as Y,Dn as Z,oe as _,Bn as a,ln as a0,ue as a1,nn as a2,fn as a3,an as a4,cn as a5,Pe as a6,dn as a7,hn as a8,wn as a9,Hn as aa,Mn as ab,tn as ac,rn as ad,un as ae,sn as af,on as ag,D as ah,ke as ai,_n as aj,vn as ak,pn as al,Vn as am,Yn as an,xn as ao,Ln as ap,qn as aq,Fn as b,kn as c,Cn as d,le as e,l as f,Se as g,Sn as h,a as i,Un as j,On as k,zn as l,An as m,ae as n,Ge as o,Kn as p,$e as q,Z as r,Le as s,Qe as t,ge as u,Xe as v,fe as w,Ie as x,B as y,Ze as z};